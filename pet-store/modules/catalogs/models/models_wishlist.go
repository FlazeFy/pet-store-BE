package models

type (
	GetMyWishlist struct {
		CatalogType  string `json:"catalog_type"`
		CatalogId    string `json:"catalog_id"`
		CatalogSlug  string `json:"catalog_slug"`
		CatalogName  string `json:"catalog_name"`
		CatalogPrice int    `json:"catalog_price"`

		// Props
		CreatedAt string `json:"created_at"`
	}
	PostWishlist struct {
		CatalogType string `json:"catalog_type"`
		CatalogId   string `json:"catalog_id"`
	}
)
