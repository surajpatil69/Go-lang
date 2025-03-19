package helper

import "strings"

func IsvalidUseInput(username string, surname string, email string) (bool, bool) {
	// validation of strings
	var isvalidname bool = len(username) >= 2 && len(surname) >= 2
	isvlaliedemail := strings.Contains(email, "@")

	return isvalidname, isvlaliedemail
}
