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

type CreateAccesskey struct {
	_        string `action: "create" entity: "accesskey" awsAPI: "iam"`
	logger   *logger.Logger
	api      iamiface.IAMAPI
	User     *struct{} `templateName: "user" required: ""`
	NoPrompt *struct{} `templateName: "no-prompt"`
}
type DeleteAccesskey struct {
	_      string `action: "delete" entity: "accesskey" awsAPI: "iam" awsCall: "DeleteAccessKey" awsInput: "DeleteAccessKeyInput" awsOutput: "DeleteAccessKeyOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Id     *string `awsName: "AccessKeyId" awsType: "awsstr" templateName: "id" required: ""`
	User   *string `awsName: "UserName" awsType: "awsstr" templateName: "user"`
}
