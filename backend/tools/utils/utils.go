package utils

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// IsValidEmail checks if the provided string is a valid email format.
func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}


var jwtSecret = []byte("your-super-secret-key") // 🔒 Replace with env var in production

const tokenExpiryDuration = 72 * time.Hour // 3 days validity

// GenerateJWT generates a JWT string with the userID as a claim.
func GenerateJWT(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(tokenExpiryDuration).Unix(),
        "iat":     time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateJWT validates the JWT string and returns the userID if valid.
func ValidateJWT(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtSecret, nil
    })

    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID, ok := claims["user_id"].(string)
        if !ok || userID == "" {
            return "", errors.New("user_id not found in token")
        }
        return userID, nil
    }

    return "", errors.New("invalid token")
}
