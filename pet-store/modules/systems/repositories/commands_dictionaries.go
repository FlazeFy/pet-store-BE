package repositories

import (
	"net/http"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func HardDelDctById(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "id")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
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

func PostDct(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Data
	id := uuid.Must(uuid.NewRandom())
	dctName := data.FormValue("dictionaries_name")
	dctType := data.FormValue("dictionaries_type")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, dictionaries_type, dictionaries_name) " +
		"VALUES (?,?,?)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, dctType, dctName)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
