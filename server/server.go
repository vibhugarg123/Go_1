package server

import (
	"github.com/vibhugarg123/Go_1/config"
	"log"
	"net/http"
)

func Start() {
	router := Router()
	log.Println(config.GetMySqlConnection())

	log.Fatal(http.ListenAndServe(":8089", router))
	log.Println("successfully started book my show server")
}
