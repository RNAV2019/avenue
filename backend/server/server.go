package server

import (
	"context"
	"fmt"
	"os"
	"rnav2022/avenue/model"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	supa "github.com/nedpals/supabase-go"
)

var supabase *supa.Client

type Description struct {
	Description string `json:"description"`
}

func getAvenueFromJWT(tokenString string) (model.Avenue, error) {
	user, err := supabase.Auth.User(context.TODO(), tokenString)
	if err != nil {
		return model.Avenue{}, err
	}
	id := user.ID
	avenue, err := model.GetAvenueByUID(id)
	if err != nil {
		return model.Avenue{}, err
	}
	return avenue, nil
}

func createAvenue(c *fiber.Ctx) error {
	var tokenString string
	reqHeaders := c.GetReqHeaders()
	reqArray := strings.Split(reqHeaders["Authorization"], " ")

	if len(reqArray) == 2 {
		tokenString = reqArray[1]
	} else {
		return fmt.Errorf("no authorization headers found")
	}

	user, err := supabase.Auth.User(context.TODO(), tokenString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Error with authorization request",
		})
	}
	id := user.ID

	avenue, err := model.CreateAvenue(id, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(avenue)
}

func updateAvenue(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"Error":  "Error with getting the avenue",
				"Avenue": avenue,
				"err":    err,
			},
		)
	}

	var description Description
	err = c.BodyParser(&description)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"Error":  "Error with parsing the description",
				"Avenue": avenue,
				"err":    err,
			},
		)
	}
	avenue.Description = description.Description
	newAvenue, err := model.UpdateAvenue(avenue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(newAvenue)
}

func getAvenueByUID(c *fiber.Ctx) error {
	uid := c.Params("uid")

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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"profile_picture": profileURL, "name": displayName, "statistic": statistic, "avenue": avenue})
}

func createLink(c *fiber.Ctx) error {
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
	}

	if tokenString == "" {
		return fmt.Errorf("no authorization headers found")
	}
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
	}

	if tokenString == "" {
		return fmt.Errorf("no authorization headers found")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	err = model.DeleteLink(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Link deleted"})
}

func getAllStatistics(c *fiber.Ctx) error {
	uid := c.Params("uid")
	avenue, err := model.GetAvenueByUID(uid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	statistics, err := model.GetAllStatistics(int(avenue.ID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	clicks, err := model.GetAggregateClicks(int(avenue.ID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"statistics": statistics, "clicks": clicks, "avenue": avenue})
}

func SetupAndListen() {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase = supa.CreateClient(supabaseUrl, supabaseKey)

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Geo-Location, Referer",
		AllowMethods: "GET, POST, PATCH, DELETE",
	}))

	avenueRoute := router.Group("/avenue")
	avenueRoute.Post("/create", createAvenue)
	avenueRoute.Patch("/update", updateAvenue)
	avenueRoute.Get("/find/:uid", getAvenueByUID)
	avenueRoute.Get("/userinfo/:uid", getPictureAndName)

	linksRoute := router.Group("/links")
	linksRoute.Post("/create", createLink)
	linksRoute.Patch("/update", updateLink)
	linksRoute.Delete("/delete/:id", deleteLink)

	dashboardRoute := router.Group("/dashboard")
	dashboardRoute.Get("/all/:uid", getAllStatistics)

	router.Listen("0.0.0.0:3030")
}
