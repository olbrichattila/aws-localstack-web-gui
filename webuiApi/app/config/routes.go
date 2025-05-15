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
		Fn:          controller.ListBucketsAction,
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: http.MethodPost,
		Fn:          controller.CreateBucketAction,
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: http.MethodDelete,
		Fn:          controller.DeleteBucketAction,
	},
	{
		Path:        "/api/s3/buckets",
		RequestType: http.MethodDelete,
		Fn:          controller.DeleteBucketAction,
	},
	{
		Path:        "/api/s3/list/:bucketName",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/list/:bucketName",
		RequestType: http.MethodGet,
		Fn:          controller.GetBucketContent,
	},
	{
		Path:        "/api/s3/file_upload",
		RequestType: http.MethodPost,
		Fn:          controller.FileUpload,
	},
	{
		Path:        "/api/s3/buckets/upload",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/buckets/upload",
		RequestType: http.MethodPost,
		Fn:          controller.FileUploadToS3,
	},
	{
		Path:        "/api/s3/load",
		RequestType: http.MethodPost,
		Fn:          controller.ViewFile,
	},
	{
		Path:        "/api/s3/buckets/delete/object",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/s3/buckets/delete/object",
		RequestType: http.MethodDelete,
		Fn:          controller.DeleteFile,
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

	// SQL
	{
		Path:        "/api/sqs/attributes",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/attributes",
		RequestType: http.MethodGet,
		Fn:          controller.SQSGetAttributesAction,
	},
	{
		Path:        "/api/sqs/attributes",
		RequestType: http.MethodPost,
		Fn:          controller.SQSGetAttributeAction,
	},
	{
		Path:        "/api/sqs",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs",
		RequestType: http.MethodPost,
		Fn:          controller.SQSCreateQueueAction,
	},
	{
		Path:        "/api/sqs",
		RequestType: http.MethodDelete,
		Fn:          controller.SQSDeleteQueueAction,
	},
	{
		Path:        "/api/sqs/purge",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/purge",
		RequestType: http.MethodDelete,
		Fn:          controller.SQSPurgeQueueAction,
	},
	{
		Path:        "/api/sqs/message/send",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/message/send",
		RequestType: http.MethodPost,
		Fn:          controller.SQSendMessageAction,
	},
	{
		Path:        "/api/sqs/message/receive",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sqs/message/receive",
		RequestType: http.MethodPost,
		Fn:          controller.SQSReceiveMessages,
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
		Fn:          controller.SNSGetAttributes,
	},
	{
		Path:        "/api/sns",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sns",
		RequestType: http.MethodPost,
		Fn:          controller.SNSCreateTopic,
	},
	{
		Path:        "/api/sns",
		RequestType: http.MethodDelete,
		Fn:          controller.SNSDeleteTopic,
	},
	{
		Path:        "/api/sns/sub/:arn/publish",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sns/sub/:arn/publish",
		RequestType: http.MethodPost,
		Fn:          controller.SNSPublishToTopicARN,
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: http.MethodGet,
		Fn:          controller.SNSGetSubscriptionsByARN,
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: http.MethodPost,
		Fn:          controller.SNSCreateSubscriptionForARN,
	},
	{
		Path:        "/api/sns/sub/:arn",
		RequestType: http.MethodDelete,
		Fn:          controller.SNSDeleteSubscriptionByARN,
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
		Fn:          controller.DynamoDBListTables,
	},
	{
		Path:        "/api/dynamodb-list/:itemCount/:exclusiveStartTable",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb-list/:itemCount/:exclusiveStartTable",
		RequestType: http.MethodGet,
		Fn:          controller.DynamoDBListTablesWithStartTable,
	},
	{
		Path:        "/api/dynamodb",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb",
		RequestType: http.MethodPost,
		Fn:          controller.DynamoDBNewTable,
	},

	{
		Path:        "/api/dynamodb/:tableName",
		RequestType: http.MethodOptions,
		Fn:          func() {},
	},
	{
		Path:        "/api/dynamodb/:tableName",
		RequestType: http.MethodDelete,
		Fn:          controller.DynamoDBDeleteTable,
	},
}
