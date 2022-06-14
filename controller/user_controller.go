package controller

import (
	"net/http"
	"spser/model"
	"spser/service"

	"github.com/go-chi/render"
)

type userController struct {
	userService service.UserService
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
	users, err := c.userService.GetAll()
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    users,
		Message: "OK",
		Success: true,
	}
	render.JSON(w, r, res)
}

func NewUserController() userController {
	return userController{
		userService: service.NewUserService(),
	}
}

func internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	w.WriteHeader(http.StatusInternalServerError)
	render.JSON(w, r, &model.Response{
		Data:    nil,
		Message: err.Error(),
		Success: false,
	})
	return
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
	w.WriteHeader(http.StatusBadRequest)
	render.JSON(w, r, &model.Response{
		Data:    nil,
		Message: err.Error(),
		Success: false,
	})
	return
}
