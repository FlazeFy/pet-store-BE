package repositories

import (
	"net/http"
	"pet-store/modules/people/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
)

func GetAllDoctorSchedule() (response.Response, error) {
	// Declaration
	var obj models.GetDoctorSchedule
	var arrobj []models.GetDoctorSchedule
	var res response.Response
	var baseTable = "doctors"
	var sqlStatement string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	joinTemplate := builders.GetTemplateJoin("total", baseTable, "id", "doctors_schedule", "doctors_id", false)
	hour := "schedule_hour"
	selectHour := builders.GetFormulaQuery(&hour, "hour")

	sqlStatement = "SELECT " + selectTemplate + ", schedule_day, " + selectHour + " schedule_hour " +
		"FROM " + baseTable + " " +
		joinTemplate + " "
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
			&obj.DoctorSlug,
			&obj.DoctorName,
			&obj.ScheduleDay,
			&obj.ScheduleHour,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = arrobj

	return res, nil
}
