# Set base URL
$baseUrl = "http://localhost:8080"

Write-Host "`n=== HEALTH CHECK ==="
Invoke-RestMethod -Uri "$baseUrl/healthcheck" -Method GET

Write-Host "`n=== GET all students (should be empty initially) ==="
Invoke-RestMethod -Uri "$baseUrl/students" -Method GET

Write-Host "`n=== ADD new student ==="
$newStudent = @{
    name = "John Smith"
    age = 21
} | ConvertTo-Json
Invoke-RestMethod -Uri "$baseUrl/students" -Method POST -Body $newStudent -ContentType "application/json"

Write-Host "`n=== ADD another student ==="
$secondStudent = @{
    name = "Sarah Connor"
    age = 29
} | ConvertTo-Json
Invoke-RestMethod -Uri "$baseUrl/students" -Method POST -Body $secondStudent -ContentType "application/json"

Write-Host "`n=== GET all students again (should show both) ==="
Invoke-RestMethod -Uri "$baseUrl/students" -Method GET

Write-Host "`n=== GET student by ID (1) ==="
Invoke-RestMethod -Uri "$baseUrl/students/1" -Method GET

Write-Host "`n=== UPDATE student (ID 1) ==="
$updateStudent = @{
    name = "John Updated"
    age = 22
} | ConvertTo-Json
Invoke-RestMethod -Uri "$baseUrl/students/1" -Method PATCH -Body $updateStudent -ContentType "application/json"

Write-Host "`n=== GET student ID 1 after update ==="
Invoke-RestMethod -Uri "$baseUrl/students/1" -Method GET

Write-Host "`n=== DELETE student (ID 2) ==="
Invoke-RestMethod -Uri "$baseUrl/students/2" -Method DELETE

Write-Host "`n=== GET all students again (should only show updated John) ==="
Invoke-RestMethod -Uri "$baseUrl/students" -Method GET

Write-Host "`n=== DONE ==="
