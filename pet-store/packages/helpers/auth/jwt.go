package auth

import (
	"pet-store/configs"
	"pet-store/packages/builders"
	"strings"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(pass string) string {
	hash, _ := HashPassword(pass)

	return hash
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetJWTConfiguration(name string) string {
	if name == "exp" {
		conf := configs.GetConfigJWT()
		return conf.JWT_EXP
	}
	return ""
}

func GetTokenHeader(c echo.Context) (bool, string) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return false, "No authorization header present"
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return false, "Invalid authorization header format"
	}

	token := authHeader[len(bearerPrefix):]
	return true, token
}

func GetAuthQuery(baseTable, token string) string {
	token = strings.Replace(token, "Bearer ", "", -1)
	var prop = "created_by"

	// Query builder
	if baseTable == "customers" {
		prop = "id"
	}

	join := builders.GetTemplateJoin("total", baseTable, prop, "users_tokens", "context_id", false)

	sqlStatement :=
		join + " " +
			"WHERE token = '" + token + "' "

	return sqlStatement
}
