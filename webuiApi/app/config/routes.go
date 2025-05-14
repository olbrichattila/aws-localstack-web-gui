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
}
