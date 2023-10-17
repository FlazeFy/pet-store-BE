package models

type (
	GetAllTag struct {
		TagSlug string `json:"tag_slug"`
		TagName string `json:"tag_name"`
	}
)
