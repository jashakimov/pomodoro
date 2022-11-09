package main

import (
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"strconv"
	"time"
)

func main() {
	min := flag.Int("time", 25, "Alerting time")
	flag.Parse()
	if err := beeep.Alert("Pomodoro", fmt.Sprint("Запущена задача на ", *min, " минут"), ""); err != nil {
		panic(err)
	}
	time.Sleep(time.Duration(*min) * time.Minute)
	if err := beeep.Alert("Pomodoro", strconv.Itoa(*min)+" минут прошли", ""); err != nil {
		panic(err)
	}
}
