package v1

import (
	"database/sql"
	"hazzikostas-api/pkg/db"
	"log"
)

// Checks the server where the bot is online now and deletes from the DB where he's not online
func CheckServer(serverID string) error {
	cursor, err := db.SetConnection()
	defer func() {
		err := cursor.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	status := false
	if err != nil {
		log.Println(err)
		return err
	}
	result, err := cursor.Query("SELECT PK_ServerID FROM Discord_Bots.HK_Servers WHERE PK_ServerID = ?", serverID)
	if err != nil {
		log.Println(err)
		return err
	}
	if result.Next() {
		status = true
	}
	log.Println(status)
	if !status {
		err := addServer(serverID, cursor)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// Includes the server ID in the database
func addServer(serverID string, cursor *sql.DB) error {
	_, err := cursor.Query("INSERT INTO Discord_Bots.HK_Servers (PK_ServerID) VALUES (?)", serverID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

/*
// Function not used yet
func deleteServer(serverID string, cursor *sql.DB) error {
	_, err := cursor.Query("DELETE FROM Discord_Bots.HK_Affixes_Channels where PK_ServerID = ?", serverID)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = cursor.Query("DELETE FROM Discord_Bots.HK_Toons_Relation where PK_ServerID = ?", serverID)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = cursor.Query("DELETE FROM Discord_Bots.HK_Servers where PK_ServerID = ?", serverID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

*/
