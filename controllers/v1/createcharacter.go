package v1

import (
	"hazzikostas-api/pkg/db"
	"log"
)

func CreateCharacter(name string, region string, realm string) (*bool, error) { //nolint:whitespace
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
	if !*status {
		_, err := cursor.Query("insert into HK_Toons_1 (toon_name, region, realm)VALUES (?,?,?)", name, region, realm)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return status, nil
}
