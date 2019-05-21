package razorpay

import (
	"encoding/json"
)

// VAReceiverType currently it is the only receiver type supported by virtual account endpoint
const VAReceiverType = "bank_account"

// Receiver struct for holding bank account info
type Receiver struct {
	ID        string `json:"id"`
	Entity    string `json:"entity"`
	Name      string `json:"name"`
	AccountNo string `json:"account_number"`
	IFSC      string `json:"ifsc"`
}

// VirtualAccount struct holds info regarding the virtual account
type VirtualAccount struct {
	ID          string     `json:"id"`
	Entity      string     `json:"entity"`
	Description string     `json:"description"`
	CustomerID  string     `json:"customer_id"`
	Status      string     `json:"status"`
	AmountPaid  int        `json:"amount_paid"`
	Receivers   []Receiver `json:"receivers"`

	Notes     interface{} `json:"notes,omitempty"`
	CreatedAt int         `json:"created_at"`
}

// VirtualAccountParams struct contains the params required for creating a virtual account
type VirtualAccountParams struct {
	Description string `json:"description"`
	CustomerID  string `json:"customer_id"`

	Receivers VirtualAccountReceiver `json:"receivers"`

	Notes map[string]string `json:"notes,omitempty"`
}

// VirtualAccountReceiver struct contain the receiver properties
type VirtualAccountReceiver struct {
	Types []string `json:"types"`
}

// New method will create a customer object and return a pointer to it
func (va *VirtualAccount) New() Resource {
	var obj = &VirtualAccount{}
	return obj
}

// Endpoint method returns the endpoint of the resource
func (va *VirtualAccount) Endpoint() string {
	return "/virtual_accounts"
}

// Create method will try to create a customer on razorpay
func (va *VirtualAccount) Create(params *VirtualAccountParams, client *Client) (*VirtualAccount, error) {
	var body, _ = json.Marshal(params)
	resp, err := client.Post(va.Endpoint(), body)

	vacc, err := sendResp(resp, err, va)
	return vacc.(*VirtualAccount), err
}

// FindOne tries to find the customer with given id
func (va *VirtualAccount) FindOne(id string, client *Client) (*VirtualAccount, error) {
	resp, err := client.Get(va.Endpoint() + "/" + id)
	vacc, err := sendResp(resp, err, va)
	return vacc.(*VirtualAccount), err
}

// Close method will close the virtual account
func (va *VirtualAccount) Close(id string, client *Client) (*VirtualAccount, error) {
	var params = map[string]string{"status": "closed"}
	var body, _ = json.Marshal(params)
	resp, err := client.Patch(va.Endpoint()+"/"+id, body)

	vacc, err := sendResp(resp, err, va)
	return vacc.(*VirtualAccount), err
}
