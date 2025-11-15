# Student API

A simple RESTful API to manage student records, built with **Go**, **Gin**, and **SQLite**. Fully containerized with **Docker** for easy local setup and portability.  

---

## Features

- Add a new student
- Get all students
- Get a student by ID
- Update student information
- Delete a student
- Healthcheck endpoint
- API versioning: `/api/v1/...`
- Logging for requests and errors
- Persistent SQLite database

---

## Prerequisites

- **Docker** installed on your machine
- (Optional) Postman or any HTTP client for testing API endpoints
- (Optional) Git to clone the repository

---

## Environment Variables

Copy `.env.example` to `.env`:

```powershell
copy .env.example .env

# Path to SQLite database inside container
STUDENT_DB=/app/data/students.db

# Port the API will run on
PORT=8080
```
The container will read these variables when it starts.

---

## Running the API with Docker

1. Build the Docker image
```powershell
docker build -t student-api ./web-service-gin
```
2. Run the container
```powershell
docker run -p 8080:8080 -v C:\Users\Zak\Documents\SRE_BootCamp\data:/app/data --env-file ./web-service-gin/.env student-api
```
3. Test endpoints

    We provide a PowerShell script to test all endpoints quickly.
    1. Make sure your API is running (e.g., http://localhost:8080)
    2. Open PowerShell in the project folder
    3. Run the script 
```
.\testAPICalls.ps1
``` 
The script performs:

    Get all students
    Get a student by ID
    Add a new student
    Update a student
    Delete a student
    Healthcheck

You can edit testAPICalls.ps1 to change IDs, names, or other request data as needed.

---

## API Versioning

All endpoints follow versioning
```bash
/api/v1/students
```
This allows for future improvements without breaking existing clients

---

## Running Test

Unit test are included for all handlers using an in-memory SQLite database
```powershell
cd web-service-gin
go test ./handlers -v
```
The tests do not require a real database and will not affect your persistent data.

---

## Logging

The API uses a **custom logger** implemented in `utils/logger.go` to structured logging for both requests and errors.  

### Features

- **InfoLogger** → general informational messages  
- **ErrorLogger** → logs errors with context  
- **Writes to file** → logs are persisted to `web-service-gin/logs/app.log`  
- Logs are also printed to the console for real-time debugging  

All API handlers use this logger to track key events and errors. For example:

- When fetching students: utils.InfoLogger.Printf("Fetched %d students", len(students))
- When an operation fails: utils.ErrorLogger.Printf("Failed to add student: %v", err)