package controllers

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"hello/src/database"
	"hello/src/models"
)

func Ambassador(c *fiber.Ctx) error {
	var users []models.User
	if err := database.DB.Where("is_ambassador = true").Find(&users).Error; err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	return c.JSON(users)
}

func Rankings(c *fiber.Ctx) error {
	rankings, err := database.Cache.ZRevRangeByScoreWithScores(context.Background(), "rankings", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	
	if err != nil {
		return err
	}
	
	result := make(map[string]float64)
	for _, ranking := range rankings{
		result[ranking.Member.(string)] = ranking.Score
	}
	
	return c.JSON(result)
}