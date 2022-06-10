package main

import (
	"fmt"
	"klikdokter/authorization"
	"klikdokter/handler"
	"klikdokter/helper"
	"klikdokter/products"
	"klikdokter/users"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/klikdokter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection Database is Successful")

	userRepository := users.NewRepository(db)
	productRepository := products.NewRepository(db)

	userServices := users.NewServices(userRepository)
	productServices := products.NewServices(productRepository)

	authServices := authorization.NewServices()

	userHandler := handler.NewUserHandler(userServices, authServices)
	productHandler := handler.NewProductHandler(productServices, authServices)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	api := router.Group("/api/v1")

	// API FOR USERS Authentication

	api.POST("/auth/login", userHandler.Login)
	api.POST("/register", userHandler.RegisterUser)
	api.GET("/items", productHandler.GetAllProduct)

	// API FOR PRODUCTS

	api.Use(authMiddleware(authServices, userServices))
	api.POST("/item/add", productHandler.SaveProduct)
	api.POST("/item/update", productHandler.UpdateDataBySKU)
	api.POST("/item/remove", productHandler.RemoveDataBySku)
	api.GET("/item/search", productHandler.GetBySku)

	router.Run(":3500")

}

func authMiddleware(authServices authorization.Services, userServices users.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized Access", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authServices.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized Access", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized Access", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))
		user, err := userServices.GetUsersByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized Access", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)

	}
}
