package routine

import (
	"github.com/go-co-op/gocron"
	v1 "hazzikostas-api/controllers/v1"
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
	cron1.Every(1).Day().At("15:55")
	_, err := cron1.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron2 := gocron.NewScheduler(time.UTC)
	cron2.Every(1).Day().At("21:55")
	_, err = cron2.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron3 := gocron.NewScheduler(time.UTC)
	cron3.Every(1).Day().At("5:55")
	_, err = cron3.Do(ExecuteRoutine)
	if err != nil {
		log.Println(err)
	}
	cron4 := gocron.NewScheduler(time.UTC)
	cron4.Every(1).Day().At("23:55")
	_, err4 := cron4.Do(ExecuteRoutine)
	if err4 != nil {
		log.Println(err)
	}
	cron1.StartAsync()
	cron2.StartAsync()
	cron3.StartAsync()
	cron4.StartAsync()
	log.Println("Cron routines started successfully")
}
