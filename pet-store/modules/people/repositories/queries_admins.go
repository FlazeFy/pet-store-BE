package repositories

import (
	"database/sql"
	"net/http"
	"pet-store/modules/people/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/auth"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
)

func GetMyProfileAdmin(path string, token string) (response.Response, error) {
	// Declaration
	var obj models.GetAdminsDetail
	var arrobj []models.GetAdminsDetail
	var res response.Response
	var baseTable = "admins"
	var sqlStatement string

	// Nullable column
	var AdminEmail sql.NullString
	var AdminImage sql.NullString

	joinAuth := auth.GetAuthQuery(baseTable, token)

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)

	sqlStatement = "SELECT " + selectTemplate + ", email, admins_image " +
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
			&obj.AdminSlug,
			&obj.AdminName,
			&AdminEmail,
			&AdminImage,
		)

		if err != nil {
			return res, err
		}

		// Nullable column
		obj.AdminEmail = converter.CheckNullString(AdminEmail)
		obj.AdminImage = converter.CheckNullString(AdminImage)

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
