package controller

import (
	"net/http"
	"spser/model"

	"github.com/go-chi/render"
)

type userController struct {
}

type UserController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

// GetAll returns all users
// @tags user-apis
// @Summary returns all users
// @Description returns all users
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /user/all [get]
func (c *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	render.JSON(w, r, res)
}

func NewUserController() userController {
	return userController{}
}
