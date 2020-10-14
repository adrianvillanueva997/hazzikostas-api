package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type message struct {
	Message string `json:"message"`
}

func buildURL(username string, password string) string {
	return fmt.Sprintf("https://auth.thexiao77.xyz/api/auth?username=%s&password=%s",
		username, password)
}

func AuthenticateUser(username string, password string) (*bool, error) {
	var resp, err = http.Get(buildURL(username, password))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	data := message{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	log.Println(data.Message)
	if data.Message == "Ok" {
		status := true
		return &status, nil
	}
	status := false
	return &status, nil
}
