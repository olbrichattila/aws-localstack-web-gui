package appconfig

import (
	"webuiApi/app/repositories/awsshared"
	"webuiApi/app/repositories/database"
	"webuiApi/app/repositories/snslistener"

	"github.com/olbrichattila/godi"
	"github.com/olbrichattila/gofra/pkg/app/config"
)

var DiBindings = []config.DiCallback{
	func(di godi.Container) (string, interface{}, error) {
		return "app.repositories.database.Database", database.New(), nil
		// return "app.repositories.database.Database",
		// 	func() any {
		// 		return database.New()
		// 	},
		// 	nil
	},
	func(di godi.Container) (string, interface{}, error) {
		awsshared := awsshared.New()
		return "app.repositories.awsshared.AWSShared", awsshared, nil
	},

	func(di godi.Container) (string, interface{}, error) {
		snslistener := snslistener.New()
		return "app.repositories.snslistener.SNSListener", snslistener, nil
	},
}
