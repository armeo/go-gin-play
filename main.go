package main

import "github.com/gin-gonic/gin"

type Response struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", pongV1)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", pongV2)
	}

	r.Run(":8080")
}

func pongV1(c *gin.Context) {
	c.String(200, "pong")
}

func pongV2(c *gin.Context) {
	resp := Response{
		Message: "pongpong",
	}

	c.JSON(200, resp)
}
