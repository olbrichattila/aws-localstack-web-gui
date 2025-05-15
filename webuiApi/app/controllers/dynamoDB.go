package controller

import (
	"encoding/json"
	"webuiApi/app/repositories/awsshared"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/olbrichattila/gofra/pkg/app/request"
)

type tableField struct {
	AttributeName string `json:"attributeName"`
	AttributeType string `json:"attributeType"`
	KeyType       string `json:"keyType"`
}

type tableCreateRequest struct {
	Name   string       `json:"name"`
	Fields []tableField `json:"fields"`
}

func DynamoDBListTablesWithStartTable(limit int, exclusiveStartTable string, awsShared awsshared.AWSShared, r request.Requester) (string, error) {
	return dynamoDBListTables(limit, &exclusiveStartTable, awsShared, r)
}

func DynamoDBListTables(limit int, awsShared awsshared.AWSShared, r request.Requester) (string, error) {
	var exclusiveStartTable *string = nil
	return dynamoDBListTables(limit, exclusiveStartTable, awsShared, r)
}

func dynamoDBListTables(limit int, exclusiveStartTable *string, awsShared awsshared.AWSShared, r request.Requester) (string, error) {
	int32Limit := int32(limit)

	client, ctx, err := awsShared.GetDynamoDBClient()
	if err != nil {
		return "", err
	}

	tables, err := client.ListTables(*ctx, &dynamodb.ListTablesInput{
		Limit:                   &int32Limit,
		ExclusiveStartTableName: exclusiveStartTable,
	})

	if err != nil {
		return "", err
	}

	result, err := json.Marshal(tables.TableNames)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func DynamoDBNewTable(req tableCreateRequest, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetDynamoDBClient()
	if err != nil {
		return "", err
	}

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

	_, err = client.CreateTable(*ctx, &dynamodb.CreateTableInput{
		TableName:            &req.Name,
		AttributeDefinitions: attributes,
		KeySchema:            keySchemas,

		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		return "", err
	}

	return "{}", nil
}

func DynamoDBDeleteTable(tableName string, awsShared awsshared.AWSShared) (string, error) {
	client, ctx, err := awsShared.GetDynamoDBClient()
	if err != nil {
		return "", nil
	}

	client.DeleteTable(*ctx, &dynamodb.DeleteTableInput{
		TableName: &tableName,
	})

	return "{}", nil
}
