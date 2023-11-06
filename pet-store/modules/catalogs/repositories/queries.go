package repositories

import (
	"math"
	"net/http"
	"pet-store/modules/catalogs/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
	"strconv"
)

func GetAllCatalogs(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetCatalogs
	var arrobj []models.GetCatalogs
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Converted column
	var CatalogPrice string

	sqlStatement = "SELECT 'animal' as catalog_type, animals_slug as catalog_slug, animals_name as catalog_name, animals_bio as catalog_bio, animals_gender as catalog_gender, animals_price as catalog_price, animals_stock as catalog_stock " +
		"FROM animals " +
		"UNION ALL " +
		"SELECT 'plant' as catalog_type, plants_slug as catalog_slug, plants_name as catalog_name, plants_bio as catalog_bio, 0 as catalog_gender, plants_price as catalog_price, plants_stock as catalog_stock " +
		"FROM plants " +
		"ORDER BY catalog_name " +
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
			&obj.CatalogType,
			&obj.CatalogSlug,
			&obj.CatalogName,
			&obj.CatalogBio,
			&obj.CatalogGender,
			&CatalogPrice,
			&obj.CatalogStock,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intCatalogPrice, err := strconv.Atoi(CatalogPrice)
		if err != nil {
			return res, err
		}

		obj.CatalogPrice = intCatalogPrice

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
