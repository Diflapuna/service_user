package main

import "github.com/NotYourAverageFuckingMisery/animello/internal/service"

func main() {
	s := service.NewService()
	s.Log.Fatal(s.Start())
}
