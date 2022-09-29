package routes

import (
	"go-rrr/pkg/controllers"
	"time"

	"github.com/go-chi/httprate"
	"github.com/gorilla/mux"
)

var RegisterRedirectRoutes = func(router *mux.Router) {
	router.HandleFunc("/redirect", controllers.CreateRedirect).Methods("POST")
	router.HandleFunc("/redirect/{ShortUrl}", controllers.GetRedirectById).Methods("GET")
	// Too Many Requests
	router.Use(httprate.LimitByIP(20, 1*time.Minute))
}
