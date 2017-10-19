package awsspec

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/wallix/awless/aws/doc"
	"github.com/wallix/awless/cloud"
)

const (
	dryRunOperation = "DryRunOperation"
	notFound        = "NotFound"
)

type BeforeRunner interface {
	BeforeRun(ctx, params map[string]interface{}) error
}

type AfterRunner interface {
	AfterRun(ctx map[string]interface{}, output interface{}) error
}

type ManualValidator interface {
	ManualValidateCommand(params map[string]interface{}) []error
}

type command interface {
	ParamsHelp() string
	ValidateParams([]string) ([]string, error)
	ValidateCommand(map[string]interface{}) []error
	inject(params map[string]interface{}) error
	Run(ctx map[string]interface{}, params map[string]interface{}) (interface{}, error)
}

func implementsBeforeRun(i interface{}) (BeforeRunner, bool) {
	v, ok := i.(BeforeRunner)
	return v, ok
}

func implementsAfterRun(i interface{}) (AfterRunner, bool) {
	v, ok := i.(AfterRunner)
	return v, ok
}

func implementsManualValidator(i interface{}) (ManualValidator, bool) {
	v, ok := i.(ManualValidator)
	return v, ok
}

func validateParams(cmdName string, cmd command, params []string) ([]string, error) {
	paramsDefinitions := structListParamsKeys(cmd)
	var missing []string
	for n, isRequired := range paramsDefinitions {
		if isRequired && !contains(params, n) {
			missing = append(missing, n)
		}
	}

	var unexpected []string
	for _, p := range params {
		_, ok := paramsDefinitions[p]
		if !ok {
			unexpected = append(unexpected, p)
		}
	}

	if len(unexpected) > 0 {
		return missing, fmt.Errorf("%s: unexpected param(s) '%s'\n%s", cmdName, strings.Join(unexpected, "', '"), cmd.ParamsHelp())
	}
	return missing, nil
}

func generateParamsHelp(commandKey string, params map[string]bool) string {
	var buff bytes.Buffer
	var extra, required []string
	for n, isRequired := range params {
		if isRequired {
			required = append(required, n)
		} else {
			extra = append(extra, n)
		}
	}
	var anyRequired bool
	if len(required) > 0 {
		anyRequired = true
		buff.WriteString("\tRequired params:")
		for _, req := range required {
			buff.WriteString(fmt.Sprintf("\n\t\t- %s", req))
			if d, ok := awsdoc.TemplateParamsDoc(commandKey, req); ok {
				buff.WriteString(fmt.Sprintf(": %s", d))
			}
		}
	}

	if len(extra) > 0 {
		if anyRequired {
			buff.WriteString("\n\tExtra params:")
		} else {
			buff.WriteString("\n\tParams:")
		}
		for _, ext := range extra {
			buff.WriteString(fmt.Sprintf("\n\t\t- %s", ext))
			if d, ok := awsdoc.TemplateParamsDoc(commandKey, ext); ok {
				buff.WriteString(fmt.Sprintf(": %s", d))
			}
		}
	}
	return buff.String()
}

func fakeDryRunId(entity string) string {
	suffix := rand.Intn(1e6)
	switch entity {
	case cloud.Instance:
		return fmt.Sprintf("i-%d", suffix)
	case cloud.Subnet:
		return fmt.Sprintf("subnet-%d", suffix)
	case cloud.Vpc:
		return fmt.Sprintf("vpc-%d", suffix)
	case cloud.Volume:
		return fmt.Sprintf("vol-%d", suffix)
	case cloud.SecurityGroup:
		return fmt.Sprintf("sg-%d", suffix)
	case cloud.InternetGateway:
		return fmt.Sprintf("igw-%d", suffix)
	case cloud.NatGateway:
		return fmt.Sprintf("nat-%d", suffix)
	case cloud.RouteTable:
		return fmt.Sprintf("rtb-%d", suffix)
	default:
		return fmt.Sprintf("dryrunid-%d", suffix)
	}
}

type paramRule struct {
	errPrefix string
	tree      ruleNode
	extras    []string
}

func (p paramRule) help() string {
	if len(p.extras) == 0 {
		return p.tree.help()
	}

	return fmt.Sprintf("%s or extra params: \"%s\"", p.tree.help(), strings.Join(p.extras, "\", \""))
}

func (p paramRule) verify(keys []string) ([]string, error) {
	if p.tree == nil {
		return nil, nil
	}
	var unexpected []string
	for _, key := range keys {
		if p.tree.unexpected(key) && !contains(p.extras, key) {
			unexpected = append(unexpected, key)
		}
	}
	if len(unexpected) > 0 {
		return nil, fmt.Errorf("%sunexpected param(s) '%s': expecting %s", p.errPrefix, strings.Join(unexpected, "', '"), p.help())
	}
	missings, _, errs := p.tree.missings(keys)
	if len(errs) > 0 {
		var errStr bytes.Buffer
		for _, e := range errs {
			errStr.WriteString(e.Error())
		}
		return nil, errors.New(p.errPrefix + errStr.String())
	}
	return missings, nil
}

type ruleNode interface {
	help() string
	unexpected(string) bool
	missings([]string) ([]string, int, []error)
}

type oneOfNode struct {
	children []ruleNode
}

func (o oneOfNode) help() string {
	var childrenHint []string
	for _, c := range o.children {
		childrenHint = append(childrenHint, c.help())
	}
	return fmt.Sprintf("(%s)", strings.Join(childrenHint, " or "))
}

func (o oneOfNode) unexpected(s string) bool {
	for _, c := range o.children {
		if !c.unexpected(s) {
			return false
		}
	}
	return true
}

func (o oneOfNode) missings(keys []string) ([]string, int, []error) {
	var errs []error
	maxFound := -1
	var missings []string
	for _, child := range o.children {
		cMissings, nbFound, err := child.missings(keys)
		errs = append(errs, err...)
		if nbFound > maxFound {
			missings = cMissings
			maxFound = nbFound
		}
	}
	return missings, maxFound, nil
}

type oneOfNodeWithError struct {
	oneOfNode
}

func (o oneOfNodeWithError) missings(keys []string) (missings []string, found int, errs []error) {
	var hasFoundChild bool
	for _, child := range o.children {
		_, nbFound, _ := child.missings(keys)
		if nbFound > 0 {
			hasFoundChild = true
		}
	}
	missings, found, errs = o.oneOfNode.missings(keys)
	if !hasFoundChild {
		errs = append(errs, fmt.Errorf("expecting %s", o.help()))
	}
	return
}

type allOfNode struct {
	children []ruleNode
}

func (a allOfNode) help() string {
	var childrenHint []string
	for _, c := range a.children {
		childrenHint = append(childrenHint, c.help())
	}
	return fmt.Sprintf("(%s)", strings.Join(childrenHint, " and "))
}

func (a allOfNode) unexpected(s string) bool {
	for _, c := range a.children {
		if !c.unexpected(s) {
			return false
		}
	}
	return true
}

func (a allOfNode) missings(keys []string) (missings []string, nbFound int, errs []error) {
	for _, child := range a.children {
		cMissings, cFound, err := child.missings(keys)
		errs = append(errs, err...)
		if len(cMissings) > 0 {
			missings = append(missings, cMissings...)
		} else {
			nbFound += cFound
		}
	}
	return
}

type stringNode struct {
	key string
}

func (k stringNode) help() string {
	return fmt.Sprintf("\"%s\"", k.key)
}

func (k stringNode) unexpected(s string) bool {
	return k.key != s
}

func (k stringNode) missings(keys []string) (missings []string, nbFound int, errs []error) {
	if contains(keys, k.key) {
		nbFound++
		return
	}
	missings = append(missings, k.key)
	return
}

func oneOf(nodes ...ruleNode) ruleNode {
	return oneOfNode{children: nodes}
}

func oneOfE(nodes ...ruleNode) ruleNode {
	return oneOfNodeWithError{oneOfNode{children: nodes}}
}

func allOf(nodes ...ruleNode) ruleNode {
	return allOfNode{children: nodes}
}

func node(key string) ruleNode {
	return stringNode{key}
}

func String(v string) *string {
	return &v
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
