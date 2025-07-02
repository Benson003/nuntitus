package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Benson003/nuntius/database"
	"github.com/golang-jwt/jwt/v5"
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

func GenerateJWT(userID uuid.UUID) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID.String(),                       // subject claim as userID
        "exp": time.Now().Add(JWTExpiryTime).Unix(),  // expiry
        "iat": time.Now().Unix(),                     // issued at
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
