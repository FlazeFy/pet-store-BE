package repositories

import (
	"database/sql"
	"math"
	"net/http"
	"pet-store/modules/warehouse/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
)

func GetAllShelf(page, pageSize int, path string, view string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetAllShelfs
	var arrobj []models.GetAllShelfs
	var res response.Response
	var baseTable = "shelfs"
	var secondTable = "dictionaries"
	var sqlStatement string

	// Nullable column
	var ShelfTag sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	firstLogicWhere := builders.GetTemplateLogic(view)
	whereActive := baseTable + firstLogicWhere
	join1 := builders.GetTemplateJoin("total", baseTable, "shelfs_category", secondTable, "id", true)
	order := "shelfs_name DESC " + ord

	sqlStatement = "SELECT " + selectTemplate + " " +
		"FROM " + baseTable + " " +
		join1 +
		"WHERE " + whereActive +
		"ORDER BY " + order +
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
			&obj.ShelfCategoryName,
			&obj.ShelfSlug,
			&obj.ShelfName,
			&ShelfTag,
		)

		if err != nil {
			return res, err
		}

		obj.ShelfTag = converter.CheckNullString(ShelfTag)

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
