package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hazzikostas-api/pkg/routine"
	"hazzikostas-api/routes"
	"log"
)

//nolint:funlen
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	go routine.LoadCronRoutines()
	go routine.ExecuteRoutine()
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	routes.GetCharacterRoutes(router)
	log.Println("Server running!")
	err = router.Run("0.0.0.0:5000")
	if err != nil {
		log.Println(err)
	}
}
