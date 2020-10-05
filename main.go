package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

package main

import (
"github.com/gin-gonic/gin"
"github.com/joho/godotenv"
v1 "ion-api/routes/api/v1"
"log"
)
type Message struct {
	Message string `json:"message"`
}

func main() {
	//routine.GetRaiderData("us", "wyrmrest-accord", "uwupolicia")
	// routine.GetRaiderData()
	err := godotenv.Load() //nolint:ineffassign
	if err != nil {
		log.Fatalln(err)
	}
	msg := Message{Message: "pong"}
	router := gin.Default()
	router.GET("/api/v1/characters", func(context *gin.Context) {
		characters, err := v1.GetCharacters()
		if err != nil {
			log.Fatalln(err)
		}
		context.JSON(200, characters)
	})
	router.GET("/api/v1/updatecharacter", func(context *gin.Context) {
		context.JSON(200, msg)
	})
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, msg)
	})
	log.Println("Server running!")
	err = router.Run("0.0.0.0:3000")
	if err != nil {
		log.Fatalln(err)
	}

}
