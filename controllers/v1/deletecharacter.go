package v1

import (
	"hazzikostas-api/pkg/db"
	"log"
)

func DeleteCharacter(name string) (*bool, error) {
	cursor, err := db.SetConnection()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		err := cursor.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	status, err := checkCharacter(name, cursor)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if *status {
		result, err := cursor.Query("DELETE from HK_Toons_1 where toon_name = ?", name)
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return status, nil
		}
	}
	return status, nil
}
