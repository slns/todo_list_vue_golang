package main

import (
	"fmt"

    "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/slns/todoListvueGo/Go_Server/database"
	"github.com/slns/todoListvueGo/Go_Server/models"
)

func main() {
	app := fiber.New()

    // Default config
    app.Use(cors.New())
    
	initDatabase()
	setupRoutes(app)

	app.Listen(":8080")
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("database/DB_Sqlite.sqlite"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Database connected!")
	// Migrate the schema
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
	
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.GetTodoById)
	app.Post("/todos", models.CreateTodo)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}
