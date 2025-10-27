package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ZHosking/SREBootcamp/web-service-gin/models"
	"github.com/ZHosking/SREBootcamp/web-service-gin/utils"
	"github.com/gin-gonic/gin"
)

// helper to create a test router with the handler
func setupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/students", GetStudentsHandler(db))
	r.GET("/students/:id", GetStudentByIDHandler(db))
	r.POST("/students", AddStudentHandler(db))
	r.PATCH("/students/:id", UpdateStudentHandler(db))
	r.DELETE("/students/:id", DeleteStudentHandler(db))
	return r

}

func setupTestDB() *sql.DB {

	db, err := models.ConnectDB("file::memory:?cache=shared")
	if err != nil {
		log.Fatalf("Failed to init test DB: %v", err)
	}

	models.Migrate(db)

	student := models.Student{
		Name: "Zak Hosking",
		Age:  26,
	}

	err = models.AddStudent(db, student)
	if err != nil {
		log.Fatalf("Failed to seed test DB: %v", err)
	}
	return db

}

func TestGetStudentsHandler(t *testing.T) {

	utils.Init()
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

	req, _ := http.NewRequest("GET", "/students", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expected := `[{"id":1,"name":"Zak Hosking","age":26}]`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, w.Body.String())
	}
}

func TestGetStudentByIDHandler(t *testing.T) {

	utils.Init()
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

	//test to check if student exists - should return 200
	req, _ := http.NewRequest("GET", "/students/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	//test to check if student doesn't exists
	req2, _ := http.NewRequest("GET", "/students/999", nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for non-existent student, got %d", w2.Code)
	}

	t.Logf("Response body: %s", w.Body.String())
	t.Logf("Non-existent student response: %s", w2.Body.String())

}

func TestAddStudentHandler(t *testing.T) {

	utils.Init()
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

	newStudentJSON := `{"name":"Josh Hosking","age":22}`
	req, _ := http.NewRequest("POST", "/students", strings.NewReader(newStudentJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	expected := `{"id":0,"name":"Josh Hosking","age":22}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, w.Body.String())
	}
}

func TestDeleteStudentHandler(t *testing.T) {

	utils.Init()
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

	req, _ := http.NewRequest("DELETE", "/students/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expected := `{"message":"Student deleted"}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, w.Body.String())
	}

	// confirm student is gone
	req2, _ := http.NewRequest("GET", "/students/1", nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 after delete, got %d", w2.Code)
	}

}

func TestUpdateStudentHandler(t *testing.T) {

	utils.Init()
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

	updateJSON := `{"name":"Zak Updated","age":30}`
	req, _ := http.NewRequest("PATCH", "/students/1", strings.NewReader(updateJSON))
	req.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expected := `{"id":1,"name":"Zak Updated","age":30}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, w.Body.String())
	}
}
