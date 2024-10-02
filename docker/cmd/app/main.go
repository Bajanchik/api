package main

import (
	"docker/internal/database"
	"docker/internal/handlers"
	"docker/internal/messagesService"
	"docker/internal/userService"
	"docker/internal/web/messages"
	"docker/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&messagesService.Message{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	} else if err := database.DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	messagesRepo := messagesService.NewMessageRepository(database.DB)
	messagesService := messagesService.NewService(messagesRepo)
	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewService(usersRepo)

	messagesHandler := handlers.NewMessageHandler(messagesService)
	usersHandler := handlers.NewUserHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictMessageHandler := messages.NewStrictHandler(messagesHandler, nil) // тут будет ошибка
	messages.RegisterHandlers(e, strictMessageHandler)
	strictUserHandler := users.NewStrictHandler(usersHandler, nil) // тут будет ошибка
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
