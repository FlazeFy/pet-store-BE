package repositories

import (
	"database/sql"
	"math"
	"net/http"
	"pet-store/modules/people/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
	"strconv"
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

func GetAllDoctors(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetDoctors
	var arrobj []models.GetDoctors
	var res response.Response
	var baseTable = "doctors"
	var sqlStatement string

	// Nullable column
	var DoctorDesc sql.NullString
	var DoctorImageUrl sql.NullString
	var UpdatedAt sql.NullString

	// Converted column
	var IsReady string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "doctors_name")

	sqlStatement = "SELECT " + selectTemplate + ", doctors_desc, is_ready, doctors_img_url, created_at, updated_at " +
		"FROM " + baseTable + " " +
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
			&obj.DoctorSlug,
			&obj.DoctorName,
			&DoctorDesc,
			&IsReady,
			&DoctorImageUrl,
			&obj.CreatedAt,
			&UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intIsReady, err := strconv.Atoi(IsReady)
		if err != nil {
			return res, err
		}

		// Nullable column
		obj.DoctorDesc = converter.CheckNullString(DoctorDesc)
		obj.DoctorImageUrl = converter.CheckNullString(DoctorImageUrl)
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)

		obj.IsReady = intIsReady

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
