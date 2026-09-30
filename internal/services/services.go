package services

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maiconDeSouza/read_log_api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func generateJWT(user models.User, secretKey string) (string, *models.AppErr) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		msg := "Erro ao gerar JWT!"
		code := http.StatusInternalServerError
		return "", models.NewAppErr(msg, code, err)
	}

	return tokenString, nil
}

func generateCookie(jwt string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "session_token_user",
		Value:    jwt,
		Path:     "/ap1/v1/user",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	return cookie
}
