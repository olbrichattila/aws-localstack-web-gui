package controller

import (
	"context"
	"encoding/json"
	"net/url"
	"webuiApi/app/repositories/awsshared"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type topicRequest struct {
	Name string `json:"name"`
}

type topicMessageRequest struct {
	Message string `json:"message"`
}

type topicSubscribeRequest struct {
	Url string `json:"url"`
}

// SnsAction function can take any parameters defined in the Di config
func SNSGetAttributes(awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSNSClient()
	if err != nil {
		return "", err
	}

	topics, err := client.ListTopics(*ctx, &sns.ListTopicsInput{})
	if err != nil {
		return "", err
	}

	response := make([]map[string]string, len(topics.Topics))
	for i, topic := range topics.Topics {
		attrs, err := getSNSAttribute(client, ctx, *topic.TopicArn)
		if err != nil {
			return "", err
		}

		response[i] = attrs
	}

	res, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func getSNSAttribute(client *sns.Client, ctx *context.Context, topicUrl string) (map[string]string, error) {
	attrsOutput, err := client.GetTopicAttributes(*ctx, &sns.GetTopicAttributesInput{
		TopicArn: aws.String(topicUrl),
	})
	if err != nil {
		return nil, err
	}

	return attrsOutput.Attributes, nil
}

func SNSCreateTopic(req topicRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSNSClient()
	if err != nil {
		return "", err
	}

	_, err = client.CreateTopic(*ctx, &sns.CreateTopicInput{
		Name: &req.Name,
	})

	if err != nil {
		return "", err
	}

	return SNSGetAttributes(awsShared)
}

func SNSDeleteTopic(req topicRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSNSClient()
	if err != nil {
		return "", err
	}

	_, err = client.DeleteTopic(*ctx, &sns.DeleteTopicInput{
		TopicArn: aws.String(req.Name),
	})

	if err != nil {
		return "", err
	}

	return SNSGetAttributes(awsShared)
}

func SNSPublishToTopicARN(topicArn string, req topicMessageRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSNSClient()
	if err != nil {
		return "", err
	}

	_, err = client.Publish(*ctx, &sns.PublishInput{
		Message:  &req.Message,
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		return "", err
	}

	return "{}", nil
}

func SNSGetSubscriptionsByARN(topicArn string, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSNSClient()
	if err != nil {
		return "", err
	}

	topics, err := client.ListSubscriptionsByTopic(*ctx, &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		return "", err
	}

	res, err := json.Marshal(topics.Subscriptions)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func SNSCreateSubscriptionForARN(topicArn string, req topicSubscribeRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSNSClient()
	if err != nil {
		return "", err
	}

	parsedUrl, err := url.Parse(req.Url)
	if err != nil {
		return "", err
	}

	protocol := parsedUrl.Scheme
	_, err = client.Subscribe(*ctx, &sns.SubscribeInput{
		Protocol: aws.String(protocol),
		TopicArn: aws.String(topicArn),
		Endpoint: aws.String(req.Url),
	})

	if err != nil {
		return "", err
	}

	topics, err := client.ListSubscriptionsByTopic(*ctx, &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		return "", err
	}

	result, err := json.Marshal(topics.Subscriptions)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func SNSDeleteSubscriptionByARN(topicArn string, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSNSClient()
	if err != nil {
		return "", err
	}

	_, err = client.Unsubscribe(*ctx, &sns.UnsubscribeInput{
		SubscriptionArn: aws.String(topicArn),
	})

	return "{}", nil
}
