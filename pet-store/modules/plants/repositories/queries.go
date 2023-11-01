package repositories

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"pet-store/modules/plants/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
)

func GetAllPlants(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetPlants
	var arrobj []models.GetPlants
	var res response.Response
	var baseTable = "plants"
	var sqlStatement string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	order := "plants_name " + ord

	sqlStatement = "SELECT " + selectTemplate + ", plants_bio, plants_price, plants_stock " +
		"FROM " + baseTable + " " +
		"ORDER BY " + order + " " +
		"LIMIT ? OFFSET ?"

	// Exec
	con := database.CreateCon()
	offset := (page - 1) * pageSize
	rows, err := con.Query(sqlStatement, pageSize, offset)
	defer rows.Close()
	fmt.Println(sqlStatement)

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.PlantSlug,
			&obj.PlantName,
			&obj.PlantBio,
			&obj.PlantPrice,
			&obj.PlantStock,
		)

		if err != nil {
			return res, err
		}

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

func GetPlantDetailBySlug(path string, slug string) (response.Response, error) {
	// Declaration
	var obj models.GetPlantDetail
	var arrobj []models.GetPlantDetail
	var res response.Response
	var baseTable = "plants"
	var sqlStatement string

	// Nullable column
	var UpdatedAt sql.NullString
	var UpdatedBy sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	propsTemplate := builders.GetTemplateSelect("properties_full", nil, nil)

	sqlStatement = "SELECT " + selectTemplate + ", plants_bio, plants_price, plants_stock, plants_detail, " + propsTemplate + " " +
		"FROM " + baseTable + " " +
		"WHERE plants_slug = '" + slug + "'"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	fmt.Println(sqlStatement)

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.PlantSlug,
			&obj.PlantName,
			&obj.PlantBio,
			&obj.PlantPrice,
			&obj.PlantStock,
			&obj.PlantDetail,

			// Props
			&obj.CreatedAt,
			&obj.CreatedBy,
			&UpdatedAt,
			&UpdatedBy,
		)

		if err != nil {
			return res, err
		}

		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)
		obj.UpdatedBy = converter.CheckNullString(UpdatedBy)

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
