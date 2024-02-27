package validations

import (
	"pet-store/modules/auth/models"
	"pet-store/packages/helpers/converter"
	"pet-store/packages/helpers/generator"
	"pet-store/packages/utils/validator"
)

func GetValidateRegister(body models.UserRegister) (bool, string) {
	var msg = ""
	var status = true

	// Rules
	minUname, maxUname := validator.GetValidationLength("customers_slug")
	minPass, maxPass := validator.GetValidationLength("password")
	minEmail, maxEmail := validator.GetValidationLength("email")
	minCName, maxCName := validator.GetValidationLength("customers_name")

	// Value
	uname := converter.TotalChar(body.Username)
	pass := converter.TotalChar(body.Password)
	email := converter.TotalChar(body.Email)
	cname := converter.TotalChar(body.CustomerName)

	// Validate
	if uname <= minUname || uname >= maxUname {
		status = false
		msg += generator.GenerateValidatorMsg("Username", minUname, maxUname)
	}
	if pass <= minPass || pass >= maxPass {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Password", minPass, maxPass)
	}
	if email <= minEmail || email >= maxEmail {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Email", minEmail, maxEmail)
	}
	if cname <= minCName || cname >= maxCName {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("First name", minCName, maxCName)
	}

	if status {
		return status, "Validation success"
	} else {
		return status, msg
	}
}

func GetValidateLogin(username, password, role string) (bool, string) {
	var msg = ""
	var status = true

	// Rules
	minUname, maxUname := validator.GetValidationLength("customers_slug")
	minPass, maxPass := validator.GetValidationLength("password")

	// Value
	uname := converter.TotalChar(username)
	pass := converter.TotalChar(password)

	// Validate
	if uname <= minUname || uname >= maxUname {
		status = false
		msg += generator.GenerateValidatorMsg("Username", minUname, maxUname)
	}
	if pass <= minPass || pass >= maxPass {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Password", minPass, maxPass)
	}
	if role != "admin" && role != "staff" && role != "customer" && role != "doctor" {
		status = false
		msg += "Your role is invalid"
	}

	if status {
		return status, "Validation success"
	} else {
		return status, msg
	}
}
