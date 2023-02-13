package main

import (
	"fmt"
	"log"
	"membership/config"
	"membership/injector"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load .env %v: ", err)
	}

	conns := config.NewConn()
	product := conns.ProductConn()
	order := conns.OrderConn()
	customer := conns.CustomerConn()
	user := conns.UserConn()

	port := os.Getenv("PORT")

	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./public", fiber.Static{Compress: true, Browse: true})

	api := app.Group("api")

	injector.NewProductInjector(product, api)
	injector.NewOrderInjector(order, api)
	injector.NewCustomerInjector(customer, api)
	injector.NewUserInjector(user, api)

	fmt.Printf("⚡️[server]: Server is running on porting %s\n", port)

	log.Fatal(app.Listen(port))

}
