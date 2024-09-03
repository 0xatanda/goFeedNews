package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/0xatanda/goFeedNews/config"
	"github.com/0xatanda/goFeedNews/handlers"
	db "github.com/0xatanda/goFeedNews/sql/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT isn't found in the env")
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("db_url is  not found in the env")
	}

	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("can't conn to db", err)
	}

	apicfg := config.APIConfig{
		DB: db.New(conn),
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
	routerV1.GET("/err", handlers.HandlerError)
	routerV1.POST("/createUser", handlers.HandlerCreateUser(&apicfg))

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is starting on port: ", portString)

}
