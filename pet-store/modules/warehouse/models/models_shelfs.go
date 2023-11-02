package models

type (
	GetAllShelfs struct {
		ShelfCategoryName string `json:"shelfs_category_name"`
		ShelfSlug         string `json:"shelfs_slug"`
		ShelfName         string `json:"shelfs_name"`
		ShelfTag          string `json:"shelfs_tag"`
	}
)
