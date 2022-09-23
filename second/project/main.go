package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

//! constructor
func NewWorker(id int, wp chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		WorkerPool: wp,
		JobQueue:   make(chan Job),
		QuitChan:   make(chan bool),
	}
}

func NewDispatcher(jqueue chan Job, maxworker int) *Dispatcher {
	worker := make(chan chan Job, maxworker)
	return &Dispatcher{
		JobQueue:   jqueue,
		MaxWorkers: maxworker,
		WorkerPool: worker,
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id: %d started", w.Id)
				fib := fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worked end, result %d", fib)

			case <-w.QuitChan:
				fmt.Printf("Worker %d stopped", w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()

		}
	}
}

func (d *Dispatcher) RunDispatch() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func RequestHanler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	job := Job{
		Name: name, Delay: delay, Number: value,
	}

	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.RunDispatch()

	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHanler(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
