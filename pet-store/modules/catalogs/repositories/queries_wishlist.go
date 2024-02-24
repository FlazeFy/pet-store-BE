package repositories

import (
	"math"
	"net/http"
	"pet-store/modules/catalogs/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
	"strconv"
)

func GetMyWishlist(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetMyWishlist
	var arrobj []models.GetMyWishlist
	var res response.Response
	var baseTable = "animals"
	var sqlStatement string

	// Converted column
	var CatalogPrice string

	catalogSlugTmt := map[string]interface{}{
		"end_as":      "catalog_slug",
		"table_list":  []string{"animals", "plants", "goods"},
		"column_list": []string{"animals_slug", "plants_slug", "goods_slug"},
	}
	catalogPriceTmt := map[string]interface{}{
		"end_as":      "catalog_price",
		"table_list":  []string{"animals", "plants", "goods"},
		"column_list": []string{"animals_price", "plants_price", "goods_price"},
	}
	catalogNameTmt := map[string]interface{}{
		"end_as":      "catalog_name",
		"table_list":  []string{"animals", "plants", "goods"},
		"column_list": []string{"animals_name", "plants_name", "goods_name"},
	}

	catalogSlugQuery := converter.MapToString(catalogSlugTmt)
	catalogNameQuery := converter.MapToString(catalogNameTmt)
	catalogPriceQuery := converter.MapToString(catalogPriceTmt)

	catalogSlugFinalQuery := builders.GetTemplateLogic("multi_content", &catalogSlugQuery)
	catalogNameFinalQuery := builders.GetTemplateLogic("multi_content", &catalogNameQuery)
	catalogPriceFinalQuery := builders.GetTemplateLogic("multi_content", &catalogPriceQuery)

	sqlStatement = "SELECT wishlists.catalog_type, catalog_id, " +
		catalogSlugFinalQuery + " " +
		catalogNameFinalQuery + " " +
		catalogPriceFinalQuery + " " +
		",wishlists.created_at " +
		"FROM " + baseTable + " " +
		"LEFT JOIN animals ON wishlists.catalog_id = animals.id AND wishlists.catalog_type = 'animals' " +
		"LEFT JOIN plants ON wishlists.catalog_id = plants.id AND wishlists.catalog_type = 'plants' " +
		"LEFT JOIN goods ON wishlists.catalog_id = goods.id AND wishlists.catalog_type = 'goods' " +
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
