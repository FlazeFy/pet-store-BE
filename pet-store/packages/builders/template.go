package builders

import (
	"fmt"
	"pet-store/packages/helpers/converter"
)

func GetTemplateSelect(name string, firstTable, secondTable *string) string {
	if name == "content_info" {
		return *firstTable + "_slug," + *firstTable + "_name"
	} else if name == "properties_time" {
		return "created_at, updated_at"
	} else if name == "properties_full" {
		return "created_at, created_by, updated_at, updated_by"
	} else if name == "user_credential" {
		return "customers_slug, email, password, customers_image"
	} else if name == "user_access" {
		return "context_type, context_id"
	} else if name == "auth" {
		return *firstTable + "s_slug, password"
	}
	return ""
}

func GetTemplateConcat(name, col string) string {
	if name == "value_group" {
		return "GROUP_CONCAT(" + col + " SEPARATOR ', ') as " + col
	}
	return ""
}

func GetTemplateOrder(name, tableName, ext string) string {
	if name == "permanent_data" {
		return tableName + ".created_at DESC, " + tableName + "." + ext + " DESC "
	} else if name == "dynamic_data" {
		return tableName + ".updated_at DESC, " + tableName + ".created_at DESC, " + tableName + "." + ext + " DESC "
	} else if name == "most_used_normal" {
		return " COUNT(1) DESC"
	}
	return ""
}

func GetTemplateJoin(typeJoin, firstTable, firstCol, secondTable, secondCol string, nullable bool) string {
	var join = ""
	if nullable {
		join = "LEFT JOIN "
	} else {
		join = "JOIN "
	}
	if typeJoin == "same_col" {
		return join + secondTable + " USING(" + firstCol + ") "
	} else if typeJoin == "total" {
		return join + secondTable + " ON " + secondTable + "." + secondCol + " = " + firstTable + "." + firstCol + " "
	}
	return ""
}

func GetTemplateGroup(is_multi_where bool, col string) string {
	var ext = " WHERE "
	if is_multi_where {
		ext = " AND "
	}

	return ext + col + " IS NOT NULL AND trim(" + col + ") != '' GROUP BY " + col + " "
}

func GetTemplateLogic(name string, colTarget *string) string {
	if name == "active" {
		return ".deleted_at IS NULL "
	} else if name == "trash" {
		return ".deleted_at IS NOT NULL "
	} else if name == "multi_content" {
		return "CASE " +
			"WHEN catalog_type = 'animals' THEN animals_" + *colTarget + " " +
			"WHEN catalog_type = 'plants' THEN plants_" + *colTarget + " " +
			"WHEN catalog_type = 'goods' THEN goods_" + *colTarget + " " +
			"END AS catalog_" + *colTarget
	}
	return ""
}

func GetWhereMine(token string) string {
	return "users_tokens.token ='" + token + "'"
}

func GetTemplateCommand(name, tableName, colName string) string {
	if name == "soft_delete" {
		return "UPDATE " + tableName + " SET deleted_at = ?, deleted_by = ? WHERE " + tableName + "." + colName + " = ?"
	} else if name == "hard_delete" {
		return "DELETE FROM " + tableName + " WHERE " + tableName + "." + colName + " = ?"
	} else if name == "recover" {
		return "UPDATE " + tableName + " SET deleted_at = null, deleted_by = null WHERE " + tableName + "." + colName + " = ?"
	}
	return ""
}

// Stats
func GetTemplateStats(ctx, firstTable, name string, ord string, joinArgs *string) string {
	// Nullable args
	var args string
	if joinArgs == nil {
		args = ""
	} else {
		args = *joinArgs
	}
	// Notes :
	// Full query
	if name == "most_appear" {
		return "SELECT " + ctx + " as context, " + GetFormulaQuery(nil, "total_item") + " total FROM " + firstTable + " " + args + " GROUP BY " + ctx + " ORDER BY total " + ord
	}

	return ""
}

func GetFormulaQuery(colTarget *string, name string) string {
	if name == "average" {
		return "CEIL(SUM(" + *colTarget + ") / COUNT(1)) AS "
	} else if name == "total_item" {
		return "COUNT(1) AS "
	} else if name == "total_sum" {
		return "SUM(" + *colTarget + ") AS "
	} else if name == "total_condition" {
		// Column target with condition
		return "COUNT(CASE WHEN " + *colTarget + " THEN 1 END) AS "
	} else if name == "max" {
		return "MAX(" + *colTarget + ") AS "
	} else if name == "min" {
		return "MIN(" + *colTarget + ") AS "
	} else if name == "max_object" || name == "min_object" {
		prop, err := converter.StringToMap(*colTarget)
		if err != nil {
			fmt.Println("Error:", err)
			return ""
		}

		toCount, _ := prop["to_count"]
		toGet, _ := prop["to_get"]
		fromTable, _ := prop["from_table"]

		if name == "max_object" {
			return "(SELECT " + toGet + " FROM " + fromTable + " WHERE " + toCount + " = (SELECT MAX(" + toCount + ") FROM " + fromTable + ")) AS "
		} else if name == "min_object" {
			return "(SELECT " + toGet + " FROM " + fromTable + " WHERE " + toCount + " = (SELECT MIN(" + toCount + ") FROM " + fromTable + ")) AS "
		}
	} else if name == "hour" {
		return "CONCAT(EXTRACT(HOUR FROM " + *colTarget + "),':00') AS "
	}
	return ""
}
