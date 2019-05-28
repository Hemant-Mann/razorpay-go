package razorpay

import (
	"encoding/json"
)

// Payment struct represents the information of the payment
type Payment struct {
	ID             string `json:"id"`
	Entity         string `json:"entity"`
	Amount         int64  `json:"amount"`
	Currency       string `json:"currency"`
	Status         string `json:"status"`
	OrderID        string `json:"order_id"`
	InvoiceID      string `json:"invoice_id"`
	International  bool   `json:"international"`
	Method         string `json:"method"`
	AmountRefunded int64  `json:"amount_refunded"`
	Captured       bool   `json:"captured"`
	Description    string `json:"description"`
	CardID         string `json:"card_id"`
	Card           Card   `json:"car"`
	Bank           string `json:"bank"`
	Email          string `json:"email"`
	Contact        string `json:"contact"`

	Notes     interface{} `json:"notes,omitempty"`
	Fee       int64       `json:"fee"`
	Tax       int64       `json:"tax"`
	CreatedAt int         `json:"created_at"`

	AmountPaid int64 `json:"amount_paid"`
	AmountDue  int64 `json:"amount_due"`
}

// PaymentParams struct represents the information to create a payment
type PaymentParams struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`

	PaymentCapture bool `json:"payment_capture"`

	Notes map[string]string `json:"notes"`
}

// New method will create a payment object and return a pointer to it
func (p *Payment) New() Resource {
	var obj = &Payment{}
	return obj
}

// Endpoint method returns the endpoint of the resource
func (p *Payment) Endpoint() string {
	return "/payments"
}

// GetNotes function is used to return a map of the notes
func (p *Payment) GetNotes() map[string]string {
	var resultMap = make(map[string]string)
	var jsonBytes, _ = json.Marshal(p.Notes)
	json.Unmarshal(jsonBytes, &resultMap)
	return resultMap
}

// FindOne tries to find the payment with given id
func (p *Payment) FindOne(id string, client *Client) (*Payment, error) {
	resp, err := client.Get(p.Endpoint() + "/" + id)

	obj, err := sendResp(resp, err, p)
	return obj.(*Payment), err
}
