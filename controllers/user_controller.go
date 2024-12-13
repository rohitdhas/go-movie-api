package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rohitdhas/movie-api-go/configs"
	"github.com/rohitdhas/movie-api-go/models"
	"github.com/rohitdhas/movie-api-go/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	limit := int64(10)
	findOptions := options.Find().SetLimit(limit)
	cursor, err := userCollection.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	defer cursor.Close(ctx)

	var users []models.User

	if err := cursor.All(ctx, &users); err != nil {
		c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error decoding users",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"data": users}})
}
