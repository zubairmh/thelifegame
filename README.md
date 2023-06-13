# Conway's Game of Life

This is a simple implementation of Conway's Game of Life written in Go. The Game of Life is a cellular automaton devised by the mathematician John Conway, where cells in a grid evolve based on a set of rules. It is a zero-player game, meaning its evolution is determined by its initial state, requiring no further input.

### Rules

The rules of Conway's Game of Life are as follows:

	1) Any live cell with fewer than two live neighbours dies.
	2) Any live cell with more than three live neighbours dies.
	3) Any live cell with two or three live neighbours lives, unchanged, to the next generation.
	4) Any dead cell with exactly three live neighbours will come to life.


### Features

  -  Adjustable grid size: You can set the size of the grid to accommodate different patterns. ✅
  -  Configurable simulation speed: You can control the speed at which the generations evolve. ✅
  -  ASCII representation: The game state is displayed using ASCII characters on the command line. ✅

### Installation

To run the Conway's Game of Life implementation, make sure you have Go installed on your system. Then, follow these steps:

Clone this repository to your local machine.

    git clone https://github.com/zubairmh/thelifegame
    cd thelifegame

Build the executable.

`go build`

### Usage

To start the game, run the following command:

`./gameoflife`

This will start the simulation with default settings: a grid size of 20x20, a glider pattern and a simulation speed of 5 ticks per second.

### License

This project is licensed under the [MIT License](https://github.com/zubairmh/thelifegame/blob/main/LICENSE).
Acknowledgements

- John Conway for creating the Game of Life and introducing it to the world.
- The Go programming language community for providing a powerful and efficient language for building this implementation.

Feel free to contribute to this project by submitting issues or pull requests. Enjoy the Game of Life!
