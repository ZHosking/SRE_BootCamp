package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct for data for each student
type student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

// students slice to seed student data
var students = []student{

	{ID: "1", Name: "John Smith", Age: "21"},
	{ID: "2", Name: "Nathan Todd", Age: "19"},
	{ID: "3", Name: "Adam James", Age: "32"},
}

func main() {

	router := gin.Default()
	router.GET("/students", getStudents)
	router.POST("/students", postStudents)

	router.Run("localhost:8080")

}

// getStudents responds with list of students as JSON
func getStudents(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, students)

}

// postStudents adds a studen from JSON received in the request
func postStudents(c *gin.Context) {

	var newStudent student

	// Call BindJSON to bind the received JSON to newStudent Var
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)

}
