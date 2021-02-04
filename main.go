package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/vibhugarg123/Go_1/server"
)

func main() {
	server.Start()
}
