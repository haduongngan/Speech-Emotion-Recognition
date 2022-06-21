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

type customerController struct {
	customerService service.CustomerService
}
type CustomerController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllCall(w http.ResponseWriter, r *http.Request)
	FilterCallInTime(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}

// GetAll returns all customers
// @tags customer-manager-apis
// @Summary Get all customers
// @Description Get all customers
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /customer/all [get]
func (c *customerController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	customers, err := c.customerService.GetAll()
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    customers,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// CreateCustomer create one customer
// @tags customer-manager-apis
// @Summary Get all customers
// @Description Get all customers
// @Accept json
// @Produce json
// @Param customer body model.Customer true "customer"
// @Success 200 {object} model.Response
// @Router /customer/create [post]
func (c *customerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var customer model.Customer

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.customerService.CreateCustomer(&customer); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "CREATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// UpdateCustomer update one customer
// @tags customer-manager-apis
// @Summary Update one customer
// @Description Update one customer
// @Accept json
// @Produce json
// @Param customer body model.CustomerPhoneUpdate true "customer id and phone info"
// @Success 200 {object} model.Response
// @Router /customer/update [put]
func (c *customerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var customer model.CustomerPhoneUpdate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.customerService.UpdateCustomer(&customer); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "UPDATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetById returns one customer
// @tags customer-manager-apis
// @Summary Get one customer
// @Description Get one customer
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} model.Response
// @Router /customer/{id} [get]
func (c *customerController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	customer, err := c.customerService.GetById(id)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    customer,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// DeleteCustomer delete one customer
// @tags customer-manager-apis
// @Summary Delete one customer
// @Description Delete one customer
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} model.Response
// @Router /customer/delete/{id} [delete]
func (c *customerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.customerService.DeleteCustomer(id); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "DELETE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetAllCall returns all calls of an user with id
// @tags customer-manager-apis
// @Summary Get all calls of an user with id
// @Description Get all calls of an user with id
// @Accept json
// @Produce json
// @Param phone body model.HistoryPayload true "Customer's phone number"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /customer/calls [put]
func (c *customerController) GetAllCall(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	// reqToken := r.Header.Get("Authorization")
	// user, err := middleware.GetClaimsData(reqToken)

	// if err != nil {
	// 	badRequestResponse(w, r, err)
	// 	return
	// }
	var payload model.HistoryPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}
	calls, err := c.customerService.GetAllCall(payload.Phone)
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
// @tags customer-manager-apis
// @Summary Get all calls in timeframe of an user with id
// @Description model.CallTimeFilterPayload => []Call
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body model.CallTimeFilterPayload true "payload"
// @Success 200 {object} model.Response
// @Router /customer/calls/filter [put]
func (c *customerController) FilterCallInTime(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var payload model.CallTimeFilterPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	// reqToken := r.Header.Get("Authorization")
	// user, err := middleware.GetClaimsData(reqToken)

	// if err != nil {
	// 	badRequestResponse(w, r, err)
	// 	return
	// }

	// payload.UserId = user.Id
	calls, err := c.customerService.FilterCallInTime(&payload)
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
func NewCustomerController() CustomerController {
	return &customerController{
		customerService: service.NewCustomerService(),
	}
}
