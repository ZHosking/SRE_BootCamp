package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ZHosking/SREBootcamp/web-service-gin/models"
	"github.com/ZHosking/SREBootcamp/web-service-gin/utils"
	"github.com/gin-gonic/gin"
)

func GetStudentsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		students, err := models.GetAllStudents(db)
		if err != nil {
			utils.ErrorLogger.Printf("Failed to fetch students: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if students == nil {
			students = []models.Student{}
		}

		utils.InfoLogger.Printf("Fetched %d students", len(students))
		c.JSON(http.StatusOK, students)
	}
}

func GetStudentByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		student, err := models.GetStudentByID(db, id)
		if err != nil {
			utils.Error(err, "Student not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if student == nil {
			utils.InfoLogger.Printf("Student not found")
			c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
			return
		}
		utils.Info("Successfully found student by ID")
		c.JSON(http.StatusOK, student)
	}
}

func AddStudentHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var s models.Student
		if err := c.BindJSON(&s); err != nil {
			utils.ErrorLogger.Printf("Invalid JSON %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		if err := models.AddStudent(db, s); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, s)
	}
}

func UpdateStudentHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var s models.Student
		if err := c.BindJSON(&s); err != nil {
			utils.ErrorLogger.Printf("Invalid JSON %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		s.ID = id
		if err := models.UpdateStudent(db, s); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, s)
	}
}

func DeleteStudentHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := models.DeleteStudent(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
	}
}
