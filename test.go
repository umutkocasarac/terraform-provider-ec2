package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
)

type apiClient struct {
	client *ec2.Client
}

func main2() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-west-2"),
		config.WithSharedConfigProfile("asd"),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client := ec2.NewFromConfig(cfg)
	opts := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-e7527ed7"),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     1,
		MaxCount:     1,
	}
	reservation, err := client.RunInstances(context.TODO(), opts)
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	instanceID := *reservation.Instances[0].InstanceId
	println("Instance created %v", instanceID)
	println(resourceScaffoldingRead(client, instanceID).Reservations[0].Instances)
	println("Instance readed")
	println(resourceScaffoldingDelete(client, instanceID).TerminatingInstances[0].CurrentState.Name)
	println("Instance deleted")
}

func resourceScaffoldingRead(client *ec2.Client, instanceID string) *ec2.DescribeInstancesOutput {
	instanceIDs := make([]string, 1)
	instanceIDs[0] = instanceID
	input := &ec2.DescribeInstancesInput{
		InstanceIds: instanceIDs,
	}
	result, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		panic("can not describe instance  error, " + err.Error())
	}
	return result
}

func resourceScaffoldingDelete(client *ec2.Client, instanceID string) *ec2.TerminateInstancesOutput {
	instanceIDs := make([]string, 1)
	instanceIDs[0] = instanceID
	deleteRequest := &ec2.TerminateInstancesInput{
		InstanceIds: instanceIDs,
	}
	result, err := client.TerminateInstances(context.TODO(), deleteRequest)
	if err != nil {
		panic("can not terminate instance instance  error, " + err.Error())
	}
	return result
}
