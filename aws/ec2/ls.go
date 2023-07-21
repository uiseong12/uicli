package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

func Ls(*cobra.Command, []string) {
	svc := ec2.New(
		session.New(),
		&aws.Config{
			Region: aws.String("ap-northeast-2"),
		},
	)
	param := &ec2.DescribeInstanceAttributeInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
				},
			},
		},
	}
}
