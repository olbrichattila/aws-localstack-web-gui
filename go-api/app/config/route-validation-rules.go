package appconfig

import (
	"github.com/olbrichattila/gofra/pkg/app/validator"
)

var RouteValidationRules = map[string]validator.ValidationRule{
	"register": {
		Redirect: "/register",
		Rules: map[string]string{
			"password": "minSize:6|maxSize:255",
			"name":     "minSize:6|maxSize:255",
			"email":    "required|email",
		},
		// CustomRule: func(fields map[string]string) (validator.ValidationErrors, bool) {
		// 	return validator.ValidationErrors{"name": []string{"error1", "error2"}}, false
		// },
	},
}
