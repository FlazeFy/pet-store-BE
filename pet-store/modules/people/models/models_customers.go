package models

type (
	GetCustomers struct {
		CustomerSlug  string `json:"customer_slug"`
		CustomerName  string `json:"customer_name"`
		CustomerEmail string `json:"email"`
	}
)
