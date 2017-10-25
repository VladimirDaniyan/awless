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
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/logger"
)

type CreatePolicy struct {
	_           string `action: "create" entity: "policy" awsAPI: "iam"`
	logger      *logger.Logger
	api         iamiface.IAMAPI
	Name        *string   `awsName: "PolicyName" awsType: "awsstr" templateName: "name" required: ""`
	Effect      *struct{} `templateName: "effect" required: ""`
	Action      *struct{} `templateName: "action" required: ""`
	Resource    *struct{} `templateName: "resource" required: ""`
	Description *string   `awsName: "Description" awsType: "awsstr" templateName: "description"`
	Conditions  *struct{} `templateName: "conditions"`
}
type UpdatePolicy struct {
	_          string `action: "update" entity: "policy" awsAPI: "iam"`
	logger     *logger.Logger
	api        iamiface.IAMAPI
	Arn        *struct{} `templateName: "arn" required: ""`
	Effect     *struct{} `templateName: "effect" required: ""`
	Action     *struct{} `templateName: "action" required: ""`
	Resource   *struct{} `templateName: "resource" required: ""`
	Conditions *struct{} `templateName: "conditions"`
}
type DeletePolicy struct {
	_           string `action: "delete" entity: "policy" awsAPI: "iam"`
	logger      *logger.Logger
	api         iamiface.IAMAPI
	Arn         *string   `awsName: "PolicyArn" awsType: "awsstr" templateName: "arn" required: ""`
	AllVersions *struct{} `templateName: "all-versions"`
}
type AttachPolicy struct {
	_       string `action: "attach" entity: "policy" awsAPI: "iam"`
	logger  *logger.Logger
	api     iamiface.IAMAPI
	Arn     *struct{} `templateName: "arn"`
	Service *struct{} `templateName: "service"`
	Access  *struct{} `templateName: "access"`
	User    *struct{} `templateName: "user"`
	Group   *struct{} `templateName: "group"`
	Role    *struct{} `templateName: "role"`
}
type DetachPolicy struct {
	_       string `action: "detach" entity: "policy" awsAPI: "iam"`
	logger  *logger.Logger
	api     iamiface.IAMAPI
	Arn     *struct{} `templateName: "arn"`
	Service *struct{} `templateName: "service"`
	Access  *struct{} `templateName: "access"`
	User    *struct{} `templateName: "user"`
	Group   *struct{} `templateName: "group"`
	Role    *struct{} `templateName: "role"`
}
