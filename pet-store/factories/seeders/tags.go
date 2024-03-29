package seeders

import (
	"fmt"
	"pet-store/modules/systems/models"
	"pet-store/modules/systems/repositories"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"

	"github.com/bxcodec/faker/v3"
)

func SeedTags(total int, showRes bool) {
	var obj models.GetAllTag
	idx := 0
	var logs string

	for idx < total {
		// Data
		name := faker.Word()
		obj.TagSlug = generator.GetSlug(name)
		obj.TagName = name

		result, err := repositories.PostTag(obj)
		if err != nil {
			fmt.Println(err.Error())
		}

		if showRes {
			fmt.Println(result.Data)
			if strData, ok := result.Data.(string); ok {
				logs += strData + "\n"
			}
		}
		idx++
	}

	if showRes {
		response.ResponsePrinter("txt", "seeder_tags", logs)
	}
}
