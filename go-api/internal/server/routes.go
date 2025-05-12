package server

import "net/http"

func (s *server) initRoutes() {
	jsonGetMiddlewareGroup := []handleFuncWithReturn{
		s.corsMiddleware,
		s.jSONMiddleware,
		s.getMiddleware,
	}

	jsonPostMiddlewareGroup := []handleFuncWithReturn{
		s.corsMiddleware,
		s.jSONMiddleware,
		s.postMIddleware,
	}

	jsonGetPostMiddlewareGroup := []handleFuncWithReturn{
		s.corsMiddleware,
		s.jSONMiddleware,
		s.getPostMiddleware,
	}

	jsonGetPostDeleteMiddlewareGroup := []handleFuncWithReturn{
		s.corsMiddleware,
		s.jSONMiddleware,
		s.getPostDeleteMiddleware,
	}

	// Serve files in the current directory at root
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.Handle("/s3/", http.StripPrefix("/s3/", fs))
	http.Handle("/sqs/", http.StripPrefix("/sqs/", fs))
	http.Handle("/sqdynamodb/", http.StripPrefix("/sqdynamodb/", fs))
	http.Handle("/settings/", http.StripPrefix("/settings/", fs))

	// Serve files in ./static under /static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// S3 endpoints
	http.HandleFunc(
		"/api/health",
		s.initMiddlewares(
			s.healthHandler,
			jsonGetMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/settings",
		s.initMiddlewares(
			s.getSettingsHandler,
			jsonGetPostMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/s3/buckets",
		s.initMiddlewares(
			s.getS3BucketListHandler,
			jsonGetPostDeleteMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/s3/list/",
		s.initMiddlewares(
			s.getS3BucketContentListHandler,
			jsonGetMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/s3/file_upload",
		s.initMiddlewares(
			s.uploadToServerHandler,
			jsonPostMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/s3/buckets/upload",
		s.initMiddlewares(
			s.uploadToS3Handler,
			jsonPostMiddlewareGroup...,
		),
	)
	http.HandleFunc(
		"/api/s3/buckets/delete/object",
		s.initMiddlewares(
			s.deleteS3ObjectHandler,
			s.corsMiddleware,
			s.jSONMiddleware,
			s.getPostDeleteMiddleware,
		),
	)

	http.HandleFunc(
		"/api/s3/load",
		s.initMiddlewares(
			s.loads3ObjectHandler,
			s.getPostMiddleware,
		),
	)

	// SQS endpoints
	http.HandleFunc(
		"/api/sqs/attributes",
		s.initMiddlewares(
			s.getSqsListHandler,
			jsonGetPostDeleteMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/sqs",
		s.initMiddlewares(
			s.addSqsQueueHandler,
			jsonGetPostDeleteMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/sqs/purge",
		s.initMiddlewares(
			s.purgeSqsQueueHandler,
			jsonGetPostDeleteMiddlewareGroup..., // Should be delete only, not yet exists
		),
	)

	http.HandleFunc(
		"/api/sqs/message/send",
		s.initMiddlewares(
			s.sendMessageHandler,
			jsonPostMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/sqs/message/receive",
		s.initMiddlewares(
			s.getMessagesHandler,
			jsonPostMiddlewareGroup...,
		),
	)

	// SNS routes
	http.HandleFunc(
		"/api/sns/attributes",
		s.initMiddlewares(
			s.getSNSAttributes,
			jsonGetMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/sns",
		s.initMiddlewares(
			s.newTopicHandler,
			jsonGetPostDeleteMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/sns/",
		s.initMiddlewares(
			s.topicApiHandler,
			jsonGetPostDeleteMiddlewareGroup...,
		),
	)

	// DynamoDB routes
	http.HandleFunc(
		"/api/dynamodb-list/",
		s.initMiddlewares(
			s.getDynamoDBTables,
			jsonGetMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/dynamodb",
		s.initMiddlewares(
			s.handleDynamoDBTable,
			jsonGetPostDeleteMiddlewareGroup...,
		),
	)

	http.HandleFunc(
		"/api/dynamodb/",
		s.initMiddlewares(
			s.handleDynamoDBTable,
			jsonGetPostDeleteMiddlewareGroup...,
		),
	)

}
