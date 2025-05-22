package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"webuiApi/app/repositories/awsshared"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/olbrichattila/gofra/pkg/app/gofraerror"
)

type SQSController struct {
	client *sqs.Client
	ctx    *context.Context
}

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

type sqsSendFIFOMessageRequest struct {
	MessageGroupId         string `json:"messageGroupId"`
	MessageDeduplicationId string `json:"messageDeduplicationId"`
	QueueUrl               string `json:"queueUrl"`
	MessageBody            string `json:"messageBody"`
}

type sqsReadMessageRequest struct {
	QueueUrl            string `json:"queueUrl"`
	MaxNumberOfMessages int    `json:"maxNumberOfMessages"`
}

func (s *SQSController) Before(awsShared awsshared.AWSShared) error {
	var err error
	s.client, s.ctx, err = awsShared.GetSQSClient()

	if err != nil {
		return gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (s *SQSController) SQSGetAttributesAction() (string, error) {
	listQueuesOutput, err := s.client.ListQueues(*s.ctx, &sqs.ListQueuesInput{})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	var result []map[string]interface{}

	for _, queueUrl := range listQueuesOutput.QueueUrls {
		// Get attributes for each queue
		attrsOutput, err := s.client.GetQueueAttributes(*s.ctx, &sqs.GetQueueAttributesInput{
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
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(res), nil
}

func (s *SQSController) SQSCreateQueueAction(req sqsAddQueueRequest) (string, error) {
	_, err := s.client.CreateQueue(*s.ctx, &sqs.CreateQueueInput{
		QueueName: &req.QueueName,
		Attributes: map[string]string{
			string(types.QueueAttributeNameVisibilityTimeout):  strconv.Itoa(req.DelaySeconds),
			string(types.QueueAttributeNameMaximumMessageSize): strconv.Itoa(req.MaximumMessageSize),
		},
	})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusBadRequest)
	}

	return s.SQSGetAttributesAction()
}

func (s *SQSController) SQSCreateFIFOQueueAction(req sqsAddQueueRequest) (string, error) {
	_, err := s.client.CreateQueue(*s.ctx, &sqs.CreateQueueInput{
		QueueName: &req.QueueName,
		Attributes: map[string]string{
			"FifoQueue":                 "true",
			"ContentBasedDeduplication": "true",              // or set to "false" and provide MessageDeduplicationId manually
			"DeduplicationScope":        "messageGroup",      // optional: messageGroup or queue
			"FifoThroughputLimit":       "perMessageGroupId", // optional: perQueue or perMessageGroupId
			string(types.QueueAttributeNameVisibilityTimeout):  strconv.Itoa(req.DelaySeconds),
			string(types.QueueAttributeNameMaximumMessageSize): strconv.Itoa(req.MaximumMessageSize),
		},
	})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusBadRequest)
	}

	return s.SQSGetAttributesAction()
}

func (s *SQSController) SQSDeleteQueueAction(req sqsQueueRequest) (string, error) {

	_, err := s.client.DeleteQueue(*s.ctx, &sqs.DeleteQueueInput{
		QueueUrl: &req.QueueUrl,
	})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return s.SQSGetAttributesAction()
}

func (s *SQSController) SQSGetAttributeAction(req sqsQueueRequest) (string, error) {
	attrsOutput, err := s.client.GetQueueAttributes(*s.ctx, &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(req.QueueUrl),
		AttributeNames: []types.QueueAttributeName{
			types.QueueAttributeNameAll,
		},
	})

	res, err := json.Marshal(attrsOutput.Attributes)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(res), nil
}

func (s *SQSController) SQSPurgeQueueAction(req sqsQueueRequest) (string, error) {
	_, err := s.client.PurgeQueue(*s.ctx, &sqs.PurgeQueueInput{
		QueueUrl: &req.QueueUrl,
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return s.SQSGetAttributesAction()
}

func (s *SQSController) SQSendMessageAction(req sqsSendMessageRequest) (string, error) {
	delay, err := strconv.Atoi(req.DelaySeconds)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	_, err = s.client.SendMessage(*s.ctx, &sqs.SendMessageInput{
		DelaySeconds: int32(delay),
		MessageBody:  &req.MessageBody,
		QueueUrl:     &req.QueueUrl,
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return "{}", nil
}

func (s *SQSController) SQSendFIFOMessageAction(req sqsSendFIFOMessageRequest) (string, error) {
	_, err := s.client.SendMessage(*s.ctx, &sqs.SendMessageInput{
		MessageGroupId:         &req.MessageGroupId,
		MessageDeduplicationId: &req.MessageDeduplicationId,
		MessageBody:            &req.MessageBody,
		QueueUrl:               &req.QueueUrl,
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return "{}", nil
}

func (s *SQSController) SQSReceiveMessages(req sqsReadMessageRequest) (string, error) {
	messages, err := s.client.ReceiveMessage(*s.ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &req.QueueUrl,
		MaxNumberOfMessages: int32(req.MaxNumberOfMessages),
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	if messages.Messages == nil {
		return "[]", nil
	}

	res, err := json.Marshal(messages.Messages)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(res), nil
}
