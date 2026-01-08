package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/repository"
	"github.com/iamskyy666/golang-job-portal-api/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err:=godotenv.Load();err!=nil{
		log.Fatal("⚠️ERROR loading .env file:",err.Error())
	}
	db,err:=repository.InitDB()
	if err!=nil{
		log.Fatal("ERROR:",err)
	}
	defer db.Close()
	r:=gin.Default()
	routes.InitRoutes(r,db)

	port:=os.Getenv("SERVER_PORT")
	if port==""{
		port="8080"
	}
	r.Run(":"+port)
	
}

//00:30:07