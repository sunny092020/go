package main

import (
    "log"
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

type AddParams struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type PrintJob struct {
    JobId int `json:"jobId" binding:"required,gte=10000"`
    Pages int `json:"pages" binding:"required,gte=1,lte=100"`
}

func addPost(c *gin.Context) {
    var ap AddParams
    if err := c.ShouldBindJSON(&ap); err != nil {
        c.JSON(400, gin.H{"error": "Calculator error"})
        return
    }

    c.JSON(200,  gin.H{"answer": ap.X + ap.Y})
}

func FindUserAgent() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println(c.GetHeader("User-Agent"))
        // Before calling handler
        c.Next()
        // After calling handler
    }
}

func main() {
    router := gin.Default()

    router.Use(FindUserAgent())
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Middleware works!"})
    })

    router.POST("/print", func(c *gin.Context) {
        var p PrintJob
        if err := c.ShouldBindJSON(&p); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input!"})
            return
        }
        c.JSON(200, gin.H{"message": 
            fmt.Sprintf("PrintJob #%v started!", p.JobId)})
    })

    router.GET("/add/:x/:y", add)
    router.POST("/add", addPost)
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
