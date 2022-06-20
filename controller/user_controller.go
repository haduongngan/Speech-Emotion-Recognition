package controller

import (
	"encoding/json"
	"net/http"
	"spser/model"
	"spser/service"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByUsername(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	LoginWithToken(w http.ResponseWriter, r *http.Request)
}

// GetAll returns all users
// @tags user-manager-apis
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

// GetByUsername gets user currently in table "users" with username
// @tags user-manager-apis
// @Summary get user with usn
// @Description input username => user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param username query string true "username"
// @Success 200 {object} model.Response
// @Router /user/wname [get]
func (c *userController) GetByUsername(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.Response
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest), 400)
		jsonResponse = &model.Response{
			Message: "Username is required",
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}
	user, err := c.userService.GetByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		jsonResponse = &model.Response{
			Message: err.Error(),
			Success: false,
		}
	} else {
		jsonResponse = &model.Response{
			Data:    user,
			Message: "OK",
			Success: true,
		}
	}
	render.JSON(w, r, jsonResponse)
}

// CreateUser creates an user with given data
// @tags user-manager-apis
// @Summary	creates new user
// @Description creates new user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param UserInfo body model.User true "User information"
// @Success 200 {object} model.CreateResponse
// @Router /user/create [post]
func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.CreateResponse
	var newUser *model.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest), 400)
		jsonResponse = &model.CreateResponse{
			Message: err.Error(),
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}

	userResponse, err := c.userService.CreateUser(newUser)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	jsonResponse = &model.CreateResponse{
		Username:    userResponse.Username,
		Role:        userResponse.Role,
		CompanyName: userResponse.CompanyName,
		Message:     "OK",
		Success:     true,
	}
}

// DeleteUser deletes user with UserID
// @tags user-manager-apis
// @Summary delete user
// @Description delete user
// @Accept json
// @Produce json
// @Param uid path integer true "User ID"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/delete/{uid} [delete]
func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.Response

	strID := chi.URLParam(r, "uid")
	uid, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		jsonResponse = &model.Response{
			Data:    nil,
			Message: "Error decoding request body:" + err.Error(),
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}

	if _, err := c.userService.DeleteUser(uid); err != nil {
		jsonResponse = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		jsonResponse = &model.Response{
			Message: "User deleted successfully!",
			Success: true,
		}
	}

	render.JSON(w, r, jsonResponse)
}

// Login log user in if they have valid credential
// @tags user-manager-apis
// @Summary log user in
// @Description log user in
// @Accept json
// @Produce json
// @Param LoginPayload body model.UserPayload true "username & password"
// @Success 200
// @Router /user/login [post]
func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.LoginResponse
	var loginDetail model.UserPayload

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginDetail); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		jsonResponse = &model.LoginResponse{
			Token:        "",
			RefreshToken: "",
			Message:      "Bad request: " + err.Error(),
			Code:         "400",
			Success:      false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}

	user, token, refreshToken, err := c.userService.LoginRequest(loginDetail.Username, loginDetail.Password)
	if err != nil {
		jsonResponse = &model.LoginResponse{
			Token:        token,
			RefreshToken: refreshToken,
			Message:      "Wrong username or password. Info:" + err.Error(),
			Code:         "400",
			Success:      false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}
	//user da bi xoa
	if user == nil {
		jsonResponse = &model.LoginResponse{
			Token:        "user has been deleted!",
			RefreshToken: "user has been deleted!",
			Message:      "user has been deleted!",
			Code:         "406",
			Success:      true,
		}
	} else {
		jsonResponse = &model.LoginResponse{
			Token:        token,
			RefreshToken: refreshToken,
			UserId:       user.Id,
			Role:         user.Role,
			Username:     user.Username,
			CompanyName:  user.CompanyName,
			Message:      "Logged in successfully as " + user.Role,
			Code:         "200",
			Success:      true,
		}
	}

	render.JSON(w, r, jsonResponse)
}

// LoginWithToken provides token each login attempt
// @tags user-manager-apis
// @Summary login user
// @Description login user, return new token string jwt
// @Accept json
// @Produce json
// @Param TokenPayload body controller.TokenPayload true "Insert your access token"
// @Success 200 {object} model.LoginResponse
// @Router /user/login/jwt [post]
func (c *userController) LoginWithToken(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.LoginResponse
	var refToken TokenPayload

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&refToken); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		jsonResponse = &model.LoginResponse{
			Token:   "",
			Message: err.Error(),
			Code:    "400",
			Success: false,
		}
	}

	user, accessToken, refreshToken, success, err := c.userService.LoginWithToken(refToken.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		jsonResponse = &model.LoginResponse{
			Token:   "",
			Message: err.Error(),
			Code:    "401",
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}
	if !success {
		jsonResponse = &model.LoginResponse{
			Token:   accessToken,
			Message: err.Error(),
			Code:    "400",
			Success: false,
		}
	} else {
		jsonResponse = &model.LoginResponse{
			Token:        accessToken,
			RefreshToken: refreshToken,
			UserId:       user.Id,
			Username:     user.Username,
			Role:         user.Role,
			CompanyName:  user.CompanyName,
			Message:      "jwt login successful!",
			Code:         "200",
			Success:      true,
		}
	}
	render.JSON(w, r, jsonResponse)
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
		Message: err.Error(),
		Success: false,
	})
	return
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
	w.WriteHeader(http.StatusBadRequest)
	render.JSON(w, r, &model.Response{
		Message: err.Error(),
		Success: false,
	})
	return
}

type TokenPayload struct {
	RefreshToken string `json:"refreshToken"`
}
