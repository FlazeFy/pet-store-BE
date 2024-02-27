package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"pet-store/modules/auth/models"
	"pet-store/modules/auth/validations"
	"pet-store/packages/builders"
	"pet-store/packages/database"
	"pet-store/packages/helpers/auth"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/helpers/response"
	"strconv"
	"strings"
)

func PostUserAuth(username, password, role string) (string, error, string) {
	status, msg := validations.GetValidateLogin(username, password, role)
	if status {
		// Declaration
		var obj models.UserLogin
		var pwd string
		var id string

		// Exec
		selectTemplate := builders.GetTemplateSelect("auth", &role, nil)
		baseTable := role + "s"
		sqlStatement := "SELECT id, " + selectTemplate + " " +
			"FROM " + baseTable +
			" WHERE " + role + "s_slug = ?"

		con := database.CreateCon()
		err := con.QueryRow(sqlStatement, username).Scan(
			&id, &obj.Username, &pwd,
		)

		if err == sql.ErrNoRows {
			return "", nil, "Account is not registered"
		} else if err != nil {
			return "", err, "Something went wrong. Please contact Admin"
		}

		match, err := auth.CheckPasswordHash(password, pwd)
		if !match {
			return "", nil, "Password incorrect"
		}

		if err != nil {
			return "", err, "Something went wrong. Please contact Admin"
		}

		return id, nil, ""
	} else {
		return "", nil, msg
	}
}

func PostUserRegister(body models.UserRegister) (response.Response, error) {
	var res response.Response
	status, msg := validations.GetValidateRegister(body)

	if status {
		// Declaration
		var baseTable = "customers"
		id, err := generator.GenerateUUID(16)
		if err != nil {
			return res, err
		}

		createdAt := generator.GenerateTimeNow("timestamp")
		slug := generator.GetSlug(body.Username)
		hashPass := auth.GenerateHashPassword(body.Password)

		// Query builder
		colFirstTemplate := builders.GetTemplateSelect("user_credential", nil, nil)

		if err != nil {
			return res, err
		}

		sqlStatement := "INSERT INTO " + baseTable + " " +
			"(id, " + colFirstTemplate + ", customers_name, customers_interest, is_notifable, created_at, created_by, updated_at, updated_by) " + " " +
			"VALUES (?, ?, ?, ?, null, ?, null, 1, ?, ?, null, null)"

		// Exec
		con := database.CreateCon()
		cmd, err := con.Prepare(sqlStatement)
		defer cmd.Close()

		if err != nil {
			return res, err
		}

		result, err := cmd.Exec(id, slug, body.Email, hashPass, body.CustomerName, createdAt, body.CreatedBy)
		if err != nil {
			return res, err
		}

		rowsAffected, _ := result.RowsAffected()
		resultStr := fmt.Sprintf("%d", rowsAffected)

		// Response
		res.Status = http.StatusOK
		res.Message = generator.GenerateCommandMsg("account", "register", 1)
		res.Data = map[string]string{"last_inserted_id": id, "result": resultStr + " rows affected"}
	} else {
		res.Status = http.StatusUnprocessableEntity
		res.Message = generator.GenerateCommandMsg("account. "+msg, "register", 0)
		res.Data = map[string]string{"result": "0 rows affected"}
	}
	return res, nil
}

func PostAccessToken(body models.UserToken) error {
	// Declaration
	var baseTable = "users_tokens"
	id, err := generator.GenerateUUID(16)
	if err != nil {
		return err
	}
	createdAt := generator.GenerateTimeNow("timestamp")

	// Query builder
	colFirstTemplate := builders.GetTemplateSelect("user_access", nil, nil)

	sqlStatement := "INSERT INTO " + baseTable + " " +
		"(id, " + colFirstTemplate + ", token, last_used_at, created_at) " + " " +
		"VALUES (?, ?, ?, ?, null, ?)"

	// Exec
	con := database.CreateCon()
	cmd, err := con.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	result, err := cmd.Exec(id, body.ContextType, body.ContextId, body.Token, createdAt)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}

func SignOut(token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "users_tokens"
	token = strings.Replace(token, "Bearer ", "", -1)

	sqlStatement := "DELETE FROM " + baseTable + " WHERE token= ?"

	// Exec
	con := database.CreateCon()
	cmd, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := cmd.Exec(token)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg("account", "sign out", 1)
	res.Data = map[string]string{"result": strconv.Itoa(int(rowsAffected)) + " rows affected"}

	return res, err
}

func CheckRole(token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "users_tokens"
	token = strings.Replace(token, "Bearer ", "", -1)
	var role string
	var sqlStatement string

	// Query builder
	sqlStatement = "SELECT context_type " +
		"FROM " + baseTable + " " +
		"WHERE token = '" + token + "'"

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
			&role,
		)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg("account", "check", 1)
	if role == "" {
		res.Data = nil
	} else {
		res.Data = role
	}

	return res, err
}
