package main

import (
	bootstrap "webuiApi/app"
	appconfig "webuiApi/app/config"

	"github.com/olbrichattila/gofra"
)

func main() {
	gofra.Run(
		bootstrap.Bootstrap,
		appconfig.Routes,
		appconfig.Jobs,
		appconfig.Middlewares,
		appconfig.DiBindings,
		appconfig.ConsoleCommands,
		appconfig.ViewFuncConfig,
		appconfig.TemplateAutoLoad,
		appconfig.RouteValidationRules,
		appconfig.ValidatorRules,
	)
}
