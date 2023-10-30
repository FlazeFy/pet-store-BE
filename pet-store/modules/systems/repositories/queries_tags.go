package repositories

import (
	"fmt"
	"math"
	"net/http"
	"pet-store/modules/systems/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
)

func GetAllTag(page, pageSize int, path string) (response.Response, error) {
	// Declaration
	var obj models.GetAllTag
	var arrobj []models.GetAllTag
	var res response.Response
	var baseTable = "tags"
	var sqlStatement string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	order := "tags_name DESC "

	sqlStatement = "SELECT " + selectTemplate + " " +
		"FROM " + baseTable + " " +
		"ORDER BY " + order +
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
			&obj.TagSlug,
			&obj.TagName,
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