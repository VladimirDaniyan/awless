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
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/logger"
)

type AttachInstanceprofile struct {
	_        string `action:"attach" entity:"instanceprofile" awsAPI:"ec2"`
	logger   *logger.Logger
	api      ec2iface.EC2API
	Instance *string   `awsName:"InstanceId" awsType:"awsstr" templateName:"instance" required:""`
	Name     *string   `awsName:"IamInstanceProfile.Name" awsType:"awsstr" templateName:"name" required:""`
	Replace  *struct{} `templateName:"replace"`
}

type DetachInstanceprofile struct {
	_        string `action:"detach" entity:"instanceprofile" awsAPI:"ec2"`
	logger   *logger.Logger
	api      ec2iface.EC2API
	Instance *struct{} `templateName:"instance" required:""`
	Name     *struct{} `templateName:"name" required:""`
	Replace  *struct{} `templateName:"replace"`
}

type CreateInstanceprofile struct {
	_      string `action:"create" entity:"instanceprofile" awsAPI:"iam" awsCall:"CreateInstanceProfile" awsInput:"iam.CreateInstanceProfileInput" awsOutput:"iam.CreateInstanceProfileOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Name   *string `awsName:"InstanceProfileName" awsType:"awsstr" templateName:"name" required:""`
}

func (cmd *CreateInstanceprofile) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type DeleteInstanceprofile struct {
	_      string `action:"delete" entity:"instanceprofile" awsAPI:"iam" awsCall:"DeleteInstanceProfile" awsInput:"iam.DeleteInstanceProfileInput" awsOutput:"iam.DeleteInstanceProfileOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Name   *string `awsName:"InstanceProfileName" awsType:"awsstr" templateName:"name" required:""`
}

func (cmd *DeleteInstanceprofile) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}