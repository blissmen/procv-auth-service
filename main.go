
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "auth-service healthy"})
    })
    r.POST("/login", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
    })
    r.Run(":9001")
}
