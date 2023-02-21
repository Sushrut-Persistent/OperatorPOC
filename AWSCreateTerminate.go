package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/credentials"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2CreateInstanceAPI interface {
	RunInstances(ctx context.Context,
		params *ec2.RunInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)

	CreateTags(ctx context.Context,
		params *ec2.CreateTagsInput,
		optFns ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error)
}

func CreateInstance(c context.Context, api EC2CreateInstanceAPI, input *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error) {
	return api.RunInstances(c, input)
}

func SetTags(c context.Context, api EC2CreateInstanceAPI, input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	return api.CreateTags(c, input)
}

func CreateInstanceCmd(name string, value string) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-0b5eea76982371e91"),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	}

	result, err := CreateInstance(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error creating an instance:")
		fmt.Println(err)
		return
	}

	tagInput := &ec2.CreateTagsInput{
		Resources: []string{*result.Instances[0].InstanceId},
		Tags: []types.Tag{
			{
				Key:   aws.String(name),
				Value: aws.String(value),
			},
		},
	}

	_, err = SetTags(context.TODO(), client, tagInput)
	if err != nil {
		fmt.Println("Got an error tagging the instance:")
		fmt.Println(err)
		return
	}

	fmt.Println("Created tagged instance with ID " + *result.Instances[0].InstanceId)
}

type EC2TerminateInstanceAPI interface {
	TerminateInstances(ctx context.Context,
		params *ec2.TerminateInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error)
}

func DeleteInstance(c context.Context, api EC2TerminateInstanceAPI, input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	return api.TerminateInstances(c, input)
}

func DeleteInstanceCmd(instanceId string) {

	// cfg, err := config.LoadDefaultConfig(context.TODO(), &aws.Config{
	// 	Credentials: aws.Credentials{
	// 		AccessKeyID:     "AKIAWT4IN5FUAI54GCMK",
	// 		SecretAccessKey: "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",
	// 	},
	// 	Region: "us-east-1",
	// })

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(options *config.LoadOptions) error {
		options.Credentials = aws.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "AKIAWT4IN5FUAI54GCMK",
				SecretAccessKey: "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",
			},
		}
		options.Region = "us-east-1"
		return nil
	})
	
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)

	instances := strings.Split(instanceId, ",")

	input := &ec2.TerminateInstancesInput{
		InstanceIds: instances,
		DryRun:      new(bool),
	}

	result, err := DeleteInstance(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error terminating the instance:")
		fmt.Println(err)
		return
	}

	fmt.Println("Terminated instance with id: ", *result.TerminatingInstances[0].InstanceId)
}

func main() {

	choice := flag.String("c", "", "The choice that user wants to make")
	name := flag.String("n", "", "The name of the tag to attach to the instance")
	value := flag.String("v", "", "The value of the tag to attach to the instance")
	instanceId := flag.String("i", "", "The IDs of the instance to terminate")
	flag.Parse()

	if *choice == "terminate" {

		if *instanceId == "" {
			fmt.Println("You must supply an instance ID (-i InstanceID")
			return
		}
		DeleteInstanceCmd(*instanceId)
	} else if *choice == "create" {

		if *name == "" || *value == "" {
			fmt.Println("You must supply a name and value for the tag (-n TagName -v TagValue)")
			return
		}
		CreateInstanceCmd(*name, *value)
	} else {

		fmt.Println("You must supply a proper action to perform (-c create/terminate)")
		return
	}
}

// Wait for 5 minutes
// time.Sleep(5 * time.Minute)