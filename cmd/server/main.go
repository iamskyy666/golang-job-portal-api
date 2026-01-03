package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/repository"
)

func main() {
	db,err:=repository.InitDB()
	if err!=nil{
		log.Fatal("ERROR:",err)
	}
	defer db.Close()
	
	r := gin.Default()
	r.Run(":8080")
}