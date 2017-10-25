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

type CreateGroup struct {
	_      string `action: "create" entity: "group" awsAPI: "iam" awsCall: "CreateGroup" awsInput: "CreateGroupInput" awsOutput: "CreateGroupOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Name   *string `awsName: "GroupName" awsType: "awsstr" templateName: "name" required: ""`
}
type DeleteGroup struct {
	_      string `action: "delete" entity: "group" awsAPI: "iam" awsCall: "DeleteGroup" awsInput: "DeleteGroupInput" awsOutput: "DeleteGroupOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Name   *string `awsName: "GroupName" awsType: "awsstr" templateName: "name" required: ""`
}
