package appconfig

import (
	"net/http"
	controller "webuiApi/app/controllers"

	"github.com/olbrichattila/gofra/pkg/app/router"
)

var Routes = []router.ControllerAction{
	// {
	// 	Path:        "/",
	// 	RequestType: http.MethodGet,
	// 	Fn: func() map[string]string {
	// 		return map[string]string{"result": "It works"}
	// 	},
	// },
	{
		Path:        "/api/s3/buckets",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: http.MethodGet,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "ListBucketsAction",
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "CreateBucketAction",
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: http.MethodDelete,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "DeleteBucketAction",
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: http.MethodDelete,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "DeleteBucketAction",
	},
	{
		Path:        "/api/s3/list/:bucketName",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/list/:bucketName",
		RequestType: http.MethodGet,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "GetBucketContent",
	},
	{
		Path:        "/api/s3/file_upload",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "FileUpload",
	},
	{
		Path:        "/api/s3/buckets/upload",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/buckets/upload",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "FileUploadToS3",
	},
	{
		Path:        "/api/s3/load",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "ViewFile",
	},
	{
		Path:        "/api/s3/buckets/delete/object",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/buckets/delete/object",
		RequestType: http.MethodDelete,
		Controller:  func() any { return &controller.S3Controller{} },
		ActionName:  "DeleteFile",
	},
	{
		Path:        "/api/settings",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/settings",
		RequestType: http.MethodGet,
		Fn:          controller.GetSettingsAction,
	},
	{
		Path:        "/api/settings",
		RequestType: http.MethodPost,
		Fn:          controller.SaveSettingsAction,
		// Middlewares: []any{middleware.CorsMiddleware},
	},

	// SQS
	{
		Path:        "/api/sqs/attributes",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/attributes",
		RequestType: http.MethodGet,
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSGetAttributesAction",
	},
	{
		Path:        "/api/sqs/attributes",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSGetAttributeAction",
	},
	{
		Path:        "/api/sqs",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSCreateQueueAction",
	},
	{
		Path:        "/api/sqs",
		RequestType: http.MethodDelete,
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSDeleteQueueAction",
	},
	{
		Path:        "/api/sqs/purge",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/purge",
		RequestType: http.MethodDelete,
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSPurgeQueueAction",
	},
	{
		Path:        "/api/sqs/message/send",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/message/send",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSendMessageAction",
	},
	{
		Path:        "/api/sqs/message/receive",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/message/receive",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.SQSController{} },
		ActionName:  "SQSReceiveMessages",
	},

	// SNS
	{
		Path:        "/api/sns/attributes",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sns/attributes",
		RequestType: http.MethodGet,
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSGetAttributes",
	},
	{
		Path:        "/api/sns",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sns",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSCreateTopic",
	},
	{
		Path:        "/api/sns",
		RequestType: http.MethodDelete,
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSDeleteTopic",
	},
	{
		Path:        "/api/sns/sub/:arn/publish",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sns/sub/:arn/publish",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSPublishToTopicARN",
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: http.MethodGet,
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSGetSubscriptionsByARN",
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSCreateSubscriptionForARN",
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: http.MethodDelete,
		Controller:  func() any { return &controller.SNSController{} },
		ActionName:  "SNSDeleteSubscriptionByARN",
	},

	// DynamoDB
	{
		Path:        "/api/dynamodb-list/:itemCount",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb-list/:itemCount",
		RequestType: http.MethodGet,
		Controller:  func() any { return &controller.DynamoDBController{} },
		ActionName:  "DynamoDBListTables",
	},
	{
		Path:        "/api/dynamodb-list/:itemCount/:exclusiveStartTable",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb-list/:itemCount/:exclusiveStartTable",
		RequestType: http.MethodGet,
		Controller:  func() any { return &controller.DynamoDBController{} },
		ActionName:  "DynamoDBListTablesWithStartTable",
	},
	{
		Path:        "/api/dynamodb",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb",
		RequestType: http.MethodPost,
		Controller:  func() any { return &controller.DynamoDBController{} },
		ActionName:  "DynamoDBNewTable",
	},
	{
		Path:        "/api/dynamodb/:tableName",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb/:tableName",
		RequestType: http.MethodDelete,
		Controller:  func() any { return &controller.DynamoDBController{} },
		ActionName:  "DynamoDBDeleteTable",
	},
}
