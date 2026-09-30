package repositories

import (
	"net/http"

	"github.com/maiconDeSouza/read_log_api/internal/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(user models.User) (*models.User, *models.AppErr)
	FindUserByEmail(email string) (*models.User, *models.AppErr)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user models.User) (*models.User, *models.AppErr) {
	re := r.db.Create(&user)

	if re.Error != nil {
		msg := "Erro ao salvar o usuario no banco de dados"
		code := http.StatusInternalServerError
		return nil, models.NewAppErr(msg, code, re.Error)
	}

	return &user, nil
}

func (r *UserRepo) FindUserByEmail(email string) (*models.User, *models.AppErr) {
	user := models.User{}

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		msg := "Erro no servidor!"
		code := http.StatusInternalServerError
		return nil, models.NewAppErr(msg, code, err)
	}

	return &user, nil
}
