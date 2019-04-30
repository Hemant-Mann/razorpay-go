package razorpay

// Card struct represents the information of the card
type Card struct {
	ID            string `json:"id"`
	Entity        string `json:"entity"`
	Name          string `json:"name"`
	Last4         string `json:"last4"`
	Network       string `json:"network"`
	Type          string `json:"type"`
	Issuer        string `json:"issuer"`
	International string `json:"international"`
	EMI           string `json:"emi"`
}

// CardParams struct represents the information to create a card
type CardParams struct {
	ID            string `json:"id"`
	Entity        string `json:"entity"`
	Name          string `json:"name"`
	Last4         string `json:"last4"`
	Network       string `json:"network"`
	Type          string `json:"type"`
	Issuer        string `json:"issuer"`
	International string `json:"international"`
	EMI           string `json:"emi"`
}

// New method will create a card object and return a pointer to it
func (c *Card) New() Resource {
	var obj = &Card{}
	return obj
}

// Endpoint method returns the endpoint of the resource
func (c *Card) Endpoint() string {
	return "/cards"
}

// FindOne tries to find the card with given id
func (c *Card) FindOne(id string, client *Client) (*Card, error) {
	resp, err := client.Get(c.Endpoint() + "/" + id)

	obj, err := sendResp(resp, err, c)
	return obj.(*Card), err
}
