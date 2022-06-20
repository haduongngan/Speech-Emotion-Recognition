package controller

import (
	"encoding/json"
	"net/http"
	"spser/middleware"
	"spser/model"
	"spser/service"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type staffController struct {
	staffService service.StaffService
}
type StaffController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllCall(w http.ResponseWriter, r *http.Request)
	FilterCallInTime(w http.ResponseWriter, r *http.Request)
	CreateStaff(w http.ResponseWriter, r *http.Request)
	UpdateStaff(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	DeleteStaff(w http.ResponseWriter, r *http.Request)
}

// GetAll returns all Staff
// @tags staff-manager-apis
// @Summary Get all Staffs
// @Description Get all Staffs
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /staff/all [get]
func (c *staffController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	staffs, err := c.staffService.GetAll()
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    staffs,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetAllCall returns all calls of an user with id
// @tags staff-manager-apis
// @Summary Get all calls of an user with id
// @Description Get all calls of an user with id
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /staff/calls [put]
func (c *staffController) GetAllCall(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	reqToken := r.Header.Get("Authorization")
	user, err := middleware.GetClaimsData(reqToken)

	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	calls, err := c.staffService.GetAllCall(user.Id)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    calls,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// FilterCallInTime returns all calls in timeframe of an user with id
// @tags staff-manager-apis
// @Summary Get all calls in timeframe of an user with id
// @Description model.CallTimeFilterPayload => []Call
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body model.StaffCallFilterPayload true "payload"
// @Success 200 {object} model.Response
// @Router /staff/filter/calls [put]
func (c *staffController) FilterCallInTime(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var payload model.StaffCallFilterPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	reqToken := r.Header.Get("Authorization")
	user, err := middleware.GetClaimsData(reqToken)

	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	payload.UserId = user.Id
	calls, err := c.staffService.FilterCallInTime(&payload)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    calls,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// CreateStaff create one Staff
// @tags staff-manager-apis
// @Summary Get all Staffs
// @Description Get all Staffs
// @Accept json
// @Produce json
// @Param Staff body model.Staff true "Staff"
// @Success 200 {object} model.Response
// @Router /staff/create [post]
func (c *staffController) CreateStaff(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var staff model.Staff

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&staff); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.staffService.CreateStaff(&staff); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "CREATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// UpdateStaff update one Staff
// @tags staff-manager-apis
// @Summary Update one Staff
// @Description Update one Staff
// @Accept json
// @Produce json
// @Param staffName body model.StaffNameUpdate true "Staff"
// @Success 200 {object} model.Response
// @Router /staff/update [put]
func (c *staffController) UpdateStaff(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var staff model.StaffNameUpdate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&staff); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.staffService.UpdateStaff(&staff); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "UPDATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetById returns one Staff
// @tags staff-manager-apis
// @Summary Get one Staff
// @Description Get one Staff
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} model.Response
// @Router /staff/{id} [get]
func (c *staffController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	staff, err := c.staffService.GetById(id)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    staff,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// DeleteCustomer delete one Staff
// @tags staff-manager-apis
// @Summary Delete one Staff
// @Description Delete one Staff
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} model.Response
// @Router /staff/delete/{id} [delete]
func (c *staffController) DeleteStaff(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.staffService.DeleteStaff(id); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "DELETE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

func NewStaffController() StaffController {
	return &staffController{
		staffService: service.NewStaffService(),
	}
}
