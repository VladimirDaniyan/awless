package awsat

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func TestInstance(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		userdataFile := generateTmpFile("this is my content with {{ .Variables.oneRef }} content")
		defer os.Remove(userdataFile)

		Template("oneRef=awesome\n"+
			"create instance count=3 image=ami-1234 "+
			"name=myinstance subnet=sub_1 type=t2.nano keypair=mykp ip=10.2.3.4 "+
			"userdata="+userdataFile+" securitygroup=sg-1234 lock=true role=myrole").
			Mock(&ec2Mock{
				RunInstancesFunc: func(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
					return &ec2.Reservation{Instances: []*ec2.Instance{{InstanceId: String("new-instance-id")}}}, nil
				},
				CreateTagsRequestFunc: func(input *ec2.CreateTagsInput) (req *request.Request, output *ec2.CreateTagsOutput) {
					output = &ec2.CreateTagsOutput{}
					req = request.New(aws.Config{}, metadata.ClientInfo{}, request.Handlers{}, nil, &request.Operation{}, input, output)
					return
				},
			}).ExpectInput("RunInstances", &ec2.RunInstancesInput{
			SubnetId:              String("sub_1"),
			ImageId:               String("ami-1234"),
			InstanceType:          String("t2.nano"),
			MinCount:              Int64(3),
			MaxCount:              Int64(3),
			KeyName:               String("mykp"),
			PrivateIpAddress:      String("10.2.3.4"),
			SecurityGroupIds:      []*string{String("sg-1234")},
			DisableApiTermination: Bool(true),
			IamInstanceProfile:    &ec2.IamInstanceProfileSpecification{Name: String("myrole")},
			UserData:              String(base64.StdEncoding.EncodeToString([]byte("this is my content with awesome content"))),
		}).ExpectInput("CreateTagsRequest", &ec2.CreateTagsInput{
			Resources: []*string{String("new-instance-id")},
			Tags: []*ec2.Tag{
				{Key: String("Name"), Value: String("myinstance")},
			},
		}).ExpectCommandResult("new-instance-id").ExpectCalls("RunInstances", "CreateTagsRequest").Run(t)
	})

	t.Run("delete", func(t *testing.T) {
		Template("delete instance id=id-1234").Mock(&ec2Mock{
			TerminateInstancesFunc: func(param0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) { return nil, nil },
		}).ExpectInput("TerminateInstances", &ec2.TerminateInstancesInput{InstanceIds: []*string{String("id-1234")}}).
			ExpectCalls("TerminateInstances").Run(t)
		Template("delete instance ids=id-1234,id-2345").Mock(&ec2Mock{
			TerminateInstancesFunc: func(param0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) { return nil, nil },
		}).ExpectInput("TerminateInstances", &ec2.TerminateInstancesInput{InstanceIds: []*string{String("id-1234"), String("id-2345")}}).
			ExpectCalls("TerminateInstances").Run(t)
	})

	t.Run("check", func(t *testing.T) {
		Template("check instance id=id-1234 state=running timeout=0").Mock(&ec2Mock{
			DescribeInstancesFunc: func(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
				return &ec2.DescribeInstancesOutput{Reservations: []*ec2.Reservation{
					{Instances: []*ec2.Instance{{InstanceId: input.InstanceIds[0], State: &ec2.InstanceState{Name: String("running")}}}},
				}}, nil
			}}).ExpectInput("DescribeInstances", &ec2.DescribeInstancesInput{InstanceIds: []*string{String("id-1234")}}).
			ExpectCalls("DescribeInstances").Run(t)
	})
}

func generateTmpFile(content string) (path string) {
	file, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(file.Name(), []byte(content), 0644)
	return file.Name()
}