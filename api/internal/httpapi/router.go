// Package httpapi exposes the services over HTTP. It only translates between
// HTTP and the domain packages; no business rules live here.
package httpapi

import (
	"net/http"

	"cakeshop/api/internal/cakes"
)

// Deps are the services the router needs. Orders, chat and payments are added
// in later sprints (see docs/process/sprint-plan.md).
type Deps struct {
	Cakes          *cakes.Service
	AllowedOrigins []string
}

// NewRouter wires all routes. Uses Go's standard library router (method + path patterns).
func NewRouter(d Deps) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	ch := cakeHandlers{svc: d.Cakes}
	mux.HandleFunc("GET /v1/cakes", ch.list)
	mux.HandleFunc("GET /v1/cakes/{id}", ch.get)
	mux.HandleFunc("GET /v1/cakes/{id}/quote", ch.quote)

	// Coming next (see backlog):
	//   GET  /v1/slots?date=          delivery slots          (Sprint 3)
	//   POST /v1/orders               create order            (Sprint 3)
	//   /v1/admin/...                 owner endpoints         (Sprint 4)
	//   POST /v1/webhooks/payments    payment gateway         (Sprint 5)
	//   POST /v1/chat                 AI assistant            (Sprint 6, separate Lambda)

	var h http.Handler = mux
	h = cors(d.AllowedOrigins)(h)
	h = logRequests(h)
	h = recoverPanics(h)
	return h
}
