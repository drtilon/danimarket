package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite for example
    "log"
    "danimarket/models"
)

func main() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    db.AutoMigrate(&models.User{})

    r := gin.Default()
    r.GET("/users", func(c *gin.Context) {
        var users []models.User
        db.Find(&users)
        c.JSON(200, users)
    })

    r.Run(":8080")
}

