package server

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

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

func (s *server) getSNSAttributes(w http.ResponseWriter, r *http.Request) {
	client, ctx, err := s.awsShared.GetSNSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	topics, err := client.ListTopics(*ctx, &sns.ListTopicsInput{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]map[string]string, len(topics.Topics))
	for i, topic := range topics.Topics {
		attrs, err := s.getSNSAttribute(*topic.TopicArn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response[i] = attrs
	}

	s.respondAny(w, r, response)
}

func (s *server) newTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		s.deleteTopicHandler(w, r)
		return
	}

	var req topicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSNSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = client.CreateTopic(*ctx, &sns.CreateTopicInput{
		Name: &req.Name,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.getSNSAttributes(w, r)
}

func (s *server) deleteTopicHandler(w http.ResponseWriter, r *http.Request) {
	var req topicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSNSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = client.DeleteTopic(*ctx, &sns.DeleteTopicInput{
		TopicArn: aws.String(req.Name),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.getSNSAttributes(w, r)
}

// this is a group handler responsible to route all api/sns/sub/{topicName}/{requests}
func (s *server) topicApiHandler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) == 3 {
		s.newTopicHandler(w, r)
		return
	}

	rootCategory := urlParts[3]
	if rootCategory == "sub" {
		if len(urlParts) == 5 {
			s.snsSubscriptionsHandler(w, r, urlParts[4])
			return
		}

		if len(urlParts) < 6 {
			http.Error(w, "Invalid SNS URL", http.StatusBadRequest)
			return
		}
		topicArn := urlParts[4]
		route := urlParts[5]
		switch route {
		case "publish":
			s.topicPublishHandler(w, r, topicArn)
		default:
			http.Error(w, "Invalid route", http.StatusBadRequest)
			return
		}

		return
	}

	http.Error(w, "", http.StatusNotFound)
}

func (s *server) topicPublishHandler(w http.ResponseWriter, r *http.Request, topicArn string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	var req topicMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSNSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = client.Publish(*ctx, &sns.PublishInput{
		Message:  &req.Message,
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("{}"))
}

func (s *server) snsSubscriptionsHandler(w http.ResponseWriter, r *http.Request, topicArn string) {
	if r.Method == http.MethodPost {
		s.snsSubscribeHandler(w, r, topicArn)
		return
	}

	if r.Method == http.MethodDelete {
		s.snsDeleteSubscriptionHandler(w, r, topicArn)
		return
	}

	client, ctx, err := s.awsShared.GetSNSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	topics, err := client.ListSubscriptionsByTopic(*ctx, &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.respondAny(w, r, topics.Subscriptions)
}

func (s *server) snsSubscribeHandler(w http.ResponseWriter, r *http.Request, topicArn string) {
	var req topicSubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetSNSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parsedUrl, err := url.Parse(req.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protocol := parsedUrl.Scheme
	_, err = client.Subscribe(*ctx, &sns.SubscribeInput{
		Protocol: aws.String(protocol),
		TopicArn: aws.String(topicArn),
		Endpoint: aws.String(req.Url),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	topics, err := client.ListSubscriptionsByTopic(*ctx, &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.respondAny(w, r, topics.Subscriptions)
}

func (s *server) snsDeleteSubscriptionHandler(w http.ResponseWriter, r *http.Request, subArn string) {
	client, ctx, err := s.awsShared.GetSNSClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = client.Unsubscribe(*ctx, &sns.UnsubscribeInput{
		SubscriptionArn: aws.String(subArn),
	})

	w.Write([]byte("{}"))
}

func (s *server) getSNSAttribute(topicUrl string) (map[string]string, error) {
	client, ctx, err := s.awsShared.GetSNSClient()
	if err != nil {
		return nil, err
	}

	attrsOutput, err := client.GetTopicAttributes(*ctx, &sns.GetTopicAttributesInput{
		TopicArn: aws.String(topicUrl),
	})
	if err != nil {
		return nil, err
	}

	return attrsOutput.Attributes, nil
}
