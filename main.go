package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/0xatanda/goFeedNews/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT isn't found in the env")
	}

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerV1 := router.Group("/v1")
	routerV1.GET("/health", handlers.HandlerReadiness)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is starting on port: ", portString)

}
