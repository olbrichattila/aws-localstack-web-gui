package middleware

import (
	"github.com/olbrichattila/gofra/pkg/app/request"
)

// OptionsMiddleware function can take any parameters defined in the Di config
func OptionsMiddleware(r request.Requester) bool {

	if r.GetRequest().Method == "OPTIONS" {
		return false
	}

	return true
}
