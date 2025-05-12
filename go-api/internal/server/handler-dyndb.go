package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func (s *server) getDynamoDBTables(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 4 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(urlParts[3])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var exclusiveStartTable *string = nil
	int32Limit := int32(limit)

	if len(urlParts) >= 5 {
		exclusiveStartTable = &urlParts[4]
	}

	client, ctx, err := s.awsShared.GetDynamoDBClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tables, err := client.ListTables(*ctx, &dynamodb.ListTablesInput{
		Limit:                   &int32Limit,
		ExclusiveStartTableName: exclusiveStartTable,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondAny(w, r, tables.TableNames)
}

func (s *server) handleDynamoDBTable(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		s.dynamoDBTableDelete(w, r)
		return
	}

	var req tableCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, ctx, err := s.awsShared.GetDynamoDBClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{}"))
}

func (s *server) dynamoDBTableDelete(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 4 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	tableName := urlParts[3]

	client, ctx, err := s.awsShared.GetDynamoDBClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client.DeleteTable(*ctx, &dynamodb.DeleteTableInput{
		TableName: &tableName,
	})

	w.Write([]byte("{}"))
}
