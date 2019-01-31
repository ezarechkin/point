package main

import   "fmt"
import   "log"
import   "errors"
import   "net/url"
import   "net/http"
import   "database/sql"
import _ "github.com/lib/pq"


func authenticateUser(message *string) (string, error) {

	var values	url.Values
	var output	string
	var err		error

	// Query example
	// phone=89619403010&password=PASSWORD_HASH
	if values, err = url.ParseQuery(*message); err == nil {

		// 
		var phone		string = values.Get("phone")
		var password	string = values.Get("password")
		var ip			string = values.Get("ip")
		
		// Check incoming data for length
		if len(phone) == 0 || len(password) == 0 || len(ip) == 0 {

			err = errors.New("VALIDATE Query: Empty required parameters")
			output = fmt.Sprintf("status=%d&message=%s", http.StatusBadRequest, "empty required parameters")
			return output, err
		} else {
			
			var db *sql.DB

			// Prepare instance auto destruction after use
			defer func(){

				log.Printf("CLOSE Database: phone: %s", "89619403010")
				db.Close()
			}()

			// Open database connection
			if db, err = sql.Open("postgres", DB_CONNECTION_STRING); err == nil {
				
				var id int
				
				// Get user with recived parameters	
				if err = db.QueryRow("SELECT id FROM users WHERE phone=$1 AND password=$2 LIMIT 1", phone, password).Scan(&id); err == nil {

					log.Printf("SELECT Database: phone: %s", phone)
					output = fmt.Sprintf("status=%d&id=%d", http.StatusOK, id)
					return output, nil
				} else {
					
					err = errors.New(fmt.Sprintf("SELECT Database: %s", err))
					output = fmt.Sprintf("status=%d&message=%s", http.StatusUnauthorized, "unable to find pair")
					return output, err
				}
			} else {

				err = errors.New(fmt.Sprintf("SELECT Database: %s", err))
				output = fmt.Sprintf("status=%d&message=%s", http.StatusInternalServerError, "internal server error")
				return output, err
			}
		}
	} else {

		err = errors.New(fmt.Sprintf("PARSE Query: ", err))
		output = fmt.Sprintf("status=%d&message=%s", http.StatusBadRequest, "invalid required parameters")		
		return output, err
	}
}