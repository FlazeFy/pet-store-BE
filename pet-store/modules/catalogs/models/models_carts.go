package models

type (
	GetMyCart struct {
		IsPaid    bool   `json:"is_paid"`
		CreatedAt string `json:"created_at"`
		PaidAt    string `json:"paid_at"`
	}
)
