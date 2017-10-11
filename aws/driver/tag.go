package awsdriver

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateTag struct {
	_        string `awsAPI:"ec2"` //  awsCall:"CreateTags" awsInput:"ec2.CreateTagsInput" awsOutput:"ec2.CreateTagsOutput"
	result   string
	logger   *logger.Logger
	api      ec2iface.EC2API
	sess     *session.Session
	Resource *string `awsName:"Resources" awsType:"awsstringslice" templateName:"resource" required:""`
	Key      *string `templateName:"key" required:""`
	Value    *string `templateName:"value" required:""`
}

func (cmd *CreateTag) Inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func (cmd *CreateTag) Validate() error {
	return validateStruct(cmd)
}

func (cmd *CreateTag) Run() (interface{}, error) {
	input := &ec2.CreateTagsInput{}
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("CreateTag: cannot inject in ec2.CreateTagsInput: %s", err)
	}
	input.Tags = []*ec2.Tag{{Key: cmd.Key, Value: cmd.Value}}

	start := time.Now()
	req, _ := cmd.api.CreateTagsRequest(input)
	req.Retryer = createTagRetryer{}
	if err := req.Send(); err != nil {
		return nil, fmt.Errorf("CreateTag: %s", err)
	}
	cmd.logger.ExtraVerbosef("ec2.CreateTags call took %s", time.Since(start))
	return cmd.result, nil
}

/*
type createTagRetryer struct {
	client.DefaultRetryer
}

func (d createTagRetryer) MaxRetries() int { return 5 }
func (d createTagRetryer) ShouldRetry(r *request.Request) bool {
	if d.DefaultRetryer.ShouldRetry(r) || !(r.HTTPResponse.StatusCode < 300 && r.HTTPResponse.StatusCode >= 200) {
		return true
	}

	return false
}*/
