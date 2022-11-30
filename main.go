package main

import (
	"github.com/gofiber/fiber/v2"
	"healthRoutine/pkgs/database"
)

const addr = ":3000"

func main() {
	app := fiber.New()
	db := database.Conn()

	//userRepo := repository.NewUserRepository(db)
	//
	//useCases :=

	app.Listen(addr)

}
