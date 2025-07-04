package router

import (
	"github.com/Benson003/nuntius/api"
	"github.com/Benson003/nuntius/middleware"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes mounts all API routes under /api/v1.
func RegisterRoutes(r chi.Router, h *api.Handler) {
	// AUTH routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", h.SignupHandler)
		r.Post("/login", h.LoginHandler)

		// Protected user routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuthMiddleware)
			r.Put("/user", h.UpdateUserHandler)
			r.Delete("/user", h.DeleteUserHandler)
			r.Get("/user", h.FetchUserHandler)
		})
	})

	// BLOG routes
	r.Route("/blogs", func(r chi.Router) {
		r.Use(middleware.JWTAuthMiddleware)

		r.Post("/", h.CreateBlogHandler)
		r.Put("/{blogID}", h.UpdateBlogHandler)
		r.Delete("/{blogID}", h.DeleteBlogHandler)
		r.Get("/{blogID}", h.FetchSingleBlogHandler)
		r.Get("/", h.FetchBlogsPaginatedHandler)

		// Markdown upload endpoint
		r.Post("/{blogID}/upload", h.UploadBlogMarkdownHandler)
		r.Post("/{blogID}/download", h.FetchBlogMarkdownHandler)
	})
}
