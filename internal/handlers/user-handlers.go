package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/maiconDeSouza/read_log_api/internal/models"
	"github.com/maiconDeSouza/read_log_api/internal/services"
)

type UserHandlersInterface interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
}

type UserHandlers struct {
	services services.UserServicesInterface
}

func NewUserHandlers(services services.UserServicesInterface) *UserHandlers {
	return &UserHandlers{services: services}
}

func (h *UserHandlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.UseRequest{}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		msg := "Erro no json enviado!"
		code := http.StatusBadRequest
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(models.NewAppErr(msg, code, err))
		return
	}

	user, appErr := h.services.RegisterUser(newUser)
	if appErr != nil {
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
