package routine

import (
	"github.com/go-co-op/gocron"
	v1 "hazzikostas-api/routes/api/v1"
	"log"
	"time"
)

func ExecuteRoutine() {
	log.Println("Starting routine")
	characters, err := v1.GetCharacters()
	if err != nil {
		log.Println(err)
	}
	Routine(*characters)
}

func LoadCronRoutines() {
	cron1 := gocron.NewScheduler(time.UTC)
	cron1.Every(1).Day().At("16:00")
	_, err := cron1.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron2 := gocron.NewScheduler(time.UTC)
	cron2.Every(1).Day().At("22:00")
	_, err = cron2.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron3 := gocron.NewScheduler(time.UTC)
	cron3.Every(1).Day().At("04:00")
	_, err = cron3.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron1.StartAsync()
	cron2.StartAsync()
	cron3.StartAsync()
	log.Println("Cron routines started successfully")
}