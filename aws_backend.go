package soko

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSBackend struct {
	SectionConfig

	client *ec2.EC2
}

const awsTomlTemplate = `[default]
backend = "aws"

[aws]
access_key_id = "%s"
secret_access_key = "%s"
region = "%s"
`

func NewAWSBackend(config SectionConfig) (*AWSBackend, error) {
	cred := credentials.NewStaticCredentials(config["access_key_id"], config["secret_access_key"], "")
	conf := &aws.Config{
		Credentials: cred,
		Region:      aws.String(config["region"]),
	}
	cli := ec2.New(conf)

	return &AWSBackend{
		SectionConfig: config,
		client:        cli,
	}, nil
}

func (b *AWSBackend) Save() error {
	config := b.SectionConfig
	data := fmt.Sprintf(
		awsTomlTemplate,
		config["access_key_id"],
		config["secret_access_key"],
		config["region"],
	)
	return ioutil.WriteFile(defaultConfigPath, []byte(data), 0644)
}

func (b *AWSBackend) Get(serverID string, key string) (string, error) {
	return b.get(serverID, key, false)
}

func (b *AWSBackend) get(serverID string, key string, isInternal bool) (string, error) {
	params := &ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("resource-id"),
				Values: []*string{aws.String(serverID)},
			},
			{
				Name:   aws.String("key"),
				Values: []*string{aws.String(key)},
			},
		},
	}

	tags, err := b.client.DescribeTags(params)
	if err != nil {
		return "", err
	}

	switch s := len(tags.Tags); s {
	case 0:
		if !isInternal {
			sayEmpty(key)
		}
		return "", nil
	case 1:
		return *tags.Tags[0].Value, nil
	default:
		return "", fmt.Errorf("Invalid size of key %s: %d tags exist", key, s)
	}

}

func (b *AWSBackend) Put(serverID string, key string, value string) error {
	params := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(serverID),
		},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	}
	_, err := b.client.CreateTags(params)
	return err
}

func (b *AWSBackend) Delete(serverID string, key string) error {
	currentV, _ := b.get(serverID, key, true)

	params := &ec2.DeleteTagsInput{
		Resources: []*string{
			aws.String(serverID),
		},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(currentV),
			},
		},
	}
	_, err := b.client.DeleteTags(params)
	return err
}
