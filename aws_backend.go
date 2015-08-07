package soko

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSBackend struct {
	SectionConfig

	client *ec2.EC2
}

func NewAWSBackend(config SectionConfig) (*AWSBackend, error) {
	conf := &aws.Config{Region: aws.String(config["region"])}
	cli := ec2.New(conf)

	return &AWSBackend{
		SectionConfig: config,
		client:        cli,
	}, nil
}
