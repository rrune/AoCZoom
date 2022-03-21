package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rrune/AoCZoom/api"
	"github.com/rrune/AoCZoom/models"
)

var config models.Config

func NewRouter(api api.Api, cfg models.Config) http.Handler {
	config = cfg
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("../data/public"))
	r.Handle("/*", http.StripPrefix("/", fs))

	r.Get("/getImage", api.GetImage)

	r.Group(func(r chi.Router) {
		r.Use(auth)
		r.Post("/updateImage", api.UpdateImage)
	})
	return r
}
