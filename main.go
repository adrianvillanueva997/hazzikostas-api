package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hazzikostas-api/pkg/routine"
	v1 "hazzikostas-api/routes/api/v1"
	"log"
)

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
		context.JSON(200, nil)
	})

	log.Println("Server running!")
	err = router.Run("0.0.0.0:5000")
	if err != nil {
		log.Println(err)
	}
}
