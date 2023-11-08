package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	commonHTTP "{{ module_name }}/http"
	"{{ module_name }}/users/model"
)

type usersService interface {
	UpdateUser(ctx context.Context, userID string, req *model.UserUpdateRequest) (*model.User, error)
}

type userController struct {
	usersService usersService
}

func NewUserController(usersService usersService, r *chi.Mux) {

	userController := &userController{
		usersService: usersService,
	}
	r.Group(func(r chi.Router) {
		// authn'ed routes
		r.Mount("/users", authRouter(userController))

	},
	)
}

func authRouter(userController *userController) *chi.Mux {
	r := chi.NewMux()
	r.Patch("/{userId}", userController.updateUserHandler)
	return r
}

func (ctrl *userController) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req = new(model.UserUpdateRequest)
	userID := chi.URLParam(r, "userId")
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		commonHTTP.Error(w, r, http.StatusBadRequest, err)
		return
	}
	_, err = ctrl.usersService.UpdateUser(r.Context(), userID, req)
	if err != nil {
		commonHTTP.Error(w, r, http.StatusInternalServerError, err)
		return
	}
	commonHTTP.ResponseOK(w, r, nil)
}
