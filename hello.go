package main

import(
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")
    r.GET("/index", func(c *gin.Context){
        c.HTML(http.StatusOK, "test.html", gin.H{
            "title": "Something",
        })
    })
    r.Run(":8080")
}

