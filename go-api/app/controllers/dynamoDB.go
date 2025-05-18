package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"webuiApi/app/repositories/awsshared"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/olbrichattila/gofra/pkg/app/gofraerror"
	"github.com/olbrichattila/gofra/pkg/app/request"
)

type DynamoDBController struct {
	client *dynamodb.Client
	ctx    *context.Context
}

type tableField struct {
	AttributeName string `json:"attributeName"`
	AttributeType string `json:"attributeType"`
	KeyType       string `json:"keyType"`
}

type tableCreateRequest struct {
	Name   string       `json:"name"`
	Fields []tableField `json:"fields"`
}

func (c *DynamoDBController) Before(awsShared awsshared.AWSShared) {
	// TODO add error handling
	c.client, c.ctx, _ = awsShared.GetDynamoDBClient()
}

func (c *DynamoDBController) DynamoDBListTablesWithStartTable(limit int, exclusiveStartTable string, r request.Requester) (string, error) {
	return c.dynamoDBListTables(limit, &exclusiveStartTable, r)
}

func (c *DynamoDBController) DynamoDBListTables(limit int, r request.Requester) (string, error) {
	var exclusiveStartTable *string = nil
	return c.dynamoDBListTables(limit, exclusiveStartTable, r)
}

func (c *DynamoDBController) dynamoDBListTables(limit int, exclusiveStartTable *string, r request.Requester) (string, error) {
	int32Limit := int32(limit)

	tables, err := c.client.ListTables(*c.ctx, &dynamodb.ListTablesInput{
		Limit:                   &int32Limit,
		ExclusiveStartTableName: exclusiveStartTable,
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	result, err := json.Marshal(tables.TableNames)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(result), nil
}

func (c *DynamoDBController) DynamoDBNewTable(req tableCreateRequest) (string, error) {
	attributes := make([]types.AttributeDefinition, len(req.Fields))
	keySchemas := make([]types.KeySchemaElement, len(req.Fields))

	for i, field := range req.Fields {
		attributes[i] = types.AttributeDefinition{
			AttributeName: &field.AttributeName,
			AttributeType: types.ScalarAttributeType(field.AttributeType),
		}
		keySchemas[i] = types.KeySchemaElement{
			AttributeName: &field.AttributeName,
			KeyType:       types.KeyType(field.KeyType),
		}
	}

	_, err := c.client.CreateTable(*c.ctx, &dynamodb.CreateTableInput{
		TableName:            &req.Name,
		AttributeDefinitions: attributes,
		KeySchema:            keySchemas,

		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return "{}", nil
}

func (c *DynamoDBController) DynamoDBDeleteTable(tableName string) (string, error) {
	_, err := c.client.DeleteTable(*c.ctx, &dynamodb.DeleteTableInput{
		TableName: &tableName,
	})

	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return "{}", nil
}
