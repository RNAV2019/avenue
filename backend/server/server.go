package server

import (
	"rnav2022/avenue/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Token struct {
	Token string `json:"token"`
}

func createUser(c *fiber.Ctx) error {
	var token Token

	err := c.BodyParser(&token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid body")
	}

	user := &model.User{
		FBToken: token.Token,
	}
	newUser, err := model.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(newUser)
}

func getAllUsers(c *fiber.Ctx) error {
	users, err := model.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	userRoute := router.Group("/user")
	userRoute.Post("/create", createUser)
	userRoute.Get("/list", getAllUsers)

	// avenueRoute := router.Group("/avenue")
	// avenueRoute.Get("/current", getAvenue)

	router.Listen(":3000")
}
