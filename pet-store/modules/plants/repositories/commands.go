package repositories

import (
	"net/http"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func HardDelPlantBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "plants"
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

func SoftDelPlantBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "plants"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Command builder
	sqlStatement = builders.GetTemplateCommand("soft_delete", baseTable, baseTable+"_slug")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(dt, slug)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func PostPlant(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "plants"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	plantName := data.FormValue("plants_name")
	plantlSlug := generator.GetSlug(plantName)
	plantBio := data.FormValue("plants_bio")
	plantDetail := data.FormValue("plants_detail")
	plantPrice := data.FormValue("plants_price")
	plantStock := data.FormValue("plants_stock")

	// Template
	props := builders.GetTemplateSelect("properties_full", nil, nil)

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, plants_slug, plants_name, plants_bio, plants_detail, plants_price, plants_stock, " + props + ", deleted_at, deleted_by) " +
		"VALUES (?,?,?,?,?,?,?,?,?,null,null,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, plantlSlug, plantName, plantBio, plantDetail, plantPrice, plantStock, dt, "1")
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
