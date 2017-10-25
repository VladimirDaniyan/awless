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
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/wallix/awless/logger"
)

type CreateRecord struct {
	_       string `action: "create" entity: "record" awsAPI: "route53"`
	logger  *logger.Logger
	api     route53iface.Route53API
	Zone    *struct{} `templateName: "zone" required: ""`
	Name    *struct{} `templateName: "name" required: ""`
	Type    *struct{} `templateName: "type" required: ""`
	Value   *struct{} `templateName: "value" required: ""`
	Ttl     *struct{} `templateName: "ttl" required: ""`
	Comment *struct{} `templateName: "comment"`
}
type DeleteRecord struct {
	_      string `action: "delete" entity: "record" awsAPI: "route53"`
	logger *logger.Logger
	api    route53iface.Route53API
	Zone   *struct{} `templateName: "zone" required: ""`
	Name   *struct{} `templateName: "name" required: ""`
	Type   *struct{} `templateName: "type" required: ""`
	Value  *struct{} `templateName: "value" required: ""`
	Ttl    *struct{} `templateName: "ttl" required: ""`
}
type UpdateRecord struct {
	_      string `action: "update" entity: "record" awsAPI: "route53"`
	logger *logger.Logger
	api    route53iface.Route53API
	Zone   *struct{} `templateName: "zone" required: ""`
	Name   *struct{} `templateName: "name" required: ""`
	Type   *struct{} `templateName: "type" required: ""`
	Value  *struct{} `templateName: "value" required: ""`
	Ttl    *struct{} `templateName: "ttl" required: ""`
}
