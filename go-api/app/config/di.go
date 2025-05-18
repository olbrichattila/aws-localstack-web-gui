package appconfig

import (
	"webuiApi/app/repositories/awsshared"
	"webuiApi/app/repositories/database"

	"github.com/olbrichattila/godi"
	"github.com/olbrichattila/gofra/pkg/app/config"
)

var DiBindings = []config.DiCallback{
	func(di godi.Container) (string, interface{}, error) {
		return "app.repositories.database.Database", database.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		awsshared := awsshared.New()
		return "app.repositories.awsshared.AWSShared", awsshared, nil
	},
}
