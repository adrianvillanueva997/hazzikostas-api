package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	v1 "hazzikostas-api/controllers/v1"
	"hazzikostas-api/middleware/auth"
	"hazzikostas-api/pkg/routine"
	"log"
)

//nolint:funlen
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	go routine.LoadCronRoutines()
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))
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
		character := context.Query("character")
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
	log.Println("Server running!")
	err = router.Run("0.0.0.0:5000")
	if err != nil {
		log.Println(err)
	}
}
