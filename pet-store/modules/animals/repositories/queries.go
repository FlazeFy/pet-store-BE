package repositories

import (
	"database/sql"
	"math"
	"net/http"
	"pet-store/modules/animals/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
	"strconv"
)

func GetAllAnimals(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetAnimals
	var arrobj []models.GetAnimals
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Converted column
	var AnimalPrice string
	var AnimalStock string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "animals_name")

	sqlStatement = "SELECT " + selectTemplate + ", animals_bio, animals_gender, animals_price, animals_stock " +
		"FROM " + baseTable + " " +
		"ORDER BY " + order + " " +
		"LIMIT ? OFFSET ?"

	// Exec
	con := database.CreateCon()
	offset := (page - 1) * pageSize
	rows, err := con.Query(sqlStatement, pageSize, offset)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.AnimalSlug,
			&obj.AnimalName,
			&obj.AnimalBio,
			&obj.AnimalGender,
			&AnimalPrice,
			&AnimalStock,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intAnimalPrice, err := strconv.Atoi(AnimalPrice)
		intAnimalStock, err := strconv.Atoi(AnimalStock)
		if err != nil {
			return res, err
		}

		obj.AnimalPrice = intAnimalPrice
		obj.AnimalStock = intAnimalStock

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = map[string]interface{}{
			"current_page":   page,
			"data":           arrobj,
			"first_page_url": pagination.FirstPageURL,
			"from":           pagination.From,
			"last_page":      pagination.LastPage,
			"last_page_url":  pagination.LastPageURL,
			"links":          pagination.Links,
			"next_page_url":  pagination.NextPageURL,
			"path":           pagination.Path,
			"per_page":       pageSize,
			"prev_page_url":  pagination.PrevPageURL,
			"to":             pagination.To,
			"total":          total,
		}
	}

	return res, nil
}

func GetAnimalDetailBySlug(path string, slug string) (response.Response, error) {
	// Declaration
	var obj models.GetAnimalDetail
	var arrobj []models.GetAnimalDetail
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Nullable column
	var AnimalDateBorn sql.NullString
	var UpdatedAt sql.NullString
	var UpdatedBy sql.NullString
	var DeletedAt sql.NullString
	var DeletedBy sql.NullString

	// Converted column
	var AnimalPrice string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	propsTemplate := builders.GetTemplateSelect("properties_full", nil, nil)

	sqlStatement = "SELECT id, " + selectTemplate + ", animals_bio, animals_gender, animals_price, animals_stock, animals_date_born, animals_detail, " + propsTemplate + ",deleted_at, deleted_by " +
		"FROM " + baseTable + " " +
		"WHERE animals_slug = '" + slug + "'"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.AnimalId,
			&obj.AnimalSlug,
			&obj.AnimalName,
			&obj.AnimalBio,
			&obj.AnimalGender,
			&AnimalPrice,
			&obj.AnimalStock,
			&AnimalDateBorn,
			&obj.AnimalDetail,

			// Props
			&obj.CreatedAt,
			&obj.CreatedBy,
			&UpdatedAt,
			&UpdatedBy,
			&DeletedAt,
			&DeletedBy,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intAnimalPrice, err := strconv.Atoi(AnimalPrice)
		if err != nil {
			return res, err
		}

		obj.AnimalDateBorn = converter.CheckNullString(AnimalDateBorn)
		obj.AnimalPrice = intAnimalPrice
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)
		obj.UpdatedBy = converter.CheckNullString(UpdatedBy)
		obj.DeletedAt = converter.CheckNullString(DeletedAt)
		obj.DeletedBy = converter.CheckNullString(DeletedBy)

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
