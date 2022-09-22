package main

import (
	"first-gin/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

//这里自动引入routers的包

func main() {
	fmt.Println("test")
	r := gin.Default()
	routers.IndexInit(r) //调用方法
	r.Run()
}
