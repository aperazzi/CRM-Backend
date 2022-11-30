package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// customer type represents data about a customer in the CRM.
type customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// customers slice to seed the crm data.
var customers = []customer{
	{ID: 1, Name: "Andrea", Role: "Software Engineer", Email: "andrea@company.com", Phone: "+3466578875", Contacted: false},
	{ID: 2, Name: "Adrian", Role: "Manager", Email: "adrian@hello.com", Phone: "+39993899487", Contacted: true},
	{ID: 3, Name: "Loren", Role: "SEO Specialist", Email: "lorean@seo.com", Phone: "+34773879833", Contacted: false},
	{ID: 4, Name: "Elisa", Role: "Marketing Manager", Email: "elisa@marketing.com", Phone: "+41884788493", Contacted: true},
	{ID: 5, Name: "Roby", Role: "UX Designer", Email: "ruby@design.com", Phone: "+346169614595", Contacted: true},
}

// getIndex serve static HTML
func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// getCustomers respond with the list of all the customers as JSON.
func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

// addCustomer add a customer from JSON received in the request body.
func addCustomer(c *gin.Context) {
	var newCustomer customer

	// Bind the received JSON to newCustomers.
	if err := c.BindJSON(&newCustomer); err != nil {
		return
	}

	// Loop over the list of customers, looking if already exist a customer with the
	// same ID as the new customer.
	for _, customer := range customers {
		if customer.ID == newCustomer.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "Customer with this ID already exist"})
			return
		}
	}

	// Add the new customer to the slice.
	customers = append(customers, newCustomer)
	c.IndentedJSON(http.StatusCreated, newCustomer)
}

// getCustomer locates the customer whose ID value matches the id parameter
// sent by the client, then returns that customer as a response.
func getCustomer(c *gin.Context) {

	// Get the id from the client and convert to int
	var id, _ = strconv.Atoi(c.Param("id"))

	// Loop over the list of customers, looking for a customer
	// whose ID value matches the parameter.
	for _, customer := range customers {
		if customer.ID == id {
			c.IndentedJSON(http.StatusOK, customer)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Customer not found"})
}

// deleteCustomer delete the customer whose ID value matches the id parameter
// and return the customers slice
func deleteCustomer(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))

	for i, customer := range customers {
		if customer.ID == id {
			customers = removeCustomer(customers, i)
			c.IndentedJSON(http.StatusOK, customers)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Customer not found"})
}

func removeCustomer(s []customer, i int) []customer {
	if i >= len(s) || i < 0 {
		return nil
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// updateCustomer update the customer whose ID value matches the id parameter
// and return the updated customers slice
func updateCustomer(c *gin.Context) {
	// Get the id from the client and convert to int
	var id, _ = strconv.Atoi(c.Param("id"))

	var updatedCustomer customer

	// Bind the received JSON to newCustomers.
	if err := c.BindJSON(&updatedCustomer); err != nil {
		return
	}

	// Loop over the list of customers, looking if already exist a customer with the
	// same ID as the new customer.
	for i, customer := range customers {
		if customer.ID == id {
			customers[i] = updatedCustomer
			c.IndentedJSON(http.StatusOK, customers)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The customer not exist"})
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")

	router.GET("/", getIndex)
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomer)
	router.POST("/customers", addCustomer)
	router.DELETE("/customers/:id", deleteCustomer)
	router.PATCH("/customers/:id", updateCustomer)

	router.Run("localhost:3000")
}
