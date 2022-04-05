package db

import (
	"database/sql"
)


func InsertUsers(username string,password string) error{
	if _, err := db.Query("insert into users values ($1, $2)", username, password); err != nil {
		// If there is any issue with inserting into the database, return a 500 errorJson
		return err
	}
	return nil
}

func UpdateUsers(username string, password string) error {
	if _, err := db.Query("UPDATE users set password= $2 where username=$1", username,password); err != nil {
		// If there is any issue with inserting into the database, return a 500 errorJson
		return err
	}
	return nil
}


func GetUsersPassword(username string) (errorCode int,pass string){
	// Get the existing entry present in the database for the given username
	result := db.QueryRow("select password from users where username=$1", username)
	// We create another instance of `Users` to store the credentials we get from the database
	// Store the obtained password in `storedCreds`
	var password string
	err := result.Scan(&password)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			return 401,""
		}
		return 500,""
	}
	return 200,password
}