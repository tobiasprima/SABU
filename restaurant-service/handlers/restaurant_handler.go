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
	RestaurantRepo *repository.RestaurantRepository
}

func NewRestaurantHandler() *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantRepo: repository.NewRestaurantRepository(),
	}
}

// GetRestaurants godoc
// @Summary      Get restaurants
// @Description  Retrieve all restaurant
// @Tags         Restaurant
// @Success      200 {object} []models.Restaurant
// @Failure      500 {object} map[string]string
// @Router       /restaurants [get]
func (h *RestaurantHandler) GetRestaurants(c echo.Context) error {
	restaurants, err := h.RestaurantRepo.GetRestaurants()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch restaurants"})
	}

	if len(restaurants) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "No Restaurants Found",
			"data":  nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": restaurants,
	})
}

// GetRestaurantByID godoc
// @Summary      Get restaurant by ID
// @Description  Retrieve details of a restaurant by its ID
// @Tags         Restaurant
// @Param        restaurant_id path string true "Restaurant ID"
// @Success      200 {object} models.Restaurant
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /restaurant/{restaurant_id} [get]
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
			"data":  nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": restaurant,
	})
}

// AddMeal godoc
// @Summary      Add meal
// @Description  Add meal by restaurant ID
// @Tags         Restaurant
// @Param        restaurant_id path string true "Restaurant ID"
// @Param        body body dtos.AddMealRequest true "Add meal request payload"
// @Success      201 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /restaurant/add-meal/{restaurant_id} [post]
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
		Name:         req.Name,
		Description:  req.Description,
		Price:        req.Price,
		Stock:        req.Stock,
	}
	err := h.RestaurantRepo.AddMeal(ctx, meal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add meal"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Meal added successfully"})
}

// GetMealsByRestaurantID godoc
// @Summary      Get meals by restaurant ID
// @Description  Retrieve details of meals by restaurant ID
// @Tags         Restaurant
// @Param        restaurant_id path string true "Restaurant ID"
// @Success      200 {object} []models.Meal
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /restaurant/get-meals/{restaurant_id} [get]
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
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "No meals found for the specified restaurant ID",
			"data":  nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": meals})
}

// GetMealsByID godoc
// @Summary      Get meals by  ID
// @Description  Retrieve details of meals by its ID
// @Tags         Restaurant
// @Param        meal_id path string true "Meal ID"
// @Success      200 {object} models.Meal
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /restaurant/get-meal/{meal_id} [get]
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
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Meal not found",
			"data":  nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": meal})
}

// UpdateMeal godoc
// @Summary      Update meal
// @Description  Update meal by its ID
// @Tags         Restaurant
// @Param        meal_id path string true "Meal ID"
// @Param        body body map[string]interface{} true "Update meal request payload"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /restaurant/update-meal/{meal_id} [patch]
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
	if description, ok := updates["description"].(string); ok && description != "" {
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

// DeleteMeal godoc
// @Summary      Delete meal
// @Description  Delete meal by its ID
// @Tags         Restaurant
// @Param        meal_id path string true "Meal ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /restaurant/delete-meal/{meal_id} [delete]
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
