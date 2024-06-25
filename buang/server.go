package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/judwhite/go-svc"
)

type Service struct {
	wq   sync.WaitGroup
	quit chan struct{}
}

func (s *Service) Init(env svc.Environment) error {
	log.Printf("is win service? %v\n", env.IsWindowsService())
	return nil
}

func (s *Service) Start() error {

	go func() {
		log.Println("starting service...")
		s.quit = make(chan struct{})
		s.wq.Add(1)
		s.StartHttpService()
	}()

	return nil
}

func (s *Service) Stop() error {
	log.Println("stopping service...")

	close(s.quit)
	s.wq.Wait()
	log.Println("Stopped")
	return nil
}

func (s *Service) StartHttpService() error {
	defer s.wq.Done()

	// buat service http yang simple
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title>GS</title>
			</head>
			<body>
				golang service home<br>
				<a href="/about">About</a>
			</body>
		</html>
		`)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title>GS - About</title>
			</head>
			<body>
				About Page<br>
				back to <a href="/">Home</a>
			</body>
		</html>			
		`)
	})

	return http.ListenAndServe(":8081", nil)

}
