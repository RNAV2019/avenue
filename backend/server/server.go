package server

import (
	"context"
	"fmt"
	"rnav2022/avenue/model"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	supa "github.com/nedpals/supabase-go"
)

var supabase *supa.Client

// func parseToken(tokenString string) (*jwt.Token, error) {
// 	signingKey := "super-secret-jwt-token-with-at-least-32-characters-long"
// 	fmt.Println("Token String" + tokenString)

// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		return signingKey, nil
// 	})

// 	// Check for errors
// 	if err != nil {
// 		fmt.Println("Error parsing token:", err)
// 		return nil, err
// 	}
// 	return token, nil
// }

func getAvenueFromJWT(tokenString string) (model.Avenue, error) {
	// token, err := parseToken(tokenString)
	// if err != nil {
	// 	fmt.Println("Error parsing token:", err)
	// }
	fmt.Println(tokenString)
	user, err := supabase.Auth.User(context.TODO(), tokenString)
	if err != nil {
		fmt.Println("Error parsing token:", err)
	}
	id := user.ID

	// id := token.Claims.(jwt.MapClaims)["id"].(string)
	fmt.Println(id)
	avenue, err := model.GetAvenueByUID(id)
	if err != nil {
		return model.Avenue{}, err
	}
	return avenue, nil
}

func createAvenue(c *fiber.Ctx) error {
	// Get token string from headers
	var tokenString string
	reqHeaders := c.GetReqHeaders()
	reqArray := strings.Split(reqHeaders["Authorization"], " ")

	if len(reqArray) == 2 {
		tokenString = reqArray[1]
	} else {
		return fmt.Errorf("no authorization headers found")
	}
	fmt.Println("The token string" + tokenString)

	// Get the id as a string
	user, err := supabase.Auth.User(context.TODO(), tokenString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Error with authorization request",
		})
	}
	id := user.ID
	fmt.Println(id)

	// Create the avenue and return it
	avenue, err := model.CreateAvenue(id, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(avenue)
}

func getCurrentAvenue(c *fiber.Ctx) error {
	var tokenString string

	reqHeaders := c.GetReqHeaders()
	reqArray := strings.Split(reqHeaders["Authorization"], " ")

	if len(reqArray) == 2 {
		tokenString = reqArray[1]
	} else {
		return fmt.Errorf("no authorization headers found")
	}

	avenue, err := getAvenueFromJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(avenue)
}

func getAvenueByUID(c *fiber.Ctx) error {
	uid := c.Params("uid")
	fmt.Println("This is the users uuid" + uid)

	avenue, err := model.GetAvenueByUID(uid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(avenue)
}

func getPictureAndName(c *fiber.Ctx) error {
	uid := c.Params("uid")
	profileURL, displayName, err := model.GetProfileImageAndName(uid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	avenue, err := model.GetAvenueByUID(uid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	geographicLocation := c.Get("X-Geo-Location")
	trafficSource := c.Get("Referer")
	statisticModel := model.Statistic{
		AvenueID:           avenue.ID,
		ClickTimestamp:     time.Now(),
		GeographicLocation: geographicLocation,
		TrafficSource:      trafficSource,
	}
	statistic, err := model.CreateStatistic(statisticModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"profile_picture": profileURL, "name": displayName, "statistic": statistic})
}

func createLink(c *fiber.Ctx) error {
	// Get token string from headers
	var tokenString string
	reqHeaders := c.GetReqHeaders()
	reqArray := strings.Split(reqHeaders["Authorization"], " ")

	if len(reqArray) == 2 {
		tokenString = reqArray[1]
	} else {
		return fmt.Errorf("no authorization headers found")
	}
	fmt.Println("The token string" + tokenString)

	avenue, err := getAvenueFromJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var link model.Link
	err = c.BodyParser(&link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	link.AvenueID = avenue.ID
	newLink, err := model.CreateLink(link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(newLink)
}

func updateLink(c *fiber.Ctx) error {
	var tokenString string
	reqHeaders := c.GetReqHeaders()
	reqArray := strings.Split(reqHeaders["Authorization"], " ")

	if len(reqArray) == 2 {
		tokenString = reqArray[1]
	} else {
		return fmt.Errorf("no authorization headers found")
	}
	fmt.Println("The token string" + tokenString)

	var link model.Link
	err := c.BodyParser(&link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	updatedLink, err := model.UpdateLink(link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(updatedLink)
}

func deleteLink(c *fiber.Ctx) error {
	var tokenString string
	reqHeaders := c.GetReqHeaders()
	reqArray := strings.Split(reqHeaders["Authorization"], " ")

	if len(reqArray) == 2 {
		tokenString = reqArray[1]
	} else {
		return fmt.Errorf("no authorization headers found")
	}
	fmt.Println("The token string" + tokenString)
	id, err := strconv.Atoi(c.Params("id"))
	fmt.Println(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	err = model.DeleteLink(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Link deleted"})
}

// func getAggregateClicks(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(err)
// 	}
// 	clicks, err := model.GetAggregateClicks(id)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(err)
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"clicks": clicks})
// }

func getAllStatistics(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	statistics, err := model.GetAllStatistics(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	clicks, err := model.GetAggregateClicks(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"statistics": statistics, "clicks": clicks})
}

// func GetTokenFromJWT(c *fiber.Ctx) (*auth.Token, error) {
// 	var firebaseID string

// 	reqHeaders := c.GetReqHeaders()
// 	reqArray := strings.Split(reqHeaders["Authorization"], " ")

// 	if len(reqArray) == 2 {
// 		firebaseID = reqArray[1]
// 		fmt.Println(firebaseID)
// 	} else {
// 		return nil, fmt.Errorf("no authorization headers found")
// 	}

// 	// Firebase jwt auth check
// 	token, err := client.VerifyIDToken(context.Background(), firebaseID)
// 	fmt.Println(token.UID)
// 	return token, err

// }

// func GetUserFromUID(token *auth.Token) (model.User, error) {
// 	fmt.Println("Users token UID: " + token.UID)
// 	user, err := model.GetUserByUID(token.UID)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return *user, err
// }

// func createUser(c *fiber.Ctx) error {
// 	// token, err := GetTokenFromJWT(c)
// 	// if err == nil {
// 	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User already exists so no new user created"})
// 	// }

// 	user := model.User{
// 		Avenue: model.Avenue{
// 			Description: "New Description",
// 		},
// 	}
// 	newUser, err := model.CreateUser(user)

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(err)
// 	}
// 	return c.Status(fiber.StatusOK).JSON(newUser)
// }

// func getAllUsers(c *fiber.Ctx) error {
// 	users, err := model.GetAllUsers()
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(err)
// 	}
// 	return c.Status(fiber.StatusOK).JSON(users)
// }

// func SetupFirebase() error {
// 	opt := option.WithCredentialsFile("firebase/auth/avenue-firebase.json")
// 	config := &firebase.Config{ProjectID: "avenue-f8597"}
// 	app, err := firebase.NewApp(context.Background(), config, opt)
// 	if err != nil {
// 		return err
// 	}
// 	client, err = app.Auth(context.TODO())
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func SetupAndListen() {
	supabaseUrl := "http://localhost:54321"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6ImFub24iLCJleHAiOjE5ODM4MTI5OTZ9.CRXP1A7WOeoJeXxjNni43kdQwgnWNReilDMblYTn_I0"
	supabase = supa.CreateClient(supabaseUrl, supabaseKey)

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Geo-Location, Referer",
	}))

	// userRoute := router.Group("/user")
	// userRoute.Post("/create", createUser)
	// userRoute.Get("/list", getAllUsers)

	avenueRoute := router.Group("/avenue")
	avenueRoute.Get("/current", getCurrentAvenue)
	avenueRoute.Post("/create", createAvenue)
	avenueRoute.Get("/find/:uid", getAvenueByUID)
	avenueRoute.Get("/userinfo/:uid", getPictureAndName)

	linksRoute := router.Group("/links")
	linksRoute.Post("/create", createLink)
	linksRoute.Patch("/update", updateLink)
	linksRoute.Delete("/delete/:id", deleteLink)

	dashboardRoute := router.Group("/dashboard")
	// dashboardRoute.Get("/count/:id", getAggregateClicks)
	dashboardRoute.Get("/all/:id", getAllStatistics)

	router.Listen(":3000")
}
