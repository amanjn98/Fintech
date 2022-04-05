package handler

import (
	db2 "Fintech/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var ErrPasswordCheck = fmt.Errorf("Password Check failed")
var ErrInternalServer= fmt.Errorf("Internal Server Error")
var ErrInvalidEmail=fmt.Errorf("Invalid email")
func LogIn(w http.ResponseWriter, r *http.Request){
	user := &User{}
	errorJson :=&ErrorResponse{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !isEmailValid(user.Username){
		errorJson.Message="Please enter the correct emailid"
		errorJson.Err=ErrInvalidEmail
		SendJsonResponse(w, errorJson)
		return
	}
	// validate the password against our requirements
	passwordCheck,errorString:=PasswordChecker(user.Password,user.Username)
	if !passwordCheck{
		errorJson.Message=errorString
		errorJson.Err=ErrPasswordCheck
		SendJsonResponse(w, errorJson)
		return
	}
	errCode,password:=db2.GetUsersPassword(user.Username)
	if errCode!=200{
		errorJson.Message="Unable to retrieve user details"
		w.WriteHeader(errCode)
		SendJsonResponse(w, errorJson)
		return
	}
	// Check the password with the hash password
	if !CheckPasswordHash(user.Password,password) {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	SendJsonResponse(w,map[string]interface{}{"message":"Successfully LoggedIn"})
}

func SignUp(w http.ResponseWriter, r *http.Request){
   user:=&User{}
   err:=json.NewDecoder(r.Body).Decode(user)
   if err!=nil{
	   //  Return a 400 status if there is something wrong with the request body
	   w.WriteHeader(http.StatusBadRequest)
	   return
   }

   errorJson :=&ErrorResponse{}
	if !isEmailValid(user.Username){
		errorJson.Message="Please enter the correct emailid"
		errorJson.Err=ErrInvalidEmail
		SendJsonResponse(w, errorJson)
		return
	}
   // validate the password against our requirements
	// validate the password against our requirements
	passwordCheck,errorString:=PasswordChecker(user.Password,user.Username)
	if !passwordCheck{
		errorJson.Message=errorString
		errorJson.Err=ErrPasswordCheck
		SendJsonResponse(w, errorJson)
		return
	}
   hash,err:= HashPassword(user.Password)
   if err!=nil{
	   errorJson.Message="Try again after some time"
	   errorJson.Err=ErrInternalServer
	   w.WriteHeader(http.StatusInternalServerError)
	   return
   }

	if err = db2.InsertUsers(user.Username,string(hash)); err != nil {
		// If there is any issue with inserting into the database, return a 500 errorJson
		w.WriteHeader(http.StatusInternalServerError)
		errorJson.Message="Try again after some time"
		errorJson.Err=ErrInternalServer
		SendJsonResponse(w,errorJson)
		return
	}
	SendJsonResponse(w,map[string]interface{}{"message":"Successfully Created account"})
}

func ResetPassword(w http.ResponseWriter, r *http.Request){
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// validate the password against our requirements
	errorJson :=&ErrorResponse{}
	if !isEmailValid(user.Username){
		errorJson.Message="Please enter the correct emailid"
		errorJson.Err=ErrInvalidEmail
		SendJsonResponse(w, errorJson)
		return
	}
	// validate the password against our requirements
	passwordCheck,errorString:=PasswordChecker(user.Password,user.Username)
	if !passwordCheck{
		errorJson.Message=errorString
		errorJson.Err=ErrPasswordCheck
		SendJsonResponse(w, errorJson)
		return
	}
	errCode,password:=db2.GetUsersPassword(user.Username)
	if errCode!=200{
		errorJson.Message="Unable to retrieve user details"
		w.WriteHeader(errCode)
		SendJsonResponse(w, errorJson)
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	if CheckPasswordHash(user.Password, password){
		// If the two passwords match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		SendJsonResponse(w, map[string]interface{}{"message":"Please don't use the old password"})
		return
	}
	hash,err:= HashPassword(user.Password)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Next, insert the username, along with the hashed password into the database
	if err = db2.UpdateUsers(user.Username,string(hash)); err != nil {
		// If there is any issue with inserting into the database, return a 500 errorJson
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	SendJsonResponse(w,map[string]interface{}{"message":"Successfully changed Password"})

}

func SendJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
