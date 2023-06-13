package main

import (
	"time"
)

// Global Variables, no need to interfere
var cells [][]int
var final [][]int

// User defined parameters
var size int
var fps int

func init() {
	size = 20 // Size of frame

	cells = make([][]int, size)
	final = make([][]int, size)
	for i := range cells {
		cells[i] = make([]int, size)
		final[i] = make([]int, size)
	}

	// The Glider Pattern (https://conwaylife.com/wiki/Glider)
	cells[5][5] = 1
	cells[5][6] = 1
	cells[5][4] = 1
	cells[4][6] = 1
	cells[3][5] = 1
}

func Tick() {

}

func EventLoop(fps int) {

	// Runs an Event Loop that triggers a Tick at the specified fps
	loop := time.NewTicker(time.Duration(1000/fps) * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-loop.C:
				Tick() // No goroutines are used in Ticks in order to prevent data mismatch across ticks
			case <-quit:
				loop.Stop()
				return
			}
		}
	}()
	for {
	}
}
func main() {

}
