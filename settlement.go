package razorpay

// Settlement struct contains the properties of settlement object
type Settlement struct {
	ID        string `json:"id"`
	Entity    string `json:"entity"`
	Amount    int64  `json:"amount"`
	Status    string `json:"status"`
	Fees      int64  `json:"fees"`
	Tax       int64  `json:"tax"`
	UTR       string `json:"utr"`
	CreatedAt int64  `json:"created_at"`
}

// New method will create a settlement object and return a pointer to it
func (o *Settlement) New() Resource {
	var obj = &Settlement{}
	return obj
}

// Endpoint method returns the endpoint of the resource
func (o *Settlement) Endpoint() string {
	return "/settlements"
}
