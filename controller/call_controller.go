package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"spser/model"
	"spser/service"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type callController struct {
	callService service.CallService
}

type CallController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	CreateCall(w http.ResponseWriter, r *http.Request)
	CreateMultiCall(w http.ResponseWriter, r *http.Request)
	UpdateCall(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	DeleteCall(w http.ResponseWriter, r *http.Request)
}

// GetAll returns all calls
// @tags call-manager-apis
// @Summary Get all calls
// @Description input none => []Call
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /call/all [get]
func (c *callController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	calls, err := c.callService.GetAll()
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

// CreateCall creates a new call
// @tags call-manager-apis
// @Summary Create a call
// @Description input Call => Call
// @Accept json
// @Produce json
// @Param call body model.Call true "Call"
// @Success 200 {object} model.Response
// @Router /call/create [post]
func (c *callController) CreateCall(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	var msgMapTemplate interface{}
	if err := json.Unmarshal(body, &msgMapTemplate); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	msgMap := msgMapTemplate.(map[string]interface{})
	log.Println(msgMap)

	var res *model.Response
	var call model.Call
	call.Duration = (msgMap["dur"].(string))
	call.Phone = (msgMap["phone"].(string))
	staff := msgMap["staff"].(map[string]interface{})
	call.StaffEmotion = staff["feel"].(string)
	customer := msgMap["customer"].(map[string]interface{})
	call.CustomerEmotion = customer["feel"].(string)
	call.StartTime = time.Now()
	staffIdString := msgMap["staffId"].(string)
	staffId, err := strconv.Atoi(staffIdString)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	call.StaffId = staffId
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&call); err != nil {
	// 	badRequestResponse(w, r, err)
	// 	return
	// }

	if err := c.callService.CreateCall(&call); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "CREATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// CreateMultiCall creates a new call
// @tags call-manager-apis
// @Summary Create multiple calls
// @Description input []Call => add []Call
// @Accept json
// @Produce json
// @Param calls body []model.Call true "Calls info"
// @Success 200 {object} model.Response
// @Router /call/multi/create [post]
func (c *callController) CreateMultiCall(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var calls []model.Call
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&calls); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.callService.CreateMultiCall(calls); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "CREATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// UpdateCall updates a call
// @tags call-manager-apis
// @Summary Update a call
// @Description input Call => update Call
// @Accept json
// @Produce json
// @Param call body model.Call true "Call"
// @Success 200 {object} model.Response
// @Router /call/update [put]
func (c *callController) UpdateCall(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var call model.Call
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&call); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := c.callService.UpdateCall(&call); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "UPDATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetById returns a call by id
// @tags call-manager-apis
// @Summary Get a call by id
// @Description input id => Call
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} model.Response
// @Router /call/{id} [get]
func (c *callController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	call, err := c.callService.GetById(id)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    call,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

// DeleteCall deletes a call
// @tags call-manager-apis
// @Summary Delete a call
// @Description input id => delete Call
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} model.Response
// @Router /call/delete/{id} [delete]
func (c *callController) DeleteCall(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	if err := c.callService.DeleteCall(id); err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "DELETE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

func NewCallController() CallController {
	return &callController{
		callService: service.NewCallService(),
	}
}
