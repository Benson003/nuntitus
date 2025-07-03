package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Benson003/nuntius/middleware"
	"github.com/Benson003/nuntius/tools/files"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{Type: TypeFailure, Message: "invalid json body"})
		return
	}
	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{Type: TypeFailure, Message: "unauthorized"})
		return
	}
	uid, _ := uuid.Parse(userID)

	blog, err := h.DBObject.CreateBlog(uid, req.Title, req.Summary, req.PublishTime, req.Visibility)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{Type: TypeFailure, Message: "failed to create blog"})
		return
	}
	respondJSON(w, http.StatusCreated, MessageResponse{Type: TypeSuccess, Message: "created blog data", Data: map[string]interface{}{"blog_id": blog.BlogID}})
}

func (h *Handler) UpdateBlogHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{Type: TypeFailure, Message: "invalid json body"})
		return
	}
	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{Type: TypeFailure, Message: "unauthorized"})
		return
	}
	uid, _ := uuid.Parse(userID)

	blogIDStr := chi.URLParam(r, "blog_id")
	blogID, err := uuid.Parse(blogIDStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{Type: TypeFailure, Message: "invalid blog id"})
		return
	}

	// default to false if nil
	visibility := false
	if req.Visibility != nil {
		visibility = *req.Visibility
	}

	updatedBlog, err := h.DBObject.UpdateBlog(uid, blogID, req.Title, req.Summary, visibility, req.PublishTime)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{Type: TypeFailure, Message: "failed to update blog"})
		return
	}
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"blog_id": updatedBlog.BlogID,
	})
}

func (h *Handler) DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{Type: TypeFailure, Message: "unauthorized"})
		return
	}
	uid, _ := uuid.Parse(userID)

	blogIDStr := chi.URLParam(r, "blog_id")
	blogID, err := uuid.Parse(blogIDStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{Type: TypeFailure, Message: "invalid blog id"})
		return
	}

	deletedBlog, err := h.DBObject.DeleteBlog(uid, blogID)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{Type: TypeFailure, Message: "failed to delete blog"})
		return
	}
	respondJSON(w, http.StatusOK, deletedBlog)
}

func (h *Handler) FetchSingleBlogHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{Type: TypeFailure, Message: "unauthorized"})
		return
	}
	uid, _ := uuid.Parse(userID)

	blogIDStr := chi.URLParam(r, "blog_id")
	blogID, err := uuid.Parse(blogIDStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{Type: TypeFailure, Message: "invalid blog id"})
		return
	}

	blog, err := h.DBObject.GetBlogByID(uid, blogID)
	if err != nil {
		respondJSON(w, http.StatusNotFound, MessageResponse{Type: TypeFailure, Message: "blog not found"})
		return
	}
	respondJSON(w, http.StatusOK, blog)
}

func (h *Handler) FetchBlogsPaginatedHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{Type: TypeFailure, Message: "unauthorized"})
		return
	}
	uid, _ := uuid.Parse(userID)

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))

	if page <= 0 {
		page = 1
	}
	if count <= 0 {
		count = 10
	}

	blogs, err := h.DBObject.GetBlogsPaginated(uid, page, count)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{Type: TypeFailure, Message: "failed to fetch blogs"})
		return
	}
	respondJSON(w, http.StatusOK, blogs)
}

func (h *Handler) FetchBlogMarkdownHandler(w http.ResponseWriter, r *http.Request) {
	blogIDStr := chi.URLParam(r, "blog_id")
	blogID, err := uuid.Parse(blogIDStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{Type: TypeFailure, Message: "invalid blog id"})
		return
	}

	data, err := files.ReadFile(blogID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/markdown")
	w.Write(data)
}

func (h *Handler) UploadBlogMarkdownHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from JWT middleware
	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{
			Type:    TypeFailure,
			Message: "unauthorized",
		})
		return
	}

	// Get blog_id from URL
	blogIDStr := chi.URLParam(r, "blog_id")
	blogID, err := uuid.Parse(blogIDStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{
			Type:    TypeFailure,
			Message: "invalid blog id",
		})
		return
	}

	// Ownership verification
	_, err = h.DBObject.GetBlogByID(uuid.MustParse(userID), blogID)
	if err != nil {
		respondJSON(w, http.StatusNotFound, MessageResponse{
			Type:    TypeFailure,
			Message: "blog not found or unauthorized",
		})
		return
	}

	// Read uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{
			Type:    TypeFailure,
			Message: "failed to read uploaded file",
		})
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{
			Type:    TypeFailure,
			Message: "failed to read file content",
		})
		return
	}

	if len(content) == 0 {
		respondJSON(w, http.StatusBadRequest, MessageResponse{
			Type:    TypeFailure,
			Message: "file is empty",
		})
		return
	}

	// Use your custom UploadFile correctly
	if err := files.UploadFile(blogID, content); err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{
			Type:    TypeFailure,
			Message: "failed to save file",
		})
		return
	}

	respondJSON(w, http.StatusOK, MessageResponse{
		Type:    TypeSuccess,
		Message: "markdown file uploaded successfully",
	})
}
