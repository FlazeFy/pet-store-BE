package repositories

import (
	"database/sql"
	"math"
	"net/http"
	"pet-store/modules/goods/models"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"pet-store/packages/utils/pagination"
	"strconv"
)

func GetAllGoods(page, pageSize int, path string, ord string) (response.Response, error) {
	// Declaration
	var obj models.GetGoods
	var arrobj []models.GetGoods
	var res response.Response
	var baseTable = "goods"
	var sqlStatement string

	// Nullable column
	var GoodsDesc sql.NullString

	// Converted column
	var GoodsPrice string
	var GoodsStock string

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "goods_name")

	sqlStatement = "SELECT " + selectTemplate + ", goods_desc, goods_category, goods_price, goods_stock " +
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
			&obj.GoodsSlug,
			&obj.GoodsName,
			&GoodsDesc,
			&obj.GoodsCategory,
			&GoodsPrice,
			&GoodsStock,
		)

		if err != nil {
			return res, err
		}

		// Nullable
		obj.GoodsDesc = converter.CheckNullString(GoodsDesc)

		// Converted
		intGoodsPrice, err := strconv.Atoi(GoodsPrice)
		intGoodsStock, err := strconv.Atoi(GoodsStock)
		if err != nil {
			return res, err
		}

		obj.GoodsPrice = intGoodsPrice
		obj.GoodsStock = intGoodsStock

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
