package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-co-op/gocron/v2"
)

const (
	httpPort = ":3000"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("привет как дела?"))
		if err != nil {
			log.Println(err)
		}
	})

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	jobs, err := initJobs(scheduler)
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		fmt.Printf("starting server on port %s\n", httpPort)
		if err := http.ListenAndServe(httpPort, r); err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()

		fmt.Println("starting jobs:", jobs[0].ID())
		scheduler.Start()
	}()

	wg.Wait()
}
func initJobs(scheduler gocron.Scheduler) ([]gocron.Job, error) {

	j, err := scheduler.NewJob(
		gocron.DurationJob(1*time.Second),
		gocron.NewTask(func() {
			fmt.Println("hello world")
		},
		),
	)
	if err != nil {
		return nil, err
	}

	return []gocron.Job{j}, nil
}
