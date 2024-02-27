package repositories

import (
	"net/http"
	"pet-store/modules/catalogs/models"
	"pet-store/packages/database"
	"pet-store/packages/helpers/auth"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"time"

	"github.com/google/uuid"
)

func PostWishlist(d models.PostWishlist, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	userId := auth.GetUserIdByToken(token)

	// Data
	id := uuid.Must(uuid.NewRandom())

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, catalog_type, catalog_id, created_at, created_by) " +
		"VALUES (?,?,?,?,?)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, d.CatalogType, d.CatalogId, dt, userId)
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
	res.Data = map[string]interface{}{
		"id":            id,
		"data":          d,
		"rows_affected": rowsAffected,
	}

	return res, nil
}
