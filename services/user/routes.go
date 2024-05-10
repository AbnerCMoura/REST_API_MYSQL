package user

import (
	"github.com/AbnerCMoura/REST_API_MYSQL/types"
	"github.com/AbnerCMoura/REST_API_MYSQL/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, req *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, req *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(req, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
}
