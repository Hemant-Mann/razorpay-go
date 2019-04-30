package razorpay

import (
	"encoding/json"
)

// Customer struct represents the information of the customer
type Customer struct {
	ID     string `json:"id"`
	Entity string `json:"entity"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	// Contact number of the customer
	Contact   string            `json:"contact"`
	Notes     map[string]string `json:"notes"`
	CreatedAt int               `json:"created_at"`
}

// CustomerParams struct represents the information to create a customer
type CustomerParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	// Contact number of the customer
	Contact string            `json:"contact"`
	Notes   map[string]string `json:"notes"`
}

func (c *Customer) New() Resource {
	var obj = &Customer{}
	return obj
}

func (c *Customer) Endpoint() string {
	return "/customers"
}

// Create method will try to create a customer on razorpay
func (c *Customer) Create(params *CustomerParams, client *Client) (*Customer, error) {
	var body, _ = json.Marshal(params)
	resp, err := client.Post("/customers", body)

	cust, err := sendResp(resp, err, c)
	return cust.(*Customer), err
}

// FindOne tries to find the customer with given id
func (c *Customer) FindOne(id string, client *Client) (*Customer, error) {
	resp, err := client.Get("/customers/" + id)
	cust, err := sendResp(resp, err, c)
	return cust.(*Customer), err
}
