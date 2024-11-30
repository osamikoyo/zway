package handler

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type handlerfunc func(w http.ResponseWriter, r *http.Request) error

func RegisterRoute(r *chi.Mux) {
	r.Get("/ping", getErr(pingHandler))
	r.Get("/text/get/:hash", getErr(getTextHandler))
	r.Get("/text/get", getErr(AddText))
}

func getErr(h handlerfunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := h(writer, request); err != nil {
			slog.Error(err.Error())
		}
	}
}
