package models

type (
	GetPlants struct {
		PlantSlug  string `json:"plants_slug"`
		PlantName  string `json:"plants_name"`
		PlantBio   string `json:"plants_bio"`
		PlantPrice int    `json:"plants_price"`
		PlantStock int    `json:"plants_stock"`
	}
	GetPlantDetail struct {
		PlantSlug   string `json:"plants_slug"`
		PlantName   string `json:"plants_name"`
		PlantBio    string `json:"plants_bio"`
		PlantPrice  string `json:"plants_price"`
		PlantStock  string `json:"plants_stock"`
		PlantDetail string `json:"plants_detail"`

		// Props
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
	}
)
