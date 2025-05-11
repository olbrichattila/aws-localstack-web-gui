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

}
