package controller

import (
	"encoding/json"
	"log"
	"strconv"
	"webuiApi/app/repositories/awsshared"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type sqsAddQueueRequest struct {
	QueueName          string `json:"name"`
	DelaySeconds       int    `json:"delaySeconds"`
	MaximumMessageSize int    `json:"maximumMessageSize"`
}

type sqsQueueRequest struct {
	QueueUrl string `json:"queueUrl"`
}

type sqsSendMessageRequest struct {
	QueueUrl     string `json:"queueUrl"`
	DelaySeconds string `json:"delaySeconds"`
	MessageBody  string `json:"messageBody"`
}

type sqsReadMessageRequest struct {
	QueueUrl            string `json:"queueUrl"`
	MaxNumberOfMessages int    `json:"maxNumberOfMessages"`
}

func SQSGetAttributesAction(awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSQSClient()
	if err != nil {
		return "", err
	}

	// List queues
	listQueuesOutput, err := client.ListQueues(*ctx, &sqs.ListQueuesInput{})
	if err != nil {
		return "", err
	}

	var result []map[string]interface{}

	for _, queueUrl := range listQueuesOutput.QueueUrls {
		// Get attributes for each queue
		attrsOutput, err := client.GetQueueAttributes(*ctx, &sqs.GetQueueAttributesInput{
			QueueUrl: aws.String(queueUrl),
			AttributeNames: []types.QueueAttributeName{
				types.QueueAttributeNameAll,
			},
		})
		if err != nil {
			log.Printf("failed to get attributes for queue %s: %v", queueUrl, err)
			continue
		}

		// Build result entry
		entry := map[string]interface{}{
			"url":        queueUrl,
			"attributes": attrsOutput.Attributes,
		}

		result = append(result, entry)
	}

	if result == nil {
		result = []map[string]any{}
	}

	res, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func SQSCreateQueueAction(req sqsAddQueueRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSQSClient()
	if err != nil {
		return "", err
	}

	_, err = client.CreateQueue(*ctx, &sqs.CreateQueueInput{
		QueueName: &req.QueueName,
		Attributes: map[string]string{
			string(types.QueueAttributeNameVisibilityTimeout):  strconv.Itoa(req.DelaySeconds),
			string(types.QueueAttributeNameMaximumMessageSize): strconv.Itoa(req.MaximumMessageSize),
		},
	})
	if err != nil {
		return "", nil
	}

	return SQSGetAttributesAction(awsShared)
}

func SQSDeleteQueueAction(req sqsQueueRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSQSClient()
	if err != nil {
		return "", err
	}

	_, err = client.DeleteQueue(*ctx, &sqs.DeleteQueueInput{
		QueueUrl: &req.QueueUrl,
	})
	if err != nil {
		return "", err
	}

	return SQSGetAttributesAction(awsShared)
}

func SQSGetAttributeAction(req sqsQueueRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSQSClient()
	if err != nil {
		return "", err
	}

	attrsOutput, err := client.GetQueueAttributes(*ctx, &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(req.QueueUrl),
		AttributeNames: []types.QueueAttributeName{
			types.QueueAttributeNameAll,
		},
	})

	res, err := json.Marshal(attrsOutput.Attributes)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func SQSPurgeQueueAction(req sqsQueueRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSQSClient()
	if err != nil {
		return "", err
	}

	_, err = client.PurgeQueue(*ctx, &sqs.PurgeQueueInput{
		QueueUrl: &req.QueueUrl,
	})

	if err != nil {
		return "", err
	}

	return SQSGetAttributesAction(awsShared)
}

func SQSendMessageAction(req sqsSendMessageRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSQSClient()
	if err != nil {
		return "", err
	}

	delay, err := strconv.Atoi(req.DelaySeconds)
	if err != nil {
		return "", err
	}

	_, err = client.SendMessage(*ctx, &sqs.SendMessageInput{
		DelaySeconds: int32(delay),
		MessageBody:  &req.MessageBody,
		QueueUrl:     &req.QueueUrl,
	})

	if err != nil {
		return "", err
	}

	return "{}", nil
}

func SQSReceiveMessages(req sqsReadMessageRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetSQSClient()
	if err != nil {
		return "", nil
	}

	messages, err := client.ReceiveMessage(*ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &req.QueueUrl,
		MaxNumberOfMessages: int32(req.MaxNumberOfMessages),
	})

	if err != nil {
		return "", nil
	}

	if messages.Messages == nil {
		return "[]", nil
	}

	res, err := json.Marshal(messages.Messages)
	if err != nil {
		return "", nil
	}

	return string(res), nil
}
