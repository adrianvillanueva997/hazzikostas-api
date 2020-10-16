package routes

import (
	"github.com/gin-gonic/gin"
	v1 "hazzikostas-api/controllers/v1"
	"hazzikostas-api/middleware/auth"
	"log"
)

//nolint:funlen
func GetCharacterRoutes(router *gin.Engine) {
	router.GET("/api/v1/characters", func(context *gin.Context) {
		characters, err := v1.GetCharacters()
		if err != nil {
			log.Println(err)
		}
		context.JSON(200, characters)
	})
	router.GET("/api/v1/postcharacters", func(context *gin.Context) {
		characters, err := v1.GetCharactersToPost()
		if err != nil {
			log.Println(err)
		}
		context.JSON(200, characters)
	})
	router.GET("/api/v1/updatecharacter", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		character := context.Query("character_r")
		status, err := auth.AuthenticateUser(username, password)
		if err != nil {
			log.Println(err)
			context.JSON(401, nil)
		}
		if *status {
			err := v1.UpdatePostCharacterStatus(character)
			if err != nil {
				log.Println(err)
				context.JSON(401, nil)
			}
			context.JSON(200, "Ok")
		} else {
			context.JSON(401, nil)
		}
	})
	router.GET("/api/v1/createcharacter", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		characterName := context.Query("name")
		region := context.Query("region")
		realm := context.Query("realm")
		status, err := auth.AuthenticateUser(username, password)
		if err != nil {
			log.Println(err)
			context.JSON(401, nil)
		}
		if *status {
			err := v1.CreateCharacter(characterName, region, realm)
			if err != nil {
				log.Println(err)
				context.JSON(401, nil)
			}
			context.JSON(200, "Ok")
		}
		context.JSON(401, nil)
	})
	router.GET("/api/v1/deletecharacter", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		characterName := context.Query("name")
		status, err := auth.AuthenticateUser(username, password)
		if err != nil {
			log.Println(err)
			context.JSON(401, nil)
		}
		if *status {
			err := v1.DeleteCharacter(characterName)
			if err != nil {
				log.Println(err)
				context.JSON(401, nil)
			}
			context.JSON(200, "Ok")
		}
		context.JSON(401, nil)
	})

}
