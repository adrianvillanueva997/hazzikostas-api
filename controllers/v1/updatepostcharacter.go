package v1

import (
	"hazzikostas-api/pkg/db"
	"log"
)

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
	results, err := cursor.Query("UPDATE HK_Toons_1 set Post=0 WHERE Toon_Name = ?", characterName)
	if err != nil {
		log.Println(err)
		return err
	}
	if results.Next() {
		log.Println(results)
	}
	return nil
}
