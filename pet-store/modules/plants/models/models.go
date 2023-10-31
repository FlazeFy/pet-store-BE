package models

type (
	GetPlants struct {
		ID         string `json:"id"`
		PlantSlug  string `json:"plant_slug"`
		PlantName  string `json:"plant_name"`
		PlantBio   string `json:"plant_bio"`
		PlantPrice string `json:"plant_price"`
		PlantStock string `json:"plant_stock"`
	}
)
