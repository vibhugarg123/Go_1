package handlers

import (
	"encoding/json"
	"github.com/vibhugarg123/Go_1/entities"
	"github.com/vibhugarg123/Go_1/service"
	"log"
	"net/http"
)

type AddUserHandler struct {
	Service service.UserService
}

func NewAddUserHandler(userService service.UserService) *AddUserHandler {
	return &AddUserHandler{
		Service: userService,
	}
}

func (amh *AddUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var user entities.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	log.Println("request received", user)
	user, err := amh.Service.Add(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("failed to add the user"))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}
