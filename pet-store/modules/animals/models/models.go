package models

type (
	GetAnimals struct {
		AnimalSlug   string `json:"animals_slug"`
		AnimalName   string `json:"animals_name"`
		AnimalBio    string `json:"animals_bio"`
		AnimalGender string `json:"animals_gender"`
		AnimalPrice  int    `json:"animals_price"`
		AnimalStock  int    `json:"animals_stock"`
	}
	GetAnimalDetail struct {
		AnimalId       string `json:"animals_id"`
		AnimalSlug     string `json:"animals_slug"`
		AnimalName     string `json:"animals_name"`
		AnimalBio      string `json:"animals_bio"`
		AnimalGender   string `json:"animals_gender"`
		AnimalPrice    int    `json:"animals_price"`
		AnimalStock    string `json:"animals_stock"`
		AnimalDateBorn string `json:"animals_date_born"`
		AnimalDetail   string `json:"animals_detail"`

		// Props
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
	}
)
