package v1

import (
	"database/sql"
	"hazzikostas-api/pkg/db"
	"log"
)

func CreateCharacter(character string, region string, realm string, serverID string) (*bool, error) { //nolint:whitespace
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
	status, err := checkCharacter(character, cursor)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !*status {
		_, err := cursor.Query("insert into HK_Toons_1 (toon_name, region, realm)VALUES (?,?,?)", character, region, realm)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	err = CheckServer(serverID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	characterStatus, err := checkCharacterServer(character, serverID, cursor)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !*characterStatus {
		*status = true
		err = insertCharacterRelation(character, serverID, cursor)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return characterStatus, nil
}

func checkCharacterServer(character string, serverID string, cursor *sql.DB) (*bool, error) {
	result, err := cursor.Query("Select PK_Character from Discord_Bots.HK_Toons_Relation where PK_Character = ? and ServerID = ?", character, serverID)
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

func insertCharacterRelation(character string, serverID string, cursor *sql.DB) error {
	_, err := cursor.Query("insert into Discord_Bots.HK_Toons_Relation (PK_Character, ServerID) VALUES (?,?)", character, serverID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
