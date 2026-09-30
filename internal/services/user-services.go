package services

import (
	"net/http"
	"net/mail"

	"github.com/maiconDeSouza/read_log_api/internal/models"
	"github.com/maiconDeSouza/read_log_api/internal/repositories"
)

type UserServicesInterface interface {
	RegisterUser(newUser models.UseRequest) (*models.User, *models.AppErr)
}

type UserService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(newUser models.UseRequest) (*models.User, *models.AppErr) {
	if len(newUser.Nickname) < 3 || len(newUser.Nickname) > 22 {
		msg := "O seu nickname precisa conter entre 3 à 22 caracteres"
		code := http.StatusBadRequest
		return nil, models.NewAppErr(msg, code, nil)
	}

	if _, err := mail.ParseAddress(newUser.Email); err != nil {
		msg := "email inválido!"
		code := http.StatusBadRequest
		return nil, models.NewAppErr(msg, code, err)
	}

	if newUser.Password != newUser.RepeatPassword {
		msg := "Vocẽ precisa digitar a mesma senha nos dois campos"
		code := http.StatusBadRequest
		return nil, models.NewAppErr(msg, code, nil)
	}

	passwordHash, err := hashPassword(newUser.Password)
	if err != nil {
		msg := "Erro no ervidor!"
		code := http.StatusInternalServerError
		return nil, models.NewAppErr(msg, code, err)
	}

	user := models.User{
		Nickname: newUser.Nickname,
		Email:    newUser.Email,
		Password: passwordHash,
	}

	result, appErr := s.repo.CreateUser(user)
	if appErr != nil {
		return result, appErr
	}
	return result, appErr
}
