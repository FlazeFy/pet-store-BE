package repositories

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"pet-store/modules/people/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/auth"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
)

func GetAllCustomer(page, pageSize int, path string, view string) (response.Response, error) {
	// Declaration
	var obj models.GetCustomers
	var arrobj []models.GetCustomers
	var res response.Response
	var baseTable = "customers"
	var sqlStatement string
	var where string

	// Nullable column
	var CustomerEmail sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "customers_name")

	if view == "all" || view == "regist" || view == "unregist" {
		if view == "all" {
			where = " "
		} else if view == "regist" {
			where = "WHERE email is NOT NULL"
		} else if view == "unregist" {
			where = "WHERE email is NULL"
		}

		sqlStatement = "SELECT " + selectTemplate + ", email " +
			"FROM " + baseTable + " " +
			where + " " +
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
				&obj.CustomerSlug,
				&obj.CustomerName,
				&CustomerEmail,
			)

			if err != nil {
				return res, err
			}

			obj.CustomerEmail = converter.CheckNullString(CustomerEmail)

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
	} else {
		res.Status = http.StatusBadRequest
		res.Message = "View not available"
		res.Data = nil
	}

	return res, nil
}

func GetMyProfile(path string, token string) (response.Response, error) {
	// Declaration
	var obj models.GetCustomersDetail
	var arrobj []models.GetCustomersDetail
	var res response.Response
	var baseTable = "customers"
	var sqlStatement string

	// Nullable column
	var CustomerEmail sql.NullString
	var CustomerImage sql.NullString
	var CustomerInterest sql.NullString

	// Converted column
	var IsNotifable string

	joinAuth := auth.GetAuthQuery(baseTable, token)

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)

	sqlStatement = "SELECT " + selectTemplate + ", email, customers_interest, customers_image, is_notifable, customers.created_at " +
		"FROM " + baseTable + " " +
		joinAuth

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
			&obj.CustomerSlug,
			&obj.CustomerName,
			&CustomerEmail,
			&CustomerInterest,
			&CustomerImage,
			&IsNotifable,
			&obj.CreatedAt,
		)

		if err != nil {
			return res, err
		}

		// Nullable column
		obj.CustomerEmail = converter.CheckNullString(CustomerEmail)
		obj.CustomerInterest = converter.CheckNullString(CustomerInterest)
		obj.CustomerImage = converter.CheckNullString(CustomerImage)

		obj.IsNotifable = converter.ConvertStringBool(IsNotifable)

		arrobj = append(arrobj, obj)
	}

	total := len(arrobj)

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
