package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"webuiApi/app/repositories/awsshared"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/olbrichattila/gofra/pkg/app/gofraerror"
)

type SNSController struct {
	client *sns.Client
	ctx    *context.Context
}

type topicRequest struct {
	Name string `json:"name"`
}

type topicMessageRequest struct {
	Message string `json:"message"`
}

type topicSubscribeRequest struct {
	Url string `json:"url"`
}

func (c *SNSController) Before(awsShared awsshared.AWSShared) error {
	var err error
	c.client, c.ctx, err = awsShared.GetSNSClient()
	if err != nil {
		return gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return nil
}

// SnsAction function can take any parameters defined in the Di config
func (c *SNSController) SNSGetAttributes() (string, error) {
	topics, err := c.client.ListTopics(*c.ctx, &sns.ListTopicsInput{})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	response := make([]map[string]string, len(topics.Topics))
	for i, topic := range topics.Topics {
		attrs, err := c.getSNSAttribute(c.client, c.ctx, *topic.TopicArn)
		if err != nil {
			return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
		}

		response[i] = attrs
	}

	res, err := json.Marshal(response)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(res), nil
}

func (c *SNSController) getSNSAttribute(client *sns.Client, ctx *context.Context, topicUrl string) (map[string]string, error) {
	attrsOutput, err := client.GetTopicAttributes(*ctx, &sns.GetTopicAttributesInput{
		TopicArn: aws.String(topicUrl),
	})
	if err != nil {
		return nil, gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return attrsOutput.Attributes, nil
}

func (c *SNSController) SNSCreateTopic(req topicRequest) (string, error) {
	_, err := c.client.CreateTopic(*c.ctx, &sns.CreateTopicInput{
		Name: &req.Name,
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return c.SNSGetAttributes()
}

func (c *SNSController) SNSDeleteTopic(req topicRequest) (string, error) {
	_, err := c.client.DeleteTopic(*c.ctx, &sns.DeleteTopicInput{
		TopicArn: aws.String(req.Name),
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return c.SNSGetAttributes()
}

func (c *SNSController) SNSPublishToTopicARN(topicArn string, req topicMessageRequest) (string, error) {
	_, err := c.client.Publish(*c.ctx, &sns.PublishInput{
		Message:  &req.Message,
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return "{}", nil
}

func (c *SNSController) SNSGetSubscriptionsByARN(topicArn string) (string, error) {
	topics, err := c.client.ListSubscriptionsByTopic(*c.ctx, &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	res, err := json.Marshal(topics.Subscriptions)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(res), nil
}

func (c *SNSController) SNSCreateSubscriptionForARN(topicArn string, req topicSubscribeRequest) (string, error) {
	parsedUrl, err := url.Parse(req.Url)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	protocol := parsedUrl.Scheme
	_, err = c.client.Subscribe(*c.ctx, &sns.SubscribeInput{
		Protocol: aws.String(protocol),
		TopicArn: aws.String(topicArn),
		Endpoint: aws.String(req.Url),
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	topics, err := c.client.ListSubscriptionsByTopic(*c.ctx, &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	result, err := json.Marshal(topics.Subscriptions)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(result), nil
}

func (c *SNSController) SNSDeleteSubscriptionByARN(topicArn string) (string, error) {
	_, err := c.client.Unsubscribe(*c.ctx, &sns.UnsubscribeInput{
		SubscriptionArn: aws.String(topicArn),
	})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return "{}", nil
}
