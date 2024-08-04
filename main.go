package main

import (
	"crud/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/insert", handler.Insert)
	r.GET("/getall", handler.GetAll)
	r.GET("/getParticular/:id", handler.GetParticular)
	r.DELETE("/delete/:id", handler.DeleteParticular)
	r.PUT("/update", handler.Update)
	r.GET("/deleteAll", handler.DeleteAll)
	r.Run()

	// slice := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println("slice", slice)
	// output := append(slice[:2], slice[3:]...)
	// fmt.Println("output", output)
	// slice = nil
	// fmt.Println(slice)

}
