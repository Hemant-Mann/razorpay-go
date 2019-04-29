package razorpay

import (
	"encoding/json"
	"net/http"
)

// Order struct represents the information of the order
type Order struct {
	ID       string `json:"id"`
	Entity   string `json:"entity"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`
	Status   string `json:"status"`
	Attempts int    `json:"attempts"`

	Notes     map[string]string `json:"notes"`
	CreatedAt int               `json:"created_at"`
}

// OrderParams struct represents the information to create a order
type OrderParams struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`

	PaymentCapture bool `json:"payment_capture"`

	Notes map[string]string `json:"notes"`
}

func (c *Order) sendResp(resp *http.Response, err error) (*Order, error) {
	var newCust = &Order{}
	if err != nil {
		return newCust, err
	}
	body, readErr := readBody(resp)
	if readErr != nil {
		return newCust, readErr
	}
	parseError := json.Unmarshal(body, newCust)
	return newCust, parseError
}

// Create method will try to create a order on razorpay
func (c *Order) Create(params *OrderParams, client *Client) (*Order, error) {
	var body, _ = json.Marshal(params)
	resp, err := client.Post("/orders", body)

	return c.sendResp(resp, err)
}

// FindOne tries to find the order with given id
func (c *Order) FindOne(id string, client *Client) (*Order, error) {
	resp, err := client.Get("/orders/" + id)
	return c.sendResp(resp, err)
}
