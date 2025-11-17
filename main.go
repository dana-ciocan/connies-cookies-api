package main

import (
	"log"
	"net/http"

	"example/web-service-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database
    database.InitDB()
    defer database.CloseDB()

    router := gin.Default()
    
    // CORS middleware for demo purposes - allows all origins
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })
    
    router.GET("/cookies", getCookies)
	router.GET("/cookies/:id", getCookieByID)

    log.Println("Starting server on localhost:8181")
    router.Run("localhost:8181")
}

func getCookies(c *gin.Context) {
    cookies, err := database.GetAllCookies()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.IndentedJSON(http.StatusOK, cookies)
}

func getCookieByID(c *gin.Context) {
    id := c.Param("id")

    cookie, err := database.GetCookieByID(id)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if cookie == nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cookie not found"})
        return
    }

    c.IndentedJSON(http.StatusOK, cookie)
}
