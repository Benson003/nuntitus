package router

import (
	"net/http"
	"time"

	embeding "github.com/Benson003/nuntius"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func New() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(cors.AllowAll().Handler)
	router.Route("/api/v1",func(r chi.Router) {

	})
	fs := http.FileServer(embeding.FS())

	// Serve index.html on unknown paths for SPA routing
	router.NotFound(func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = "/index.html"
		fs.ServeHTTP(w, req)
	})
	router.Handle("/*", fs)


	return router
}
