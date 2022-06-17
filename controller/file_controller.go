package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"spser/infrastructure"
	"spser/model"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type FileController interface {
	// UploadFile(w http.ResponseWriter, r *http.Request)
	DeleteFile(w http.ResponseWriter, r *http.Request)
	UploadMultipleFile(w http.ResponseWriter, r *http.Request)
}

type fileController struct {
	// contractService service.ContractService
}

// // Import storage file godoc
// // @tags file-manager-apis
// // @Summary Import file
// // @Description Import file
// // @Accept json
// // @Produce json
// // @Security ApiKeyAuth
// // @Param file formData file true "file data"
// // @Param id path integer true "id"
// // @Param type query string true "type" Enums(sound)
// // @Success 200 {object} model.Response
// // @Router /file/storage/{id} [post]
// func (c *fileController) UploadFile(w http.ResponseWriter, r *http.Request) {
// 	var res *model.Response

// 	queryValues := r.URL.Query()
// 	id, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		http.Error(w, http.StatusText(400), 400)
// 		res = &model.Response{
// 			Data:    nil,
// 			Message: "ie failed: " + err.Error(),
// 			Success: false,
// 		}
// 		render.JSON(w, r, res)
// 		return
// 	}
// 	typeStr := queryValues.Get("type")

// 	//25mb
// 	r.ParseMultipartForm(25 << 20)
// 	file, header, err := r.FormFile("file")
// 	if err != nil {
// 		res := &model.Response{
// 			Data:    nil,
// 			Message: "Error retrieving the file",
// 			Success: false,
// 		}
// 		render.JSON(w, r, res)
// 		return
// 	}
// 	defer file.Close()

// 	// filePath := infrastructure.GetSRFilePath() + strconv.Itoa(id) + "/"
// 	var filePath string
// 	switch typeStr {
// 	case "scientificResearch":
// 		filePath = infrastructure.GetSRFilePath() + strconv.Itoa(id) + "/"
// 	case "patent":
// 		filePath = infrastructure.GetPatentFilePath() + strconv.Itoa(id) + "/"
// 	case "solution":
// 		filePath = infrastructure.GetSolutionFilePath() + strconv.Itoa(id) + "/"
// 	case "researchProject":
// 		filePath = infrastructure.GetRPFilePath() + strconv.Itoa(id) + "/"
// 	case "monograph":
// 		filePath = infrastructure.GetMonographFilePath() + strconv.Itoa(id) + "/"
// 	case "productContract":
// 		filePath = infrastructure.GetProductContractFilePath() + strconv.Itoa(id) + "/"
// 	case "scientificResearchOther":
// 		filePath = infrastructure.GetSROtherFilePath() + strconv.Itoa(id) + "/"
// 	case "avatar":
// 		filePath = infrastructure.GetAvatarFilePath() + strconv.Itoa(id) + "/"
// 	case "studentResearch":
// 		filePath = infrastructure.GetStudentResearchFilePath() + strconv.Itoa(id) + "/"
// 	default:
// 		res := &model.Response{
// 			Data:    nil,
// 			Message: "Error In Type",
// 			Success: false,
// 		}
// 		render.JSON(w, r, res)
// 		return
// 	}
// 	if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 		os.Mkdir(filePath, 0755)
// 	}
// 	fileName := filePath + header.Filename
// 	tempFile, err := os.Create(fileName)
// 	if err != nil {
// 		res := &model.Response{
// 			Data:    nil,
// 			Message: "Error while create temp file. " + err.Error(),
// 			Success: false,
// 		}
// 		render.JSON(w, r, res)
// 		return
// 	}
// 	defer tempFile.Close()

// 	fileBytes, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		res := &model.Response{
// 			Data:    nil,
// 			Message: "Error while read content file. " + err.Error(),
// 			Success: false,
// 		}
// 		render.JSON(w, r, res)
// 		return
// 	}
// 	_, err = tempFile.Write(fileBytes)
// 	if err != nil {
// 		res := &model.Response{
// 			Data:    nil,
// 			Message: "Error while read content file. " + err.Error(),
// 			Success: false,
// 		}
// 		render.JSON(w, r, res)
// 		return
// 	}

// 	res = &model.Response{
// 		Data:    fileName[strings.Index(fileName, "/storage"):],
// 		Message: "DONE",
// 		Success: true,
// 	}
// 	render.JSON(w, r, res)
// }

// Delete file godoc
// @tags file-manager-apis
// @Summary Delete file
// @Description Delete file
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Param type query string true "type" Enums(sound)
// @Param name query string true "name"
// @Success 200 {object} model.Response
// @Router /file/storage/{id} [delete]
func (c *fileController) DeleteFile(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	queryValues := r.URL.Query()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "ie failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	fileName := queryValues.Get("name")
	typeStr := queryValues.Get("type")
	// filePath := infrastructure.GetSRFilePath() + strconv.Itoa(id) + "/" + fileName
	var filePath string
	switch typeStr {
	case "sound":
		filePath = infrastructure.GetStoragePath() + strconv.Itoa(id) + "/" + fileName
	default:
		res := &model.Response{
			Data:    nil,
			Message: "Error In Type",
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		res := &model.Response{
			Data:    nil,
			Message: "Error deleting file. " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return

	}

	res = &model.Response{
		Data:    "",
		Message: "delete file successful.",
		Success: true,
	}
	render.JSON(w, r, res)
}

// // UploadMultipleFile storage file godoc
// // @tags file-manager-apis
// // @Summary Import file
// // @Description Import file
// // @Accept json
// // @Produce json
// // @Param id path integer true "id"
// // @Param file formData file true "file data"
// // @Param hvSciRes query string true "have file or not" Enums(yes, no)
// // @Param hvPatent query string true "have file or not" Enums(yes, no)
// // @Param hvSolution query string true "have file or not" Enums(yes, no)
// // @Param hvResearchProject query string true "have file or not" Enums(yes, no)
// // @Param hvMonograph query string true "have file or not" Enums(yes, no)
// // @Param hvProductContract query string true "have file or not" Enums(yes, no)
// // @Param hvScientificResearchOther query string true "have file or not" Enums(yes, no)
// // @Param hvAvatar query string true "have file or not" Enums(yes, no)
// // @Success 200 {object} model.Response
// // @Router /file/storage/multi/{id} [post]
// func (c *fileController) UploadMultipleFile(w http.ResponseWriter, r *http.Request) {
// 	var response *model.Response
// 	var fileUrl []string
// 	queryValues := r.URL.Query()
// 	id, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		http.Error(w, http.StatusText(400), 400)
// 		response = &model.Response{
// 			Data:    nil,
// 			Message: "id failed: " + err.Error(),
// 			Success: false,
// 		}
// 		render.JSON(w, r, response)
// 		return
// 	}
// 	if err := r.ParseMultipartForm(25 << 20); err != nil {
// 		fmt.Fprintln(w, err)
// 		return
// 	}
// 	var queryArr []string
// 	queryArr = append(queryArr, queryValues.Get("hvSciRes"))
// 	queryArr = append(queryArr, queryValues.Get("hvPatent"))
// 	queryArr = append(queryArr, queryValues.Get("hvSolution"))
// 	queryArr = append(queryArr, queryValues.Get("hvResearchProject"))
// 	queryArr = append(queryArr, queryValues.Get("hvMonograph"))
// 	queryArr = append(queryArr, queryValues.Get("hvProductContract"))
// 	queryArr = append(queryArr, queryValues.Get("hvScientificResearchOther"))
// 	queryArr = append(queryArr, queryValues.Get("hvAvatar"))
// 	log.Println(queryArr)

// 	formData := r.MultipartForm
// 	files := formData.File["file"]
// 	indexArr := make([]int, len(files))

// 	for i := range files {
// 		file, err := files[i].Open()

// 		if err != nil {
// 			fmt.Fprintln(w, err)
// 		}
// 		defer file.Close()
// 		var filePath string
// 		switch queryArr[i] {
// 		case "avatar":
// 			filePath = infrastructure.GetProfilePath() + strconv.Itoa(id) + "/"
// 		default:
// 			res := &model.Response{
// 				Data:    nil,
// 				Message: "Error In Type",
// 				Success: false,
// 			}
// 			render.JSON(w, r, res)
// 			return
// 		}
// 		if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 			os.Mkdir(filePath, 0755)
// 		}
// 		fileName := filePath + files[i].Filename
// 		tempFile, err := os.Create(fileName)
// 		if err != nil {
// 			res := &model.Response{
// 				Data:    nil,
// 				Message: "Error while create temp file. " + err.Error(),
// 				Success: false,
// 			}
// 			render.JSON(w, r, res)
// 			return
// 		}
// 		defer tempFile.Close()

// 		fileBytes, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			res := &model.Response{
// 				Data:    nil,
// 				Message: "Error while read content file. " + err.Error(),
// 				Success: false,
// 			}
// 			render.JSON(w, r, res)
// 			return
// 		}
// 		_, err = tempFile.Write(fileBytes)
// 		if err != nil {
// 			res := &model.Response{
// 				Data:    nil,
// 				Message: "Error while read content file. " + err.Error(),
// 				Success: false,
// 			}
// 			render.JSON(w, r, res)
// 			return
// 		}

// 		fileUrl = append(fileUrl, fileName[strings.Index(fileName, "/storage"):])
// 	}
// 	response = &model.Response{
// 		Data:    fileUrl,
// 		Message: "DONE",
// 		Success: true,
// 	}
// 	render.JSON(w, r, response)
// }

// UploadMultipleFile storage file godoc
// @tags file-manager-apis
// @Summary Import file
// @Description Import file
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Param file formData file true "file data"
// @Param type query string true "type" Enums(sound)
// @Success 200 {object} model.Response
// @Router /file/storage/multi/{id} [post]
func (c *fileController) UploadMultipleFile(w http.ResponseWriter, r *http.Request) {
	var response *model.Response
	var fileUrl []string
	queryValues := r.URL.Query()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		response = &model.Response{
			Data:    nil,
			Message: "id failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, response)
		return
	}
	if err := r.ParseMultipartForm(25 << 20); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	typeStr := queryValues.Get("type")
	formData := r.MultipartForm
	files := formData.File["file"]
	log.Println(len(files))
	for i := range files {
		file, err := files[i].Open()

		log.Println(files[i].Filename)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		defer file.Close()
		var filePath string
		switch typeStr {
		case "sound":
			filePath = infrastructure.GetStoragePath() + strconv.Itoa(id) + "/"
		default:
			res := &model.Response{
				Data:    nil,
				Message: "Error In Type",
				Success: false,
			}
			render.JSON(w, r, res)
			return
		}
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			os.Mkdir(filePath, 0755)
		}
		fileName := filePath + files[i].Filename
		log.Println(fileName)
		tempFile, err := os.Create(fileName)
		if err != nil {
			res := &model.Response{
				Data:    nil,
				Message: "Error while create temp file. " + err.Error(),
				Success: false,
			}
			render.JSON(w, r, res)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			res := &model.Response{
				Data:    nil,
				Message: "Error while read content file. " + err.Error(),
				Success: false,
			}
			render.JSON(w, r, res)
			return
		}
		_, err = tempFile.Write(fileBytes)
		if err != nil {
			res := &model.Response{
				Data:    nil,
				Message: "Error while read content file. " + err.Error(),
				Success: false,
			}
			render.JSON(w, r, res)
			return
		}

		fileUrl = append(fileUrl, fileName[strings.Index(fileName, "/storage"):])
	}
	response = &model.Response{
		Data:    fileUrl,
		Message: "DONE",
		Success: true,
	}
	render.JSON(w, r, response)
}

func NewFileController() FileController {
	return &fileController{}
}
