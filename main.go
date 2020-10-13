package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	v1 "hazzikostas-api/routes/api/v1"
	"hazzikostas-api/src/routine"
	"log"
	"time"
)

// ExecuteRoutine
func ExecuteRoutine() {
	log.Println("Starting routine")
	characters, err := v1.GetCharacters()
	if err != nil {
		log.Println(err)
	}
	routine.Routine(*characters)
}

func loadCronRoutines() {
	cron1 := gocron.NewScheduler(time.UTC)
	cron1.Every(1).Day().At("18:00")
	_, err := cron1.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron2 := gocron.NewScheduler(time.UTC)
	cron2.Every(1).Day().At("00:00")
	_, err = cron2.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron3 := gocron.NewScheduler(time.UTC)
	cron3.Every(1).Day().At("06:00")
	_, err = cron3.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron1.StartAsync()
	cron2.StartAsync()
	cron3.StartAsync()
	log.Println("Cron routines started successfully")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	go loadCronRoutines()
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
