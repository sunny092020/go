package main

import (
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
)

func v1EndpointHandler(c *gin.Context) {
    c.String(200, "v1: %s %s", c.Request.Method, c.Request.URL.Path)
}

func v2EndpointHandler(c *gin.Context) {
    c.String(200, "v2: %s %s", c.Request.Method, c.Request.URL.Path)
}

func add(c *gin.Context) {
    x, _ := strconv.ParseFloat(c.Param("x"), 64)
    y, _ := strconv.ParseFloat(c.Param("y"), 64)
    c.String(200,  fmt.Sprintf("%f", x + y))
}

func main() {
    router := gin.Default()
    router.GET("/add/:x/:y", add)

    v1 := router.Group("/v1")

    v1.GET("/products", v1EndpointHandler)
    // Eg: /v1/products
    v1.GET("/products/:productId", v1EndpointHandler)
    v1.POST("/products", v1EndpointHandler)
    v1.PUT("/products/:productId", v1EndpointHandler) 
    v1.DELETE("/products/:productId", v1EndpointHandler)

    v2 := router.Group("/v2")

    v2.GET("/products", v2EndpointHandler)
    v2.GET("/products/:productId", v2EndpointHandler)
    v2.POST("/products", v2EndpointHandler)
    v2.PUT("/products/:productId", v2EndpointHandler)
    v2.DELETE("/products/:productId", v2EndpointHandler)

    router.Run(":5000")
}
