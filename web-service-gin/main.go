package main

import (
	"github.com/ZHosking/SREBootcamp/web-service-gin/handlers"
	"github.com/ZHosking/SREBootcamp/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db, err := models.InitDB("students.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// routes
	router.GET("/students", handlers.GetStudentsHandler(db))
	router.GET("/students/:id", handlers.GetStudentByIDHandler(db))
	router.POST("/students", handlers.AddStudentHandler(db))
	router.PATCH("/students/:id", handlers.UpdateStudentHandler(db))
	router.DELETE("/students/:id", handlers.DeleteStudentHandler(db))

	//healthcheck call
	router.GET("/healthcheck", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(503, gin.H{"status": "unhealthy"})
			return
		}
		c.JSON(200, gin.H{"status": "healthy"})
	})

	router.Run(":8080")

}
