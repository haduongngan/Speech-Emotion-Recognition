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

type segmentController struct {
	segmentService service.SegmentService
}

type SegmentController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	GetByCallId(w http.ResponseWriter, r *http.Request)
	CreateSegment(w http.ResponseWriter, r *http.Request)
	DeleteSegment(w http.ResponseWriter, r *http.Request)
	GetEmotion(w http.ResponseWriter, r *http.Request)
}

// GetAll returns all segments
// @tags segment-manager-apis
// @summary Get all segments
// @description input none => []Segment
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /segment/all [get]
func (c *segmentController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	segments, err := c.segmentService.GetAll()
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    segments,
		Message: "Success",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetById returns segment with Id
// @tags segment-manager-apis
// @summary Get segment with Id
// @description input id => Segment
// @Accept json
// @Produce json
// @Param id path integer true "Segment Id"
// @Success 200 {object} model.Response
// @Router /segment/{id} [get]
func (c *segmentController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}
	segment, err := c.segmentService.GetById(id)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    segment,
		Message: "Success",
		Success: true,
	}
	render.JSON(w, r, res)
}

// GetByCallId returns segments with CallId
// @tags segment-manager-apis
// @summary Get segment with CallId
// @description input callId => []Segment
// @Accept json
// @Produce json
// @Param callId path integer true "Call Id"
// @Success 200 {object} model.Response
// @Router /segment/call/{callId} [get]
func (c *segmentController) GetByCallId(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	strCallId := chi.URLParam(r, "callId")
	callId, err := strconv.Atoi(strCallId)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	segments, err := c.segmentService.GetByCallId(callId)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    segments,
		Message: "Success",
		Success: true,
	}

	render.JSON(w, r, res)
}

// CreateSegment create one segment into db
// @tags segment-manager-apis
// @summary Create segment
// @description input segmentInfo => Segment
// @Accept json
// @Produce json
// @Param segmentInfo body model.Segment true "Segment info"
// @Success 200 {object} model.Response
// @Router /segment/create [post]
func (c *segmentController) CreateSegment(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var segments *model.Segment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&segments); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	err := c.segmentService.CreateSegment(segments)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "CREATE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// DeleteSegment deletes one segment from db with id
// @tags segment-manager-apis
// @summary deletes one segment from db with id
// @description input id => deleted segment
// @Accept json
// @Produce json
// @Param id path integer true "Segment Id"
// @Success 200 {object} model.Response
// @Router /segment/delete/{id} [delete]
func (c *segmentController) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	err = c.segmentService.DeleteSegment(id)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Message: "DELETE SUCCESS",
		Success: true,
	}

	render.JSON(w, r, res)
}

// GetEmotion returns emotion of segment
// @tags segment-manager-apis
// @summary Get segment emotion with Id
// @description input id => emotion
// @Accept json
// @Produce json
// @Param id path integer true "Segment Id"
// @Success 200 {object} model.Response
// @Router /segment/emo/{id} [get]
func (c *segmentController) GetEmotion(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	emotion, err := c.segmentService.GetEmotion(id)
	if err != nil {
		internalServerErrorResponse(w, r, err)
		return
	}

	res = &model.Response{
		Data:    emotion,
		Message: "Success",
		Success: true,
	}

	render.JSON(w, r, res)
}

func NewSegmentController() SegmentController {
	return &segmentController{
		segmentService: service.NewSegmentService(),
	}
}
