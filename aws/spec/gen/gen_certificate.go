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
	"github.com/aws/aws-sdk-go/service/acm/acmiface"
	"github.com/wallix/awless/logger"
)

type CreateCertificate struct {
	_                 string `action: "create" entity: "certificate" awsAPI: "acm"`
	logger            *logger.Logger
	api               acmiface.ACMAPI
	Domains           *struct{} `templateName: "domains" required: ""`
	ValidationDomains *struct{} `templateName: "validation-domains"`
}
type DeleteCertificate struct {
	_      string `action: "delete" entity: "certificate" awsAPI: "acm" awsCall: "DeleteCertificate" awsInput: "DeleteCertificateInput" awsOutput: "DeleteCertificateOutput"`
	logger *logger.Logger
	api    acmiface.ACMAPI
	Arn    *string `awsName: "CertificateArn" awsType: "awsstr" templateName: "arn" required: ""`
}
type CheckCertificate struct {
	_       string `action: "check" entity: "certificate" awsAPI: "acm"`
	logger  *logger.Logger
	api     acmiface.ACMAPI
	Arn     *struct{} `templateName: "arn" required: ""`
	State   *struct{} `templateName: "state" required: ""`
	Timeout *struct{} `templateName: "timeout" required: ""`
}
