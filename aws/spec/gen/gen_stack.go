/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// DO NOT EDIT
// This file was automatically generated with go generate
package awsspec

import (
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/wallix/awless/logger"
)

type CreateStack struct {
	_               string `action: "create" entity: "stack" awsAPI: "cloudformation" awsCall: "CreateStack" awsInput: "CreateStackInput" awsOutput: "CreateStackOutput"`
	logger          *logger.Logger
	api             cloudformationiface.CloudFormationAPI
	Name            *string   `awsName: "StackName" awsType: "awsstr" templateName: "name" required: ""`
	TemplateFile    *string   `awsName: "TemplateBody" awsType: "awsfiletostring" templateName: "template-file" required: ""`
	Capabilities    *[]string `awsName: "Capabilities" awsType: "awsstringslice" templateName: "capabilities"`
	DisableRollback *bool     `awsName: "DisableRollback" awsType: "awsbool" templateName: "disable-rollback"`
	Notifications   *[]string `awsName: "NotificationARNs" awsType: "awsstringslice" templateName: "notifications"`
	OnFailure       *string   `awsName: "OnFailure" awsType: "awsstr" templateName: "on-failure"`
	Parameters      *struct{} `awsName: "Parameters" awsType: "awsparameterslice" templateName: "parameters"`
	ResourceTypes   *[]string `awsName: "ResourceTypes" awsType: "awsstringslice" templateName: "resource-types"`
	Role            *string   `awsName: "RoleARN" awsType: "awsstr" templateName: "role"`
	PolicyFile      *string   `awsName: "StackPolicyBody" awsType: "awsfiletostring" templateName: "policy-file"`
	Timeout         *int64    `awsName: "TimeoutInMinutes" awsType: "awsint64" templateName: "timeout"`
}
type UpdateStack struct {
	_                   string `action: "update" entity: "stack" awsAPI: "cloudformation" awsCall: "UpdateStack" awsInput: "UpdateStackInput" awsOutput: "UpdateStackOutput"`
	logger              *logger.Logger
	api                 cloudformationiface.CloudFormationAPI
	Name                *string   `awsName: "StackName" awsType: "awsstr" templateName: "name" required: ""`
	Capabilities        *[]string `awsName: "Capabilities" awsType: "awsstringslice" templateName: "capabilities"`
	Notifications       *[]string `awsName: "NotificationARNs" awsType: "awsstringslice" templateName: "notifications"`
	Parameters          *struct{} `awsName: "Parameters" awsType: "awsparameterslice" templateName: "parameters"`
	ResourceTypes       *[]string `awsName: "ResourceTypes" awsType: "awsstringslice" templateName: "resource-types"`
	Role                *string   `awsName: "RoleARN" awsType: "awsstr" templateName: "role"`
	PolicyFile          *string   `awsName: "StackPolicyBody" awsType: "awsfiletostring" templateName: "policy-file"`
	PolicyUpdateFile    *string   `awsName: "StackPolicyDuringUpdateBody" awsType: "awsfiletostring" templateName: "policy-update-file"`
	TemplateFile        *string   `awsName: "TemplateBody" awsType: "awsfiletostring" templateName: "template-file"`
	UsePreviousTemplate *bool     `awsName: "UsePreviousTemplate" awsType: "awsbool" templateName: "use-previous-template"`
}
type DeleteStack struct {
	_               string `action: "delete" entity: "stack" awsAPI: "cloudformation" awsCall: "DeleteStack" awsInput: "DeleteStackInput" awsOutput: "DeleteStackOutput"`
	logger          *logger.Logger
	api             cloudformationiface.CloudFormationAPI
	Name            *string   `awsName: "StackName" awsType: "awsstr" templateName: "name" required: ""`
	RetainResources *[]string `awsName: "RetainResources" awsType: "awsstringslice" templateName: "retain-resources"`
}
