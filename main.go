package main

import "github.com/gin-gonic/gin"

func main() {
	var a int
	r := gin.Default()
	r.Run(":8080")
	print(a)
}
