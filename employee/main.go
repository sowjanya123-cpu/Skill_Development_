package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Designation string  `json:"designation"`
	Technology  string  `json:"technology"`
	Commission  int     `json:"commission"`
	Salary      float32 `json:"salary"`
	Phone       int64   `json:"phone"`
}

var Employees = []Employee{
	{Id: "1001", Name: "Aliya", Designation: "Developer", Technology: "Java", Commission: 50000, Salary: 15000, Phone: 9353938926},
	{Id: "1002", Name: "John", Designation: "Manager", Technology: "Go", Commission: 60000, Salary: 20000, Phone: 9353938927},
}

func readAllEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, Employees)
}

func readEmployeeById(c *gin.Context) {
	id := c.Param("id")
	for _, employee := range Employees {
		if employee.Id == id {
			c.JSON(http.StatusOK, employee)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func createEmployee(c *gin.Context) {
	var newEmployee Employee
	err := c.BindJSON(&newEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	newEmployee.Id = fmt.Sprintf("%d", len(Employees)+1)

	Employees = append(Employees, newEmployee)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Employee created successfully",
		"employee": newEmployee,
	})
}

func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var updatedEmployee Employee
	err := c.BindJSON(&updatedEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	for i, employee := range Employees {
		if employee.Id == id {
			Employees[i] = updatedEmployee
			Employees[i].Id = id
			c.JSON(http.StatusOK, gin.H{
				"message":  "Employee updated successfully",
				"employee": Employees[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

// Delete an employee by ID
func deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	for i, employee := range Employees {
		if employee.Id == id {
			Employees = append(Employees[:i], Employees[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Employee deleted successfully",
				"id":      id,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func main() {

	r := gin.Default()

	// Routes (API Endpoints)
	r.GET("/employees", readAllEmployees)
	r.GET("/employees/:id", readEmployeeById)
	r.POST("/employees", createEmployee)
	r.PUT("/employees/:id", updateEmployee)
	r.DELETE("/employees/:id", deleteEmployee)

	// Start the server on port 8080
	r.Run(":8081")
}