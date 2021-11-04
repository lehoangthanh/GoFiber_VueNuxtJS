package controllers

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"hello/src/database"
	"hello/src/models"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Products(c *fiber.Ctx) error{
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(products)
}

func CreteProduct(c *fiber.Ctx) error {
	var product models.Product
	if err:= c.BodyParser(&product); err != nil {
		return err
	}
	
	if err := database.DB.Create(&product).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	go database.ClearCache("product_backend", "product_frontend")
	
	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id, _ := strconv.Atoi(c.Params("id"))
	product.Id = uint(id)
	if err := database.DB.Find(&product).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{}
	product.Id = uint(id)
	if err:= c.BodyParser(&product); err != nil {
		return err
	}
	if err := database.DB.Model(&product).Updates(&product).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	go database.ClearCache("product_backend", "product_frontend")

	return c.JSON(product)
}

func deleteCache(key string) {
time.Sleep(3 * time.Second)
	database.Cache.Del(context.Background(), key)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{}
	product.Id = uint(id)
	if err := database.DB.Model(&product).Delete(&product).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	go database.ClearCache("product_backend", "product_frontend")
	
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func ProductsFrontend(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()
	result, err := database.Cache.Get(ctx, "product_frontend").Result()
	if err != nil {
		if err := database.DB.Find(&products).Error; err != nil {
			panic(err)
		}
		
		BytesProduct, err :=json.Marshal(products)
		if err = database.Cache.Set(ctx, "product_frontend", BytesProduct, 30*time.Minute).Err(); err != nil {
			panic(err)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}
	
	return c.JSON(products)
}

func ProductsBackend(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()
	result, err := database.Cache.Get(ctx, "product_backend").Result()
	if err != nil {
		if err := database.DB.Find(&products).Error; err != nil {
			panic(err)
		}
		
		BytesProduct, err :=json.Marshal(products)
		if err = database.Cache.Set(ctx, "product_backend", BytesProduct, 30*time.Minute).Err(); err != nil {
			panic(err)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}
	
	var searchProducts []models.Product
	
	if s := c.Query("s"); s != "" {
		sLower := strings.ToLower(s)
		for _, product := range products {
			if strings.Contains(product.Title, sLower) || strings.Contains(product.Description, sLower) {
				searchProducts = append(searchProducts, product)
			}
		}
	} else {
		searchProducts = products
	}
	
	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(searchProducts, func(i, j int) bool {
				return searchProducts[i].Price < searchProducts[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(searchProducts, func(i, j int) bool {
				return searchProducts[i].Price > searchProducts[j].Price
			})
		}
	}
	var total = len(searchProducts)
	iPage, _ := strconv.Atoi(c.Query("page", "1"))
	var data []models.Product
	perPage := 10
	
	if total <= iPage*perPage && total >= (iPage - 1) * perPage {
		data = searchProducts[(iPage - 1)*perPage : total]
	} else if total >= iPage*perPage {
		data = searchProducts[(iPage - 1)*perPage : iPage * perPage]
	} else {
		data = []models.Product{}
	}
	
	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      iPage,
		"last_page": (total / perPage) + 1,
	})
}