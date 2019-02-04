package main

import (
	"github.com/gin-gonic/gin"
)

type Data struct {
	Data  string `json:"Data"`
	Data2 string `json:"Data2"`
}



func geoHash(c *gin.Context) {
	var data Data
	var hash Data
	c.ShouldBind(&data)
	hash = data
	t1 := c.Param("fixed")
	hash.Data2 = t1
	c.Redirect(301, "http://example.com")
	//c.JSON(200, hash)
}

func main() {

	r := gin.Default()
	r.GET("/:fixed/", geoHash)
	r.Run(":8666") // listen and serve on 0.0.0.0:8080
}