package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Benson003/nuntius/database"
	"github.com/Benson003/nuntius/middleware"
	"github.com/Benson003/nuntius/tools/hashing"
	"github.com/Benson003/nuntius/tools/utils"
	"github.com/google/uuid"
)

func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB max

	var req SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)
	email := strings.TrimSpace(req.Email)

	if username == "" || email == "" || password == "" {
		http.Error(w, "username, email, and password are required", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(email) { // optional helper
		http.Error(w, "invalid email format", http.StatusBadRequest)
		return
	}

	hash, err := hashing.HashPassword(password)
	if err != nil {
		http.Error(w, "couldn't hash password", http.StatusInternalServerError)
		return
	}

	user, err := h.DBObject.CreateUser(username, hash, email)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") { // SQLite detection
			http.Error(w, "email or username already exists", http.StatusConflict)
			return
		}
		http.Error(w, fmt.Sprintf("couldn't create user: %v", err), http.StatusInternalServerError)
		return
	}

	tokenString, err := utils.GenerateJWT(user.UserID.String())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{
			Type:    TypeFailure,
			Message: "failed to generate token",
		})
		return
	}
	respondJSON(w, http.StatusCreated, MessageResponse{
		Type:    TypeSuccess,
		Message: "user created successfully",
		Data: map[string]interface{}{
			"token": tokenString,
		},
	})
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body into LoginRequest

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{
			Type:    TypeFailure,
			Message: "invalid json body",
		})
		return
	}

	// Trim spaces for safety (optional but recommended)
	req.UsernameOrEmail = strings.TrimSpace(req.UsernameOrEmail)
	req.Password = strings.TrimSpace(req.Password)

	// Basic input validation
	if req.UsernameOrEmail == "" || req.Password == "" {
		respondJSON(w, http.StatusBadRequest, MessageResponse{
			Type:    TypeFailure,
			Message: "username/email and password required",
		})
		return
	}

	var user *database.UserTable

	// Try fetching by email
	user, err := h.DBObject.GetUserByUsernameOrEmail(req.UsernameOrEmail, req.UsernameOrEmail)

	if err != nil {
		// If both fail, the user was not found
		respondJSON(w, http.StatusNotFound, MessageResponse{
			Type:    TypeFailure,
			Message: "failed to find user",
		})
		return
	}

	// Check the password against the stored hash
	if err := hashing.VerifyPassword(user.PasswordHash, req.Password); err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{
			Type:    TypeFailure,
			Message: "invalid credentials",
		})
		return
	}

	// Generate JWT token for the user
	tokenString, err := utils.GenerateJWT(user.UserID.String())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{
			Type:    TypeFailure,
			Message: "failed to generate token",
		})
		return
	}

	// Respond with the JWT token on successful login
	respondJSON(w, http.StatusOK, MessageResponse{
		Type:    TypeSuccess,
		Message: "login successful",
		Data: map[string]interface{}{
			"token": tokenString,
		},
	})
}

func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user_id from JWT middleware
	userIDStr, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{
			Type:    TypeFailure,
			Message: "unauthorized",
		})
		return
	}
	userID, _ := uuid.Parse(userIDStr)

	// Parse incoming JSON request
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, MessageResponse{
			Type:    TypeFailure,
			Message: "invalid json body",
		})
		return
	}

	// Optional: validate email if provided
	if req.Email != "" && !utils.IsValidEmail(req.Email) {
		respondJSON(w, http.StatusBadRequest, MessageResponse{
			Type:    TypeFailure,
			Message: "invalid email format",
		})
		return
	}

	// Hash password if provided
	hashedPassword := ""
	if req.Password != "" {
		hash, err := hashing.HashPassword(req.Password)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, MessageResponse{
				Type:    TypeFailure,
				Message: "failed to hash password",
			})
			return
		}
		hashedPassword = hash
	}

	// Update user in DB
	updatedUser, err := h.DBObject.UpdateUser(userID, req.Username, hashedPassword, req.Email, req.FirstName, req.LastName)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{
			Type:    TypeFailure,
			Message: "failed to update user",
		})
		return
	}

	respondJSON(w, http.StatusOK, MessageResponse{
		Type:    TypeSuccess,
		Message: "user updated successfully",
		Data:    updatedUser,
	})
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user_id from JWT middleware
	userIDStr, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{
			Type:    TypeFailure,
			Message: "unauthorized",
		})
		return
	}
	userID, _ := uuid.Parse(userIDStr)

	// Delete user from DB
	deletedUser, err := h.DBObject.DeleteUser(userID)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, MessageResponse{
			Type:    TypeFailure,
			Message: "failed to delete user",
		})
		return
	}

	respondJSON(w, http.StatusOK, MessageResponse{
		Type:    TypeSuccess,
		Message: "user deleted successfully",
		Data:    deletedUser,
	})
}

func (h *Handler) FetchUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user_id from JWT middleware
	userIDStr, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, MessageResponse{
			Type:    TypeFailure,
			Message: "unauthorized",
		})
		return
	}
	userID, _ := uuid.Parse(userIDStr)

	// Fetch user from DB
	user, err := h.DBObject.GetUserByID(userID)
	if err != nil {
		respondJSON(w, http.StatusNotFound, MessageResponse{
			Type:    TypeFailure,
			Message: "user not found",
		})
		return
	}

	respondJSON(w, http.StatusOK, MessageResponse{
		Type:    TypeSuccess,
		Message: "user fetched successfully",
		Data:    user,
	})
}
