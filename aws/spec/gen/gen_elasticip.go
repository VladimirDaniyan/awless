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
	"github.com/wallix/awless/logger"
)

type CreateElasticip struct {
	_      string `action: "create" entity: "elasticip" awsAPI: "ec2" awsCall: "AllocateAddress" awsInput: "AllocateAddressInput" awsOutput: "AllocateAddressOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Domain *string `awsName: "Domain" awsType: "awsstr" templateName: "domain" required: ""`
}
type DeleteElasticip struct {
	_      string `action: "delete" entity: "elasticip" awsAPI: "ec2" awsCall: "ReleaseAddress" awsInput: "ReleaseAddressInput" awsOutput: "ReleaseAddressOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "AllocationId" awsType: "awsstr" templateName: "id"`
	Ip     *string `awsName: "PublicIp" awsType: "awsstr" templateName: "ip"`
}
type AttachElasticip struct {
	_                  string `action: "attach" entity: "elasticip" awsAPI: "ec2" awsCall: "AssociateAddress" awsInput: "AssociateAddressInput" awsOutput: "AssociateAddressOutput" awsDryRun: ""`
	logger             *logger.Logger
	api                ec2iface.EC2API
	Id                 *string `awsName: "AllocationId" awsType: "awsstr" templateName: "id" required: ""`
	Instance           *string `awsName: "InstanceId" awsType: "awsstr" templateName: "instance"`
	Networkinterface   *string `awsName: "NetworkInterfaceId" awsType: "awsstr" templateName: "networkinterface"`
	Privateip          *string `awsName: "PrivateIpAddress" awsType: "awsstr" templateName: "privateip"`
	AllowReassociation *bool   `awsName: "AllowReassociation" awsType: "awsbool" templateName: "allow-reassociation"`
}
type DetachElasticip struct {
	_           string `action: "detach" entity: "elasticip" awsAPI: "ec2" awsCall: "DisassociateAddress" awsInput: "DisassociateAddressInput" awsOutput: "DisassociateAddressOutput" awsDryRun: ""`
	logger      *logger.Logger
	api         ec2iface.EC2API
	Association *string `awsName: "AssociationId" awsType: "awsstr" templateName: "association" required: ""`
}
