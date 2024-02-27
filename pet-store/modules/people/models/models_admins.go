package models

type (
	GetAdminsDetail struct {
		AdminSlug  string `json:"admin_slug"`
		AdminName  string `json:"admin_name"`
		AdminEmail string `json:"email"`
		AdminImage string `json:"admin_image"`

		// Props
		CreatedAt string `json:"created_at"`
	}
)
