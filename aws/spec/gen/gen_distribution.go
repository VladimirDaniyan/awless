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
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/wallix/awless/logger"
)

type CreateDistribution struct {
	_              string `action: "create" entity: "distribution" awsAPI: "cloudfront"`
	logger         *logger.Logger
	api            cloudfrontiface.CloudFrontAPI
	OriginDomain   *struct{} `templateName: "origin-domain" required: ""`
	Certificate    *struct{} `templateName: "certificate"`
	Comment        *struct{} `templateName: "comment"`
	DefaultFile    *struct{} `templateName: "default-file"`
	DomainAliases  *struct{} `templateName: "domain-aliases"`
	Enable         *struct{} `templateName: "enable"`
	ForwardCookies *struct{} `templateName: "forward-cookies"`
	ForwardQueries *struct{} `templateName: "forward-queries"`
	HttpsBehaviour *struct{} `templateName: "https-behaviour"`
	OriginPath     *struct{} `templateName: "origin-path"`
	PriceClass     *struct{} `templateName: "price-class"`
	MinTtl         *struct{} `templateName: "min-ttl"`
}
type CheckDistribution struct {
	_       string `action: "check" entity: "distribution" awsAPI: "cloudfront"`
	logger  *logger.Logger
	api     cloudfrontiface.CloudFrontAPI
	Id      *struct{} `templateName: "id" required: ""`
	State   *struct{} `templateName: "state" required: ""`
	Timeout *struct{} `templateName: "timeout" required: ""`
}
type UpdateDistribution struct {
	_      string `action: "update" entity: "distribution" awsAPI: "cloudfront"`
	logger *logger.Logger
	api    cloudfrontiface.CloudFrontAPI
	Id     *string `awsName: "Id" awsType: "awsstr" templateName: "id" required: ""`
	Enable *bool   `awsName: "DistributionConfig.Enabled" awsType: "awsbool" templateName: "enable" required: ""`
}
type DeleteDistribution struct {
	_      string `action: "delete" entity: "distribution" awsAPI: "cloudfront"`
	logger *logger.Logger
	api    cloudfrontiface.CloudFrontAPI
	Id     *string `awsName: "Id" awsType: "awsstr" templateName: "id" required: ""`
}
