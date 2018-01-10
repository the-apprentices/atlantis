package main

import(
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/thinkerou/favicon"
)

func main() {
    r := gin.Default()
    r.Use(favicon.New("./img/fav.ico"))
    r.LoadHTMLGlob("templates/*")
    r.GET("/:name", func(c *gin.Context){
        name := c.Param("name");
        c.HTML(http.StatusOK, "test.html", gin.H{
            "title": "Something",
            "content": name,
        })
    })
    r.Run(":8080")
}

