package handlers

import (
	"net/http"
	"sabu-restaurant-service/dtos"
	"sabu-restaurant-service/models"
	"sabu-restaurant-service/repository"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type RestaurantHandler struct {
	RestaurantRepo		*repository.RestaurantRepository
}

func NewRestaurantHandler() *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantRepo: repository.NewRestaurantRepository(),
	}
}

func (h *RestaurantHandler) GetRestaurantByID(c echo.Context) error {
	restaurantID := c.Param("restaurant_id")
	if restaurantID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Restaurant ID is required"})
	}

	restaurant, err := h.RestaurantRepo.GetRestaurantByID(restaurantID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get restaurant"})
	}

	if restaurant == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Restaurant not found",
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": restaurant,
	})
}

func (h *RestaurantHandler) AddMeal(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(dtos.AddMealRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	restaurantID := c.Param("restaurant_id")
	if restaurantID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Restaurant ID is required"})
	}

	if req.Name == "" || req.Price <= 0 || req.Stock < 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Missing or invalid fields (name, price, stock)",
		})
	}

	meal := &models.Meal{
		RestaurantID: restaurantID,
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Stock: req.Stock,
	}
	err := h.RestaurantRepo.AddMeal(ctx, meal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add meal"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Meal added successfully"})
}

func (h *RestaurantHandler) GetMealsByRestaurantID(c echo.Context) error {
	ctx := c.Request().Context()
	
	restaurantId := c.Param("restaurant_id")
	if restaurantId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Restaurant ID should be provided"})
	}

	meals, err := h.RestaurantRepo.GetMealsByRestaurantID(ctx, restaurantId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get meals"})
	}

	if len(meals) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "No meals found for the specified restaurant ID",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": meals})
}

func (h *RestaurantHandler) GetMealByID(c echo.Context) error {
	ctx := c.Request().Context()

	mealId := c.Param("meal_id")
	if mealId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Meal id is required"})
	}

	meal, err := h.RestaurantRepo.GetMealByID(ctx, mealId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get meal by id"})
	}

	if meal == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Meal not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": meal})
}

func (h *RestaurantHandler) UpdateMeal(c echo.Context) error {
	ctx := c.Request().Context()

	mealId := c.Param("meal_id")
	if mealId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Meal ID is required"})
	}

	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	validUpdates := bson.M{}
	if name, ok := updates["name"].(string); ok && name != "" {
		validUpdates["name"] = name
	}
	if description, ok := updates["description"].(string); ok && description != ""{
		validUpdates["description"] = description
	}
	if price, ok := updates["price"].(float64); ok && price > 0 {
		validUpdates["price"] = price
	}
	if stock, ok := updates["stock"].(int); ok && stock >= 0 {
		validUpdates["stock"] = stock
	}

	if len(validUpdates) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No valid fields to update"})
	}

	err := h.RestaurantRepo.UpdateMeal(ctx, mealId, validUpdates)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update meal"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Meal updated successfully"})
}

func (h *RestaurantHandler) DeleteMeal(c echo.Context) error {
	ctx := c.Request().Context()

	mealId := c.Param("meal_id")
	if mealId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Meal ID is required"})
	}

	err := h.RestaurantRepo.DeleteMeal(ctx, mealId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete meal"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Meal deleted successfully"})
}