package router

import (
	"fmt"
	"net/http"

	"github.com/maiconDeSouza/read_log_api/internal/handlers"
)

type Routes struct {
	mux          *http.ServeMux
	userHandlers handlers.UserHandlersInterface
	versionAPI   string
}

func NewRoutes(mux *http.ServeMux, userHandlers handlers.UserHandlersInterface, versionAPI string) *Routes {
	return &Routes{mux: mux, userHandlers: userHandlers, versionAPI: versionAPI}
}

func (r *Routes) InitRoutes() {
	r.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	r.mux.HandleFunc(fmt.Sprintf("POST %s/user", r.versionAPI), r.userHandlers.RegisterUser)
	r.mux.HandleFunc(fmt.Sprintf("POST %s/user/login", r.versionAPI), r.userHandlers.Login)

}
