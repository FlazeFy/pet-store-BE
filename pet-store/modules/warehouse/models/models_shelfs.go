package models

type (
	GetAllShelfs struct {
		ShelfSlug string `json:"shelfs_slug"`
		ShelfName string `json:"shelfs_name"`
		ShelfTag  string `json:"shelfs_tag"`
	}
)
