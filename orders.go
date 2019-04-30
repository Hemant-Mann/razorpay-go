package razorpay

import (
	"encoding/json"
)

// Order struct represents the information of the order
type Order struct {
	ID       string `json:"id"`
	Entity   string `json:"entity"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`
	Status   string `json:"status"`
	Attempts int    `json:"attempts"`

	Notes     map[string]string `json:"notes,omitempty"`
	CreatedAt int               `json:"created_at"`
}

// OrderParams struct represents the information to create a order
type OrderParams struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`

	PaymentCapture bool `json:"payment_capture"`

	Notes map[string]string `json:"notes"`
}

// New method will create a order object and return a pointer to it
func (o *Order) New() Resource {
	var obj = &Order{}
	return obj
}

// Endpoint method returns the endpoint of the resource
func (o *Order) Endpoint() string {
	return "/orders"
}

// Create method will try to create a order on razorpay
func (o *Order) Create(params *OrderParams, client *Client) (*Order, error) {
	var body, _ = json.Marshal(params)
	resp, err := client.Post(o.Endpoint(), body)

	or, err := sendResp(resp, err, o)
	return or.(*Order), err
}

// FindOne tries to find the order with given id
func (o *Order) FindOne(id string, client *Client) (*Order, error) {
	resp, err := client.Get(o.Endpoint() + "/" + id)

	or, err := sendResp(resp, err, o)
	return or.(*Order), err
}
