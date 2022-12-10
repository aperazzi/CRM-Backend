package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// var id = uuid.New()

// customer type represents data about a customer in the CRM.
type customer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Contacted bool      `json:"contacted"`
}

// customers slice to seed the crm data.
var customers = []customer{
	{ID: uuid.New(), Name: "Andrea", Role: "Software Engineer", Email: "andrea@company.com", Phone: "+3466578875", Contacted: false},
	{ID: uuid.New(), Name: "Adrian", Role: "Manager", Email: "adrian@hello.com", Phone: "+39993899487", Contacted: true},
	{ID: uuid.New(), Name: "Loren", Role: "SEO Specialist", Email: "lorean@seo.com", Phone: "+34773879833", Contacted: false},
	{ID: uuid.New(), Name: "Elisa", Role: "Marketing Manager", Email: "elisa@marketing.com", Phone: "+41884788493", Contacted: true},
	{ID: uuid.New(), Name: "Roby", Role: "UX Designer", Email: "ruby@design.com", Phone: "+346169614595", Contacted: true},
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
	newCustomer.ID = uuid.New()

	if err := c.BindJSON(&newCustomer); err != nil {
		return
	}

	for _, customer := range customers {
		if customer.ID == newCustomer.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "Customer with this ID already exist"})
			return
		}
	}

	customers = append(customers, newCustomer)
	c.IndentedJSON(http.StatusCreated, newCustomer)
}

// getCustomer locates the customer whose ID value matches the id parameter
// sent by the client, then returns that customer as a response.
func getCustomer(c *gin.Context) {
	id := c.Param("id")

	for _, customer := range customers {
		customerID := (customer.ID).String()
		if customerID == id {
			c.IndentedJSON(http.StatusOK, customer)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Customer not found"})
}

// deleteCustomer delete the customer whose ID value matches the id parameter
// and return the customers slice
func deleteCustomer(c *gin.Context) {
	id := c.Param("id")

	for i, customer := range customers {
		customerID := (customer.ID).String()
		if customerID == id {
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
	id := c.Param("id")

	var updatedCustomer customer

	if err := c.BindJSON(&updatedCustomer); err != nil {
		return
	}

	for i, customer := range customers {
		customerID := (customer.ID).String()
		if customerID == id {
			updatedCustomer.ID = customer.ID
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

	err := router.Run("localhost:3000")
	if err != nil {
		return
	}
}
