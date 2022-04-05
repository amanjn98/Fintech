package handler

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"unicode"
)

// PasswordChecker function will validate the password against the defined rules
// min 1 lowercase and 1 uppercase
// min 1 number
// min 1 special character
// min 8 char long
// No empty string or whitespace.
func PasswordChecker(pass string, user string) (bool, string) {
	if strings.EqualFold(pass,user){
		return false, ""
	}
	var upp,low,num,sym bool
	var tot int
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false, ""
		}
	}
	if !upp{
		return false,"missing upper case letter"

	}
	if !low{
		return false,"missing lower case letter"
	}
	if !num{
		return false,"missing number"
	}
	if !sym{
		return false,"missing symbol"
	}
	if tot<8{
		return false,"too short password"
	}
	return true, ""
}

// HashPassword function salt and hash the password using the bcrypt algorithm
// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power)
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

//CheckPasswordHash function check the password with the hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
