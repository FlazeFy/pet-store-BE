package repositories

import (
	"math"
	"net/http"
	"pet-store/modules/catalogs/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/auth"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
	"strconv"
)

func GetMyWishlist(page, pageSize int, path string, ord string, token string) (response.Response, error) {
	// Declaration
	var obj models.GetMyWishlist
	var arrobj []models.GetMyWishlist
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string

	// Converted column
	var CatalogPrice string

	var col1 = "slug"
	var col2 = "name"
	var col3 = "price"

	joinAuth := auth.GetAuthQuery(baseTable, token)
	catalogSlugQuery := builders.GetTemplateLogic("multi_content", &col1)
	catalogNameQuery := builders.GetTemplateLogic("multi_content", &col2)
	catalogPriceQuery := builders.GetTemplateLogic("multi_content", &col3)

	sqlStatement = "SELECT wishlists.catalog_type, catalog_id, " +
		catalogSlugQuery + ", " +
		catalogNameQuery + ", " +
		catalogPriceQuery + ", " +
		"wishlists.created_at " +
		"FROM " + baseTable + " " +
		"LEFT JOIN animals ON wishlists.catalog_id = animals.id AND wishlists.catalog_type = 'animals' " +
		"LEFT JOIN plants ON wishlists.catalog_id = plants.id AND wishlists.catalog_type = 'plants' " +
		"LEFT JOIN goods ON wishlists.catalog_id = goods.id AND wishlists.catalog_type = 'goods' " +
		joinAuth +
		"ORDER BY wishlists.created_at " +
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
			&obj.CatalogId,
			&obj.CatalogSlug,
			&obj.CatalogName,
			&CatalogPrice,
			&obj.CreatedAt,
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

	total := len(arrobj)
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

func GetCheckWishlist(token, slug, types string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "wishlists"
	var sqlStatement string
	var rowCount int

	var col1 = "slug"
	joinAuth := auth.GetAuthQuery(baseTable, token)
	catalogSlugQuery := builders.GetTemplateLogic("multi_content", &col1)

	sqlStatement = "SELECT wishlists.id, " +
		catalogSlugQuery + " " +
		"FROM " + baseTable + " " +
		"LEFT JOIN animals ON wishlists.catalog_id = animals.id AND wishlists.catalog_type = 'animals' " +
		"LEFT JOIN plants ON wishlists.catalog_id = plants.id AND wishlists.catalog_type = 'plants' " +
		"LEFT JOIN goods ON wishlists.catalog_id = goods.id AND wishlists.catalog_type = 'goods' " +
		joinAuth + " " +
		"HAVING catalog_slug ='" + slug + "' " +
		"LIMIT 1"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	for rows.Next() {
		rowCount++

		if err != nil {
			return res, err
		}
	}

	total := rowCount

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = false
	} else {
		res.Data = true
	}

	return res, nil
}
