package repositories

import (
	"math"
	"net/http"
	"pet-store/modules/catalogs/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
)

func GetMyCart(page, pageSize int, path string, user_id string) (response.Response, error) {
	// Declaration
	var obj models.GetMyCart
	var arrobj []models.GetMyCart
	var res response.Response
	var baseTable = "carts"
	var sqlStatement string

	// Converted column
	var IsPaid string

	// Query builder
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "customers_name")

	user_id = "1" // fon now

	sqlStatement = "SELECT is_paid, created_at, paid_at " +
		"FROM carts " +
		"WHERE created_by = '" + user_id + "' " +
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
			&IsPaid,
			&obj.CreatedAt,
			&obj.PaidAt,
		)

		if err != nil {
			return res, err
		}

		// Converted
		boolIsPad := converter.ConvertStringBool(IsPaid)
		if err != nil {
			return res, err
		}

		obj.IsPaid = boolIsPad

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
