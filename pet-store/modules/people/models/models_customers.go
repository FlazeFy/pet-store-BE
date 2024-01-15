package models

type (
	GetCustomers struct {
		CustomerSlug  string `json:"customer_slug"`
		CustomerName  string `json:"customer_name"`
		CustomerEmail string `json:"email"`
	}
	GetCustomersDetail struct {
		CustomerSlug     string `json:"customer_slug"`
		CustomerName     string `json:"customer_name"`
		CustomerEmail    string `json:"email"`
		CustomerInterest string `json:"customers_interest"`
		CustomerImage    string `json:"customers_image"`
		IsNotifable      bool   `json:"is_notifable"`

		// Props
		CreatedAt string `json:"created_at"`
	}
)
