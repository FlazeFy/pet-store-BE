package models

type (
	GetAnimals struct {
		ID           string `json:"id"`
		AnimalSlug   string `json:"animal_slug"`
		AnimalName   string `json:"animal_name"`
		AnimalBio    string `json:"animal_bio"`
		AnimalGender string `json:"animal_gender"`
		AnimalPrice  string `json:"animal_price"`
		AnimalStock  string `json:"animal_stock"`
	}
)
