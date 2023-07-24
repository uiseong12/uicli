package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

type LsFlag struct {
	All bool
}

func Ls(*cobra.Command, []string) {
	svc := ec2.New(
		session.New(),
		&aws.Config{
			Region: aws.String("ap-northeast-2"),
		},
	)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
				},
			},
		},
	}

	resp, err := svc.DescribeInstances(params)
	if err != nil {
		panic(err)
	}

	for _, res := range resp.Reservations {
		for _, inst := range res.Instances {
			fmt.Println(formatInstance(inst))
		}
	}
}

func formatInstance(inst *ec2.Instance) string {
	output := []string{
		*inst.PublicIpAddress,
		*inst.InstanceId,
		*inst.State.Name,
		(*inst.LaunchTime).Format("2006-01-02 15:04:05"),
		lookupTag(inst, "Name"),
	}
	return strings.Join(output[:], "\t")
}

func lookupTag(inst *ec2.Instance, key string) string {
	var value string
	for _, t := range inst.Tags {
		if *t.Key == key {
			value = *t.Value
			break
		}
	}
	return value
}
