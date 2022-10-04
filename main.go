package main

import (
	"github.com/ainmtsn1999/go-api-get-weather/controllers/weathercontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//api
	r.GET("/api/weather", weathercontroller.Index)
	r.Run()

}
