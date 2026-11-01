package httpapi

import (
	"errors"
	"net/http"
	"strconv"

	"cakeshop/api/internal/cakes"
)

type cakeHandlers struct {
	svc *cakes.Service
}

// listParams are the only query parameters GET /v1/cakes accepts.
var listParams = map[string]bool{"occasion": true, "tag": true, "all": true}

// GET /v1/cakes?occasion=birthday&tag=eggless
func (h cakeHandlers) list(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	for name, values := range q {
		if !listParams[name] {
			writeError(w, http.StatusBadRequest, "bad_request", "unknown query parameter: "+name)
			return
		}
		if len(values) > 1 {
			writeError(w, http.StatusBadRequest, "bad_request", "parameter given more than once: "+name)
			return
		}
	}

	showAll := false
	if s := q.Get("all"); s != "" {
		b, err := strconv.ParseBool(s)
		if err != nil {
			writeError(w, http.StatusBadRequest, "bad_request", "all must be true or false")
			return
		}
		showAll = b // admins only, once auth exists (Sprint 4)
	}

	list, err := h.svc.List(cakes.Filter{
		Occasion:      q.Get("occasion"),
		Tag:           q.Get("tag"),
		OnlyAvailable: !showAll,
	})
	if errors.Is(err, cakes.ErrInvalid) {
		writeError(w, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal", "could not load cakes")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"cakes": list})
}

// GET /v1/cakes/{id}
func (h cakeHandlers) get(w http.ResponseWriter, r *http.Request) {
	c, err := h.svc.Get(r.PathValue("id"))
	if errors.Is(err, cakes.ErrNotFound) {
		writeError(w, http.StatusNotFound, "not_found", "cake not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal", "could not load cake")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// GET /v1/cakes/{id}/quote?size=8in&qty=1
func (h cakeHandlers) quote(w http.ResponseWriter, r *http.Request) {
	qty := 1
	if s := r.URL.Query().Get("qty"); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil {
			writeError(w, http.StatusBadRequest, "bad_request", "qty must be a number")
			return
		}
		qty = n
	}
	q, err := h.svc.Quote(r.PathValue("id"), r.URL.Query().Get("size"), qty)
	switch {
	case errors.Is(err, cakes.ErrNotFound):
		writeError(w, http.StatusNotFound, "not_found", "cake or size not found")
	case errors.Is(err, cakes.ErrUnavailable):
		writeError(w, http.StatusConflict, "unavailable", "this cake is not available right now")
	case err != nil:
		writeError(w, http.StatusBadRequest, "bad_request", err.Error())
	default:
		writeJSON(w, http.StatusOK, q)
	}
}