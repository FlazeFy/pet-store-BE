package repositories

import (
	"net/http"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
)

func HardDelDoctorBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "doctors"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, baseTable+"_slug")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(slug)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "permanently delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
