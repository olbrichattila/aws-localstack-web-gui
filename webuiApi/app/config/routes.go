package appconfig

import (
	"net/http"
	controller "webuiApi/app/controllers"

	"github.com/olbrichattila/gofra/pkg/app/router"
)

var Routes = []router.ControllerAction{
	// {
	// 	Path:        "/",
	// 	RequestType: []string{http.MethodGet},
	// 	Fn: func() map[string]string {
	// 		return map[string]string{"result": "It works"}
	// 	},
	// },
	{
		Path:        "/api/s3/buckets",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: []string{http.MethodGet},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "ListBucketsAction",
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "CreateBucketAction",
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: []string{http.MethodDelete},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "DeleteBucketAction",
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: []string{http.MethodDelete},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "DeleteBucketAction",
	},
	{
		Path:        "/api/s3/list/:bucketName",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/list/:bucketName",
		RequestType: []string{http.MethodGet},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "GetBucketContent",
	},
	{
		Path:        "/api/s3/file_upload",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "FileUpload",
	},
	{
		Path:        "/api/s3/buckets/upload",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/buckets/upload",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "FileUploadToS3",
	},
	{
		Path:        "/api/s3/load",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "ViewFile",
	},
	{
		Path:        "/api/s3/buckets/delete/object",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/buckets/delete/object",
		RequestType: []string{http.MethodDelete},
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "DeleteFile",
	},
	{
		Path:        "/api/settings",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/settings",
		RequestType: []string{http.MethodGet},
		Fn:          controller.GetSettingsAction,
	},
	{
		Path:        "/api/settings",
		RequestType: []string{http.MethodPost},
		Fn:          controller.SaveSettingsAction,
		// Middlewares: []any{middleware.CorsMiddleware},
	},

	// SQS
	{
		Path:        "/api/sqs/attributes",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/attributes",
		RequestType: []string{http.MethodGet},
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSGetAttributesAction",
	},
	{
		Path:        "/api/sqs/attributes",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSGetAttributeAction",
	},
	{
		Path:        "/api/sqs",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSCreateQueueAction",
	},
	{
		Path:        "/api/sqs",
		RequestType: []string{http.MethodDelete},
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSDeleteQueueAction",
	},
	{
		Path:        "/api/sqs/purge",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/purge",
		RequestType: []string{http.MethodDelete},
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSPurgeQueueAction",
	},
	{
		Path:        "/api/sqs/message/send",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/message/send",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSendMessageAction",
	},
	{
		Path:        "/api/sqs/message/receive",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/message/receive",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSReceiveMessages",
	},

	// SNS
	{
		Path:        "/api/sns/attributes",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sns/attributes",
		RequestType: []string{http.MethodGet},
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSGetAttributes",
	},
	{
		Path:        "/api/sns",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sns",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSCreateTopic",
	},
	{
		Path:        "/api/sns",
		RequestType: []string{http.MethodDelete},
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSDeleteTopic",
	},
	{
		Path:        "/api/sns/sub/:arn/publish",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sns/sub/:arn/publish",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSPublishToTopicARN",
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: []string{http.MethodGet},
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSGetSubscriptionsByARN",
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSCreateSubscriptionForARN",
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: []string{http.MethodDelete},
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSDeleteSubscriptionByARN",
	},

	// DynamoDB
	{
		Path:        "/api/dynamodb-list/:itemCount",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb-list/:itemCount",
		RequestType: []string{http.MethodGet},
		Controller:  func() any { return &controller.DynamoDBController{} },
		ActionName:  "DynamoDBListTables",
	},
	{
		Path:        "/api/dynamodb-list/:itemCount/:exclusiveStartTable",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb-list/:itemCount/:exclusiveStartTable",
		RequestType: []string{http.MethodGet},
		Controller:  func() any { return &controller.DynamoDBController{} },
		ActionName:  "DynamoDBListTablesWithStartTable",
	},
	{
		Path:        "/api/dynamodb",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb",
		RequestType: []string{http.MethodPost},
		Controller:  func() any { return &controller.DynamoDBController{} },
		ActionName:  "DynamoDBNewTable",
	},
	{
		Path:        "/api/dynamodb/:tableName",
		RequestType: []string{http.MethodOptions},
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb/:tableName",
		RequestType: []string{http.MethodDelete},
		Controller:  func() any { return &controller.DynamoDBController{} },
		ActionName:  "DynamoDBDeleteTable",
	},
}
