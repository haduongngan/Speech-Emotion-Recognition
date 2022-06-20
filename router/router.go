package router

import (
	"log"
	"net/http"
	"os"
	"spser/controller"
	"spser/infrastructure"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	_ "spser/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLog  = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// Router Root Router
func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(time.Duration(5 * time.Second)))
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// api swagger for develope mode
	r.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(infrastructure.GetHTTPSwagger()),
		httpSwagger.DocExpansion("none"),
	))

	//declare controller
	userController := controller.NewUserController()
	segmentController := controller.NewSegmentController()
	fileController := controller.NewFileController()

	r.Route("/api/v1", func(router chi.Router) {
		// Public routes

		//----- user routes-----
		router.Get("/user/all", userController.GetAll)
		router.Get("/user/wname", userController.GetByUsername)
		router.Post("/user/create", userController.CreateUser)
		router.Delete("/user/delete/{uid}", userController.DeleteUser)
		router.Post("/user/login", userController.Login)
		router.Post("/user/login/jwt", userController.LoginWithToken)

		//------ segment routers------
		router.Get("/segment/all", segmentController.GetAll)
		router.Get("/segment/{id}", segmentController.GetById)
		router.Post("/segment/create", segmentController.CreateSegment)
		router.Delete("/segment/delete/{id}", segmentController.DeleteSegment)
		router.Get("/segment/call/{callId}", segmentController.GetByCallId)
		router.Get("/segment/emo/{id}", segmentController.GetEmotion)

		//------ file routers------
		router.Post("/file/storage/multi/{id}", fileController.UploadMultipleFile)
	})

	// Protected routes
	// Create serve files api

	r.Group(func(protectedRoute chi.Router) {
		// Middleware authentication
		// protectedRoute.Use(jwtauth.Verifier(infrastructure.GetEncodeAuth()))
		// protectedRoute.Use(jwtauth.Authenticator)

		fs := http.StripPrefix("/storage", http.FileServer(http.Dir(infrastructure.GetStoragePath())))
		protectedRoute.Get("/storage/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fs.ServeHTTP(w, r)
		}))
	})
	return r

}
