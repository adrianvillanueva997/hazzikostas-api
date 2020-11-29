package v1

import (
	"database/sql"
	"log"
)
// Private function that checks if certain character exists in the DB, returns true or false.
func checkCharacter(name string, cursor *sql.DB) (*bool, error) {
	result, err := cursor.Query("Select Toon_Name from Discord_Bots.HK_Toons_1 where Toon_Name = ?", name)
	status := false
	if err != nil {
		log.Println(err)
		return &status, err
	}
	if result.Next() {
		status = true
		return &status, nil
	}
	return &status, nil
}