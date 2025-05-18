package middleware

import (
	"net/http"

	"github.com/olbrichattila/gofra/pkg/app/request"
	"github.com/olbrichattila/gofra/pkg/app/session"
)

func SessionMiddleware(w http.ResponseWriter, r request.Requester, s session.Sessioner) {
	s.Init(w, r.GetRequest())
}
