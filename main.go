package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT isn't found in the env")
	}

	app := fiber.New()

	if err := app.Listen(":" + portString); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is starting on port: ", portString)

}
