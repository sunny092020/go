package main

import (
    "github.com/gin-gonic/gin"
		"gorm.io/driver/postgres"
		"gorm.io/gorm"
		"log"

		"./models/user"
)


func v1Users(c *gin.Context) {
	
}

func main() {
	dsn := "host=10.6.0.5 user=postgres password=password dbname=oms port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Println(db.First(&OmsUser{}))
	log.Println(err)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/users", v1Users)
	router.Run(":5000")
}
