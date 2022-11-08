package ui

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/apoorvam/goterminal"
	"github.com/briandowns/spinner"
)

type MultiSpinner struct {
	mu          *sync.Mutex
	Delay       time.Duration      // Delay is the speed of the indicator
	chars       []string           // chars holds the chosen character set
	Data        string             // Suffix is the text appended to the indicator
	FinalMSG    string             // string displayed after Stop() is called
	Writer      *goterminal.Writer // to make testing better, exported so users have access. Use `WithWriter` to update after initialization.
	active      bool               // active holds the state of the spinner
	stopChan    chan struct{}
	spinnerType int // stopChan is a channel used to stop the indicator 	// hideCursor determines if the cursor is visible 	// will be triggered after every spinner update
}

func New(cs []string, d time.Duration, writer *goterminal.Writer, sType int) *MultiSpinner {
	s := &MultiSpinner{
		Delay:       d,
		chars:       cs,
		mu:          &sync.Mutex{},
		Writer:      writer,
		stopChan:    make(chan struct{}, 1),
		active:      false,
		spinnerType: sType,
	}

	return s
}

func (s *MultiSpinner) Start() {
	s.mu.Lock()
	s.Data = strings.Replace(s.Data, MULTISPINNER, "", -1)
	var arr []string
	_ = json.Unmarshal([]byte(s.Data), &arr)
	s.Data = strings.Join(arr, "\n\n")
	spinnerChars := spinner.CharSets[s.spinnerType]
	if s.active {
		s.mu.Unlock()
		return
	}
	s.active = true
	s.mu.Unlock()
	go func() {
		for {
			for i := 0; i < 10; i++ {

				select {
				case <-s.stopChan:
					//s.Writer.Clear()
					return
				default:
					s.mu.Lock()
					if !s.active {
						s.mu.Unlock()
						return
					}
					s.Writer.Clear()

					output := strings.Replace(s.Data, SPINNER, spinnerChars[i], -1) + "\n"

					fmt.Fprint(s.Writer, output)
					//fmt.Println(spinnerChars[i])
					s.Writer.Print()
					s.mu.Unlock()
					time.Sleep(s.Delay)
				}
			}
		}
	}()
}

func (s *MultiSpinner) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.active {
		s.active = false
		s.stopChan <- struct{}{}
		s.Writer.Clear()
	}
	s.Writer.Reset()
}
