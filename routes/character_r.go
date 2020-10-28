package routes

import (
	"github.com/gin-gonic/gin"
	v1 "hazzikostas-api/controllers/v1"
	"hazzikostas-api/middleware/auth"
	"log"
)

//nolint:gocognit nolint:funlen
// Function to publish the API routes
func GetCharacterRoutes(router *gin.Engine) {
	// Get all characters
	router.GET("/api/v1/characters", func(context *gin.Context) {
		characters, err := v1.GetCharacters()
		if err != nil {
			log.Println(err)
		}
		context.JSON(200, characters)
	})
	// Get all characters ready to post in the bot
	router.GET("/api/v1/postcharacters", func(context *gin.Context) {
		characters, err := v1.GetCharactersToPost()
		if err != nil {
			log.Println(err)
		}
		context.JSON(200, characters)
	})
	// Update character, used when a character was already posted to put it on 0
	router.POST("/api/v1/updatecharacter", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		character := context.Query("character_r")
		status, err := auth.AuthenticateUser(username, password)
		if err != nil {
			log.Println(err)
			context.JSON(401, nil)
		}
		if !(username == "" && password == "" && character == "") { //nolint:nestif
			if *status {
				err := v1.UpdatePostCharacterStatus(character)
				if err != nil {
					log.Println(err)
					context.JSON(401, nil)
				}
				context.JSON(200, "Ok")
			} else {
				context.JSON(401, "User authentication failed")
			}
		} else {
			context.JSON(401, "Params missing")
		}
	})
	// Used to create a new character to track
	router.POST("/api/v1/createcharacter", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		characterName := context.Query("character")
		region := context.Query("region")
		realm := context.Query("realm")
		if !(username == "" && password == "" && characterName == "" && region == "" && realm == "") { //nolint:nestif
			status, err := auth.AuthenticateUser(username, password)
			if err != nil {
				log.Println(err)
				context.JSON(400, nil)
			}
			if *status {
				characterStatus, err := v1.CreateCharacter(characterName, region, realm)
				if err != nil {
					log.Println(err)
					context.JSON(500, nil)
				}
				if !*characterStatus && characterStatus != nil {
					context.JSON(201, "Ok")
				} else {
					context.JSON(204, "Character already exists")
				}
			} else {
				context.JSON(401, "User authentication failed")
			}
		} else {
			context.JSON(401, "Params missing")
		}
	})
	// Used to delete a character
	router.DELETE("/api/v1/deletecharacter", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		characterName := context.Query("character")
		if !(username == "" && password == "" && characterName == "") { //nolint:nestif
			status, err := auth.AuthenticateUser(username, password)
			if err != nil {
				log.Println(err)
				context.JSON(401, nil)
			}
			if *status {
				deleteStatus, err := v1.DeleteCharacter(characterName)
				if err != nil {
					log.Println(err)
					context.JSON(500, nil)

				}
				if *deleteStatus {
					context.JSON(200, "Ok")
				} else {
					context.JSON(204, "Character not found")
				}
			} else {
				context.JSON(401, "User authentication failed")
			}
		} else {
			context.JSON(401, "Params missing")
		}
	})
}
