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

func HardDelAnimalBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
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

func SoftDelAnimalBySlug(slug string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
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

func PostAnimal(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	animalName := data.FormValue("animals_name")
	animalSlug := generator.GetSlug(animalName)
	animalBio := data.FormValue("animals_bio")
	animalDateBorn := data.FormValue("animals_date_born")
	animalGender := data.FormValue("animals_gender")
	animalDetail := data.FormValue("animals_detail")
	animalPrice := data.FormValue("animals_price")
	animalStock := data.FormValue("animals_stock")

	// Template
	props := builders.GetTemplateSelect("properties_full", nil, nil)

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, animals_slug, animals_name, animals_bio, animals_date_born, animals_gender, animals_detail, animals_price, animals_stock, " + props + ", deleted_at, deleted_by) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?,?,null,null,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, animalSlug, animalName, animalBio, animalDateBorn, animalGender, animalDetail, animalPrice, animalStock, dt, "1")
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
