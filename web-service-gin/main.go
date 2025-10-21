package main

import (
	"os"

	"log"

	"github.com/ZHosking/SREBootcamp/web-service-gin/handlers"
	"github.com/ZHosking/SREBootcamp/web-service-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env variables")
	}

	router := gin.Default()

	dbPath := os.Getenv("STUDENT_DB")

	if dbPath == "" {
		dbPath = "students.db"
		log.Printf("No STUDENT_DB set, falling back to default: %s", dbPath)
	}

	db, err := models.InitDB(dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// routes
	api := router.Group("/api/v1")
	{
		api.GET("/students", handlers.GetStudentsHandler(db))
		api.GET("/students/:id", handlers.GetStudentByIDHandler(db))
		api.POST("/students", handlers.AddStudentHandler(db))
		api.PATCH("/students/:id", handlers.UpdateStudentHandler(db))
		api.DELETE("/students/:id", handlers.DeleteStudentHandler(db))
		api.GET("/healthcheck", func(c *gin.Context) {
			if err := db.Ping(); err != nil {
				c.JSON(503, gin.H{"status": "unhealthy"})
				return
			}
			c.JSON(200, gin.H{"status": "healthy"})
		})
	}

	router.Run(":8080")

}
