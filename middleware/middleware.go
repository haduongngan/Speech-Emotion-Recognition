package middleware

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"spser/controller"

// 	"github.com/go-chi/render"
// 	"gopkg.in/go-playground/validator.v9"
// )

// func validate(w http.ResponseWriter, payload interface{}) error {
// 	v := validator.New()

// 	errs := v.Struct(payload)

// 	if errs != nil {
// 		validationErrorResponse(w, errs)
// 		return errs
// 	}
// 	return nil
// }

// func ValidatePayload(payload interface{}) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			err := json.NewDecoder(r.Body).Decode(&payload)
// 			if err != nil {
// 				w.WriteHeader(http.StatusBadRequest)
// 				http.Error(w, http.StatusText(400), 400)
// 				res := &controller.Response{
// 					Data:    nil,
// 					Message: "Bad request: " + err.Error(),
// 					Success: false,
// 				}
// 				render.JSON(w, r, res)
// 				return
// 			}

// 			defer r.Body.Close()

// 			err = validate(w, payload)
// 			if err != nil {
// 				return
// 			}

// 			p := payload

// 			ctx := context.WithValue(r.Context(), "payload", p)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		})
// 	}
// }

// func validationErrorResponse(w http.ResponseWriter, err error) {
// 	errResponse := make([]string, 0)

// 	for _, e := range err.(validator.ValidationErrors) {
// 		errResponse = append(errResponse, fmt.Sprint(e))
// 	}

// 	response := map[string][]string{"errors": errResponse}
// 	render.JSON(w, nil, response)
// }
