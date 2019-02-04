package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mmcloughlin/geohash"
)

type Data struct {
	Lat  float64 `json:"Lat"`
	Long float64 `json:"Long"`
}

type GeoHashData struct {
	GeoHash string `json:"GeoHash"`
}

func geoHash(c *gin.Context) {
	var data Data
	var hash GeoHashData
	c.ShouldBind(&data)
	hash.GeoHash = geohash.Encode(data.Lat, data.Long)
	c.JSON(200, hash)
}

func main() {

	r := gin.Default()
	r.POST("/geoHash", geoHash)
	r.Run() // listen and serve on 0.0.0.0:8080
}