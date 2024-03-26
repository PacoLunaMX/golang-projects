package main

import (
	"fmt"

	"github.com/francisco/go-fiber-crm/database"
	"github.com/francisco/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead", lead.DeleteLead)

}

func initDatabase() {
	var err error
	database, err := gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
