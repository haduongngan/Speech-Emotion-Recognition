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

type employeeController struct {
	employeeService service.EmployeeService
}
type EmployeeController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllCall(w http.ResponseWriter, r *http.Request)
	FilterCallInTime(w http.ResponseWriter, r *http.Request)
	CreateEmployee(w http.ResponseWriter, r *http.Request)
	UpdateEmployee(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	DeleteEmployee(w http.ResponseWriter, r *http.Request)
}

// GetAll returns all Employee
// @tags employee-manager-apis
// @Summary Get all employees
// @Description Get all employees
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /employee/all [get]
func (c *employeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	employees, err := c.employeeService.GetAll()
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    employees,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetAllCall returns all calls of an user with id
// @tags employee-manager-apis
// @Summary Get all calls of an user with id
// @Description Get all calls of an user with id
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /employee/calls [get]
func (c *employeeController) GetAllCall(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	reqToken := r.Header.Get("Authorization")
	user, err := middleware.GetClaimsData(reqToken)

	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	calls, err := c.employeeService.GetAllCall(user.Id)
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
// @tags employee-manager-apis
// @Summary Get all calls in timeframe of an user with id
// @Description model.CallTimeFilterPayload => []Call
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body model.CallTimeFilterPayload true "payload"
// @Success 200 {object} model.Response
// @Router /employee/filter/calls [put]
func (c *employeeController) FilterCallInTime(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var payload model.CallTimeFilterPayload
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
	calls, err := c.employeeService.FilterCallInTime(&payload)
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

// CreateEmployee create one employee
// @tags employee-manager-apis
// @Summary Get all employees
// @Description Get all employees
// @Accept json
// @Produce json
// @Param employee body model.Employee true "employee"
// @Success 200 {object} model.Response
// @Router /employee/create [post]
func (c *employeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var employee model.Employee

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.employeeService.CreateEmployee(&employee); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "CREATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// UpdateEmployee update one employee
// @tags employee-manager-apis
// @Summary Update one employee
// @Description Update one employee
// @Accept json
// @Produce json
// @Param employee body model.Employee true "employee"
// @Success 200 {object} model.Response
// @Router /employee/update [put]
func (c *employeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var employee model.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.employeeService.UpdateEmployee(&employee); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "UPDATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetById returns one employee
// @tags employee-manager-apis
// @Summary Get one employee
// @Description Get one employee
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} model.Response
// @Router /employee/{id} [get]
func (c *employeeController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	employee, err := c.employeeService.GetById(id)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    employee,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// DeleteCustomer delete one employee
// @tags employee-manager-apis
// @Summary Delete one employee
// @Description Delete one employee
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} model.Response
// @Router /employee/delete/{id} [delete]
func (c *employeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.employeeService.DeleteEmployee(id); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "DELETE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

func NewEmployeeController() EmployeeController {
	return &employeeController{
		employeeService: service.NewEmployeeService(),
	}
}
