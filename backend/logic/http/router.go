package http

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"veric-backend-mvp/logic/http/open_api"
)

func registerRouter() http.Handler {
	r := mux.NewRouter()

	open_api.RegisterRouter(r.PathPrefix("/api").Subrouter())

	return handlers.LoggingHandler(os.Stdout, r)
}
