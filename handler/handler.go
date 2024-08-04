package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct { //pasword field canbe hidden using json tag
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Address  string `json:"address"`
	Password string `json:"-"` // Assumed addition to hide passwords in JSON responses
}
// type Person1 struct {
// 	FirstName string
// 	LastName  string
// 	Id        int
// }

// both struct working in json data
var personSlice []Person //************[]only then you can iterate

func Insert(c *gin.Context) {
	var person []Person
	err := c.BindJSON(&person) //(bindJson:pasre the data jSON to struct)passing data to context c
	if err != nil {
		fmt.Println("error occured in binding", err)
		c.JSON(404, err)
		return
	}
	fmt.Println("person:", person)
	personSlice = append(personSlice, person...)
	fmt.Println("personSlice", personSlice)
	c.JSON(200, "yeah! Inserted")
}
func GetAll(c *gin.Context) {

	fmt.Println(personSlice)
	c.JSON(200, personSlice)
}
func GetParticular(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("error converting id to integer:", err)
		c.JSON(400, "Invalid ID")
		return
	}
	for _, person := range personSlice {
		if person.Id == id {
			c.JSON(http.StatusOK, person)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
}

func DeleteAll(c *gin.Context) {
	//personSlice = nil //
	personSlice = append(personSlice[:0], personSlice[len(personSlice):]...)
	fmt.Println("personSlice", personSlice)
	c.JSON(200, "All data deleted")
}
func DeleteParticular(c *gin.Context) {
	delId := c.Param("id")
	dd, err := strconv.Atoi(delId)
	if err != nil {
		fmt.Println(err)
	}
	var persondeleted Person
	for index, valueofperson := range personSlice {
		if valueofperson.Id == dd {
			persondeleted = valueofperson //saving a person deleted into variable
			personSlice = append(personSlice[:index], personSlice[index+1:]...)
			c.JSON(200, persondeleted)
			return
		}
	}
	fmt.Println("personSlice::", personSlice)
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})

}
func Update(c *gin.Context) {
	var person Person
	if err := c.BindJSON(&person); err != nil {
		fmt.Println("err in binding", err)
	}
	for i := range personSlice {
		if personSlice[i].Id == person.Id {
			personSlice[i].Address = person.Address
			c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully"})
			c.JSON(200, personSlice)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

//json pkg bind(),
//ioutil pkg(io.ReadAll)
