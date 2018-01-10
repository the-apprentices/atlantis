package main

import(
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/thinkerou/favicon"
)

func main() {
    r := gin.Default()
    r.Use(favicon.New("./img/fav.ico"))
    r.LoadHTMLGlob("templates/*")
    r.POST("/postf", func(c *gin.Context){
        value := c.PostForm("exval")
        c.JSON(200, gin.H{
            "result": "OK",
            "value": value,
        })
        fmt.Printf("%s", value);
    })
    r.GET("/exform", func(c *gin.Context){
        c.HTML(http.StatusOK, "exform.html", gin.H{
            "title": "Testing form",
        })
    })
    r.GET("/noform/:name", func(c *gin.Context){
        name := c.Param("name");
        c.HTML(http.StatusOK, "test.html", gin.H{
            "title": "Something",
            "content": name,
        })
    })
    r.Run(":8080")
}

