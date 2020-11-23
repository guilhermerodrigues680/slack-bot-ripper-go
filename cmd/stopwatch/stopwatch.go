package stopwatch

import (
	"fmt"
	"time"
)

// Status representa a situacao do cronometro
type Status struct {
	IsPaused    bool      `json:"isPaused"`
	StartTime   time.Time `json:"startTime,omitempty"`
	CurrentTime time.Time `json:"currentTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
}

// Stopwatch representa o cronometro
type Stopwatch struct {
	isPaused    bool
	startTime   time.Time
	currentTime time.Time
	endTime     time.Time
	ticker      *time.Ticker
}

// Start inicia a contagem do cronometro
func (s *Stopwatch) Start() {
	s.startTime = time.Now().UTC()
	s.currentTime = s.startTime
	s.endTime = s.startTime.Add(time.Hour)
	s.isPaused = false
	s.ticker.Reset(1000 * time.Millisecond)
	fmt.Println(s.startTime, s.currentTime, s.endTime)
}

// Pause pausa a contagem do cronometro
func (s *Stopwatch) Pause() {
	s.ticker.Stop()
	s.isPaused = true
}

// Reset redefine a contagem do cronometro
func (s *Stopwatch) Reset() {
	s.ticker.Stop()
}

// Status retorna um objeto contendo a situacao do cronometro
func (s Stopwatch) Status() Status {
	return Status{
		IsPaused:    s.isPaused,
		StartTime:   s.startTime,
		CurrentTime: s.currentTime,
		EndTime:     s.endTime,
	}
}

// Cronometro é uma Go Rotine que realiza a logica de contagem do cronometro
func (s *Stopwatch) Cronometro() {
	for {
		select {
		case t := <-s.ticker.C:
			s.currentTime = t.UTC()
			fmt.Printf("Ticker at: %s\n", t.UTC())
			fmt.Println(s.currentTime.Sub(s.startTime))
			fmt.Println(s.endTime)

			if s.currentTime.After(s.endTime) {
				s.ticker.Stop()
			}
		}
	}
}

// NewStopwatch retorna uma nova intancia de um cronometro
func NewStopwatch() *Stopwatch {
	s := Stopwatch{}
	s.isPaused = true
	s.ticker = time.NewTicker(1000 * time.Millisecond)
	s.ticker.Stop()
	return &s
}
