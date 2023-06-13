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
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {

			// Count the number of neighboring cells
			count := 0
			if i != 0 {
				if cells[i-1][j] == 1 {
					count++
				}
			}
			if i != size-1 {
				if cells[i+1][j] == 1 {
					count++
				}
			}
			if j != 0 {
				if cells[i][j-1] == 1 {
					count++
				}
			}
			if j != size-1 {
				if cells[i][j+1] == 1 {
					count++
				}
			}
			if i != 0 && j != 0 {
				if cells[i-1][j-1] == 1 {
					count++
				}
			}
			if i != size-1 && j != size-1 {
				if cells[i+1][j+1] == 1 {
					count++
				}
			}
			if i != 0 && j != size-1 {
				if cells[i-1][j+1] == 1 {
					count++
				}
			}
			if i != size-1 && j != 0 {
				if cells[i+1][j-1] == 1 {
					count++
				}
			}
			if count < 2 && cells[i][j] == 1 {
				final[i][j] = 0
			}
			if (count >= 2 && count <= 3) && cells[i][j] == 1 {
				final[i][j] = 1
			}
			if count > 3 && cells[i][j] == 1 {
				final[i][j] = 0
			}
			if count == 3 && cells[i][j] == 0 {
				final[i][j] = 1
			}
		}
	}

	// Copy the generated frame to current frame for next Tick
	for i := range cells {
		copy(cells[i], final[i])
	}
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
	fps = 5 // Rate at which simulation runs
	EventLoop(fps)
}
