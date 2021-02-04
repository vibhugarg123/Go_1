package server

import (
	"github.com/gorilla/mux"
	"github.com/vibhugarg123/Go_1/handlers"
	"github.com/vibhugarg123/Go_1/service"
	"net/http"
)

func Router() http.Handler {
	router := mux.NewRouter()
	//Route Handlers
	userService := service.NewUserService()
	router.Handle("/api/adduser", handlers.NewAddUserHandler(userService)).Methods("POST")
	return router
}
