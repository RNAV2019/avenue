package server

import (
	"context"
	"fmt"
	"rnav2022/avenue/model"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/api/option"
)

var client *auth.Client

func GetTokenFromJWT(c *fiber.Ctx) (*auth.Token, error) {
	var firebaseID string

	reqHeaders := c.GetReqHeaders()
	reqArray := strings.Split(reqHeaders["Authorization"], " ")

	if len(reqArray) == 2 {
		firebaseID = reqArray[1]
		fmt.Println(firebaseID)
	} else {
		return nil, fmt.Errorf("no authorization headers found")
	}

	// Firebase jwt auth check
	token, err := client.VerifyIDToken(context.Background(), firebaseID)
	fmt.Println(token.UID)
	return token, err

}

func GetUserFromUID(token *auth.Token) (model.User, error) {
	fmt.Println("Users token UID: " + token.UID)
	user, err := model.GetUserByUID(token.UID)
	if err != nil {
		return model.User{}, err
	}
	return *user, err
}

func createUser(c *fiber.Ctx) error {
	token, err := GetTokenFromJWT(c)
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User already exists so no new user created"})
	}

	user := model.User{
		FirebaseID: token.UID,
		Avenue: model.Avenue{
			Description: "New Description",
		},
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

func getCurrentAvenue(c *fiber.Ctx) error {
	token, err := GetTokenFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	user, err := GetUserFromUID(token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	avenue := model.GetCurrentAvenue(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(avenue)
}

func getUserByUID(c *fiber.Ctx) error {
	uid := c.Params("uid")

	user, err := model.GetUserByUID(uid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func getPictureAndName(c *fiber.Ctx) error {
	uid := c.Params("uid")
	profile, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"profile_picture": profile.PhotoURL, "name": profile.DisplayName})
}

func createLink(c *fiber.Ctx) error {
	token, err := GetTokenFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	user, err := GetUserFromUID(token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	var link model.Link
	err = c.BodyParser(&link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	link.AvenueID = user.Avenue.ID
	newLink, err := model.CreateLink(link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(newLink)
}

func SetupFirebase() error {
	opt := option.WithCredentialsFile("firebase/auth/avenue-firebase.json")
	config := &firebase.Config{ProjectID: "avenue-f8597"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return err
	}
	client, err = app.Auth(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	userRoute := router.Group("/user")
	userRoute.Post("/create", createUser)
	userRoute.Get("/info/:uid", getPictureAndName)
	userRoute.Get("/list", getAllUsers)

	avenueRoute := router.Group("/avenue")
	avenueRoute.Get("/current", getCurrentAvenue)
	avenueRoute.Get("/find/:uid", getUserByUID)

	linksRoute := router.Group("/links")
	linksRoute.Post("/create", createLink)
	// linksRoute.Get("/find/:uid", getLinksByUID)

	router.Listen(":3000")
}
