package appconfig

import middleware "webuiApi/app/middlewares"

// Add middlewares here to execute at every load
var Middlewares = []interface{}{
	middleware.CorsMiddleware,
	middleware.OptionsMiddleware,
	middleware.SessionMiddleware,
}
