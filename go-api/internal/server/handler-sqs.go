package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func (s *server) getSqsListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.getSqsByQueueUrlHandler(w, r)
		return
	}

	client, ctx, err := s.awsShared.GetSQSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// List queues
	listQueuesOutput, err := client.ListQueues(*ctx, &sqs.ListQueuesInput{})
	if err != nil {
		log.Fatalf("failed to list queues: %v", err)
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
		result = []map[string]interface{}{}
	}

	s.respondAny(w, r, result)
}

func (s *server) getSqsByQueueUrlHandler(w http.ResponseWriter, r *http.Request) {
	var req sqsQueueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSQSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	attrsOutput, err := client.GetQueueAttributes(*ctx, &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(req.QueueUrl),
		AttributeNames: []types.QueueAttributeName{
			types.QueueAttributeNameAll,
		},
	})

	s.respondAny(w, r, attrsOutput.Attributes)
}

func (s *server) addSqsQueueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		s.deleteSqsQueueHandler(w, r)
		return
	}

	var req sqsAddQueueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSQSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = client.CreateQueue(*ctx, &sqs.CreateQueueInput{
		QueueName: &req.QueueName,
		Attributes: map[string]string{
			string(types.QueueAttributeNameVisibilityTimeout):  strconv.Itoa(req.DelaySeconds),
			string(types.QueueAttributeNameMaximumMessageSize): strconv.Itoa(req.MaximumMessageSize),
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.getSqsListHandler(w, r)
}

func (s *server) deleteSqsQueueHandler(w http.ResponseWriter, r *http.Request) {
	var req sqsQueueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSQSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = client.DeleteQueue(*ctx, &sqs.DeleteQueueInput{
		QueueUrl: &req.QueueUrl,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.getSqsListHandler(w, r)
}

func (s *server) purgeSqsQueueHandler(w http.ResponseWriter, r *http.Request) {
	var req sqsQueueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSQSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = client.PurgeQueue(*ctx, &sqs.PurgeQueueInput{
		QueueUrl: &req.QueueUrl,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.getSqsListHandler(w, r)
}

func (s *server) sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req sqsSendMessageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("???")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSQSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delay, err := strconv.Atoi(req.DelaySeconds)
	if err != nil {
		http.Error(w, "Invalid delay seconds", http.StatusBadRequest)
		return
	}

	_, err = client.SendMessage(*ctx, &sqs.SendMessageInput{
		DelaySeconds: int32(delay),
		MessageBody:  &req.MessageBody,
		QueueUrl:     &req.QueueUrl,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{}"))
}

func (s *server) getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var req sqsReadMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSQSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messages, err := client.ReceiveMessage(*ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &req.QueueUrl,
		MaxNumberOfMessages: int32(req.MaxNumberOfMessages),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if messages.Messages == nil {
		w.Write([]byte("[]"))
		return
	}

	s.respondAny(w, r, messages.Messages)
}

func (s *server) respondAny(w http.ResponseWriter, r *http.Request, toMarshal any) {
	data, err := json.Marshal(toMarshal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
