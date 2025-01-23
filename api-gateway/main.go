package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sabu-api-gateway/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

var serviceMap map[string]string

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("failed to load .env file: %v", err)
	// }

	serviceMap = map[string]string{
		"foundation": os.Getenv("FOUNDATION_SERVICE"),
		"donor": os.Getenv("DONOR_SERVICE"),
		"restaurant": os.Getenv("RESTAURANT_SERVICE"),
		"user": os.Getenv("USER_SERVICE"),
	}
}

func forwardRequest(serviceURL string, parts []string, c echo.Context) error {
	if serviceURL == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Service not found"})
	}

	targetURL := serviceURL + "/" + parts[0] + "/" + strings.Join(parts[1:], "/")

	fmt.Println("targeturl", targetURL)

	if parts[0] == "foundation" && len(parts) > 1 && parts[1] == "complete-order"{
		email, _ := c.Get("email").(string)

		payload := map[string]interface{}{
			"email": email,
		}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create payload"})
		}

		req, err := http.NewRequest(c.Request().Method, targetURL, bytes.NewBuffer(payloadBytes))
		if err != nil {
			return err
		}

		req.Header = c.Request().Header
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
	}

	req, err := http.NewRequest(c.Request().Method, targetURL, c.Request().Body)
	if err != nil {
		return err
	}
	req.Header = c.Request().Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
}

func main() {
	e := echo.New()

	e.Any("/*", func(c echo.Context) error {
		path := c.Request().URL.Path
		parts := strings.Split(strings.Trim(path, "/"), "/")
		if len(parts) < 1 {
			return c.JSON(http.StatusBadRequest, map[string] string{"message": "Invalid path"})
		}

		serviceName := parts[0]

		if serviceName == "user" {
			return forwardRequest(serviceMap[serviceName], parts, c)
		}

		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing Token"})
		}

		tokenStr := strings.TrimPrefix(token, "Bearer ")

		claims, err := utils.ParseJWT(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}

		if email, ok := claims["email"]; ok {
			c.Set("email", email)
		}

		return forwardRequest(serviceMap[serviceName], parts, c)
	})


	e.Logger.Fatal(e.Start(":8085"))
}