package router

import (
	"log"
	"net/http"
	"os"
	"spser/controller"
	_ "spser/docs"
	"spser/infrastructure"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrLog  = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	acceptCors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  checkOrigin,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(acceptCors.Handler)
	// swagger route
	r.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(infrastructure.GetHTTPSwagger()),
	))

	userController := controller.NewUserController()
	r.Route("/api/v1", func(router chi.Router) {
		// // public routes
		// router.Post("/user/login", userController.Login)
		// router.Post("/user/login/jwt", userController.LoginWithToken)

		router.Get("/user/all", userController.GetAll)
		// protected routes
		router.Group(func(protectedRoute chi.Router) {
			// Declare middleware
			protectedRoute.Use(jwtauth.Verifier(infrastructure.GetEncodeAuth()))
			protectedRoute.Use(jwtauth.Authenticator)

			// public routers

			// //---------------User routes--------------------------------
			// protectedRoute.Route("/user", func(subRoute chi.Router) {
			// 	subRoute.Post("/create", userController.CreateUser)
			// 	subRoute.Get("/all", userController.GetAll)
			// 	subRoute.Delete("/delete/{uid}", userController.DeleteUser)
			// 	subRoute.Get("/wname", userController.GetByUsername)
			// 	subRoute.Put("/setPerm", userController.SetPermission)
			// 	subRoute.Get("/getchild", userController.GetChildUser)
			// 	subRoute.Get("/getcitizen", userController.GetChildCitizen)
			// 	subRoute.Get("/progress", userController.GetCensusProgress)
			// 	subRoute.Put("/set_progress", userController.SetProgress)
			// 	subRoute.Get("/sex_chart", userController.GetSexChart)
			// 	subRoute.Get("/age_chart", userController.GetAgeChart)
			// })

		})
	})

	return r
}
