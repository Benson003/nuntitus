package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Benson003/nuntius/database"
	"github.com/Benson003/nuntius/tools/hashing"
)

// Login request format
type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

// Sign up request format
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        respondJSON(w, http.StatusBadRequest, MessageResponse{
            Type:    TypeFailure,
            Message: "invalid json body",
        })
        return
    }

    if req.UsernameOrEmail == "" || req.Password == "" {
        respondJSON(w, http.StatusBadRequest, MessageResponse{
            Type:    TypeFailure,
            Message: "username/email and password required",
        })
        return
    }

    var user database.UserTable
    userEmail, errEmail := h.DBObject.GetUserByEmail(req.UsernameOrEmail)
    userUsername, errUsername := h.DBObject.GetUserByUsername(req.UsernameOrEmail)

    if errEmail == nil {
        user = userEmail
    } else if errUsername == nil {
        user = userUsername
    } else {
        respondJSON(w, http.StatusNotFound, MessageResponse{
            Type:    TypeFailure,
            Message: "failed to find user",
        })
        return
    }

    if err := hashing.VerifyPassword(user.PasswordHash, req.Password); err != nil {
        respondJSON(w, http.StatusUnauthorized, MessageResponse{
            Type:    TypeFailure,
            Message: "invalid credentials",
        })
        return
    }

    tokenString, err := GenerateJWT(user.UserID)
    if err != nil {
        respondJSON(w, http.StatusInternalServerError, MessageResponse{
            Type:    TypeFailure,
            Message: "failed to generate token",
        })
        return
    }

    respondJSON(w, http.StatusOK, MessageResponse{
        Type:    TypeSuccess,
        Message: "login successful",
        Data: map[string]interface{}{
            "token": tokenString,
        },
    })
}


func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}
	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)
	email := strings.TrimSpace(req.Email)
	// Basic input validation
	if username == "" || email == "" || password == "" {
		http.Error(w, "username, email, and password are required", http.StatusBadRequest)
		return
	}
	hash, err := hashing.HashPassword(password)
	if err != nil {
		http.Error(w, "couldn't hash password", http.StatusInternalServerError)
		return
	}
	_, err = h.DBObject.CreateUser(username, hash, email)
	if err != nil {
		http.Error(w, "couldn't create user", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusCreated, MessageResponse{
		Type:    TypeSuccess,
		Message: "user created successfully",
	})
}

func (h *Handler) Revalidation(w http.ResponseWriter, r *http.Request) {
	//Revalidation
}
