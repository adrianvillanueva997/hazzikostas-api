package v1

import (
	"hazzikostas-api/pkg/db"
	"log"
)
// Updates a character in the DB if it was already posted by the discord bot.
func UpdatePostCharacterStatus(characterName string) error {
	cursor, err := db.SetConnection()
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		err := cursor.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	_, err = cursor.Exec("UPDATE HK_Toons_1 set Post=0 WHERE Toon_Name=?", characterName)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
