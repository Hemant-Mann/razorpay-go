package webhook

import (
	razorpay "github.com/hemant-mann/razorpay-go"
)

type paymentEntity struct {
	Entity razorpay.Payment `json:"entity"`
}

type orderEntity struct {
	Entity razorpay.Order `json:"entity"`
}

type vaEntity struct {
	Entity razorpay.VirtualAccount `json:"entity"`
}

type settlementEntity struct {
	Entity razorpay.Settlement `json:"entity"`
}

// Payload struct contains the payload data for the webhook
type Payload struct {
	Payment        paymentEntity    `json:"payment"`
	Order          orderEntity      `json:"order"`
	VirtualAccount vaEntity         `json:"virtual_account"`
	Settlement     settlementEntity `json:"settlement"`
}

// Webhook struct contains the webhook data sent by razorpay
type Webhook struct {
	Entity    string   `json:"entity"`
	AccountID string   `json:"account_id"`
	Event     string   `json:"event"`
	Contains  []string `json:"contains"`
	Payload   Payload  `json:"payload"`
	CreatedAt int64    `json:"created_at"`
}
