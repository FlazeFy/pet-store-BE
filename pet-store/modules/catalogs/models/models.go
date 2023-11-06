package models

type (
	GetCatalogs struct {
		CatalogType   string `json:"catalog_type"`
		CatalogSlug   string `json:"catalog_slug"`
		CatalogName   string `json:"catalog_name"`
		CatalogBio    string `json:"catalog_bio"`
		CatalogGender string `json:"catalog_gender"`
		CatalogPrice  int    `json:"catalog_price"`
		CatalogStock  string `json:"catalog_stock"`
	}
)
