package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Benson003/nuntius/database"

	"github.com/google/uuid"
)

var jwtSecret = []byte("your_secret_here")
const (
	JWTExpiryTime time.Duration = time.Hour * 24 * 31 // 31 days for theh token to expire
	TypeSuccess   string        = "success"
	TypeFailure   string        = "error"
	TypeInfo      string        = "message"
)

type Handler struct {
	DBObject *database.DBObject
}

type JWT struct {
	AuthToken uuid.UUID `json:"auth_token"`
	Timeout   int64     `json:"expiry"`
}

type MessageResponse struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(payload)
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"` // optional new username
	Email    string `json:"email,omitempty"`    // optional new email
	Password string `json:"password,omitempty"` // optional new password
}

type CreateBlogRequest struct {
    Title       string    `json:"title"`
    Summary     string    `json:"summary"`
    PublishTime time.Time `json:"publish_time"`
    Visibility  bool      `json:"visibility"`
}

type UpdateBlogRequest struct {
    Title       string    `json:"title,omitempty"`
    Summary     string    `json:"summary,omitempty"`
    PublishTime time.Time `json:"publish_time,omitempty"`
    Visibility  *bool     `json:"visibility,omitempty"`
}
