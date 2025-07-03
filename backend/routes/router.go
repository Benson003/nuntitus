package router

import (
	"net/http"
	"strings"
	"time"

	embeding "github.com/Benson003/nuntius"
	"github.com/Benson003/nuntius/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(h *api.Handler) *chi.Mux {
	router := chi.NewRouter()

	// Global middleware
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.Compress(5))
	router.Use(cors.AllowAll().Handler)

	// Mount API routes
	router.Route("/api/v1", func(r chi.Router) {
		RegisterRoutes(r, h) // ✅ correct
	})

	// Static file server for embedded frontend
	fs := http.FileServer(embeding.FS())

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		r.URL.Path = "/index.html"
		fs.ServeHTTP(w, r)
	})

	router.Handle("/*", fs)

	return router
}
