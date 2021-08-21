package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Product struct {
	Name    string
	Type    string
	Price   int64
	Store   string
	Barcode string
}

var products []Product = []Product{
	{
		Name:    "Indomie",
		Type:    "Food",
		Price:   3000,
		Store:   "Cabang A",
		Barcode: "112233",
	},
	{
		Name:    "Clear Shampoo",
		Type:    "Toiletries",
		Price:   15000,
		Store:   "Cabang A",
		Barcode: "223344",
	},
	{
		Name:    "Iphone 18",
		Type:    "Electronic",
		Price:   18000000,
		Store:   "Cabang A",
		Barcode: "334455",
	},
}

var rahasia = "rahasia-cuy"

func main() {
	app := fiber.New()

	app.Post("/api/login", login)

	app.Use("/api/transactions", jwtware.New(jwtware.Config{
		SigningKey: []byte(rahasia),
	}))

	app.Get("/api/transactions", func(c *fiber.Ctx) error {
		return c.JSON(products)
	})

	app.Listen(":8000")
}

func login(c *fiber.Ctx) error {
	var user User
	c.BodyParser(&user)

	// Throws Unauthorized error
	if user.User != "doni" || user.Password != "rahasia" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "doni"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(rahasia))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}