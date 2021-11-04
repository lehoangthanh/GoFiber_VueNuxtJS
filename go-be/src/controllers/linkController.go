package controllers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"hello/src/database"
	"hello/src/middlewares"
	"hello/src/models"
	"strconv"
)

func Link(c *fiber.Ctx)  error {
 id,_ := strconv.Atoi(c.Params("id"))
 var links []models.Link
	if err := database.DB.Where("user_id = ?", id).Find(&links).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	for i, link := range links {
		var orders []models.Order
		database.DB.Where("code = ? and complete = true", link.Code).Find(&orders)
		links[i].Orders = orders
	}
	return c.JSON(links)
}

type CreateLinkRequest struct {
	Products []int
}

func CreateLink(c *fiber.Ctx) error {
	var request CreateLinkRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	id, _ := middlewares.GetUserId(c)
	link := models.Link{
		UserId: id,
		Code: faker.Username(),
	}
	
	for _, productId := range	request.Products {
		var product models.Product
		product.Id = uint(productId)
		link.Products = append(link.Products, product)
	}
	if err := database.DB.Create(&link).Error; err != nil {
		return err
	}
	
	return c.JSON(link)
}

func Stats(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)
	var links []models.Link
	database.DB.Find(&links, models.Link{
		UserId: id,
	})
	
	var results [] interface{}
	
	var orders []models.Order
	
	revenue := 0.0
	for _, link := range links {
		database.DB.Preload("OrderItems").Find(&orders, &models.Order{
			Code: link.Code,
			Complete: true,
		})
		for _, order := range orders {
			revenue += order.GetTotal()
		}
		results = append(results, fiber.Map{
			"code": link.Code,
			"count": len(orders),
			"revenue": revenue,
		})
	}
	return c.JSON(results)
}

func GetLink(c *fiber.Ctx) error {
	code := c.Params("code")
	var link models.Link

	if err := database.DB.Debug().Preload("User").Preload("Products").First(&link, models.Link{
		Code: code,
	}).Error; err != nil {
		return err
	}
	return c.JSON(link)
}