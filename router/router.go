package router

import (
	"Toto/controllers"
	"goji.io"
	"net/http"
	"goji.io/pat"
	
	"github.com/sirupsen/logrus"
	
	
)


// RegisterRoutes creates router and routes requests
func RegisterRoutes(mux *goji.Mux) {

	logrus.WithFields(logrus.Fields{"User": "walrus",}).Info("This is an info message")
	user := new(controllers.UserController)

	mux.HandleFunc(pat.Post("/user"), user.Create)
	mux.HandleFunc(pat.Get("/user/:id"), user.Get)
	mux.HandleFunc(pat.Get("/user"), user.Find)
	mux.HandleFunc(pat.Put("/user/:id"), user.Update)
	mux.HandleFunc(pat.Delete("/user/:id"), user.Delete)

	http.ListenAndServe("localhost:8000", mux)

}
