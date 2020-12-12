package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const(
	N = "N"
	S = "S"
	E = "E"
	W = "W"
	R = "R"
	L = "L"
	F = "F"
)

type Move struct {
	dir string
	steps int
}

func main() {
	f, err := os.Open("11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var moves []Move
	for scanner.Scan() {
		text := scanner.Text()
		dir := string(text[0])
		steps, err := strconv.Atoi(text[1:])
		if err != nil {
			log.Panic(err)
		}

		moves = append(moves, Move{
			dir: dir,
			steps: steps,
		})
	}

	solution := solvePartOne(moves)
	fmt.Printf("Part One -- The solution is %d\n", solution)


/*	solution2 := solvePartTwo(lay)
	fmt.Printf("Part Two -- The solution is %d\n", solution2)*/
}

func forward(dir, steps int) (int, int) {
}

func solvePartOne(moves []Move) int {
	x := 0
	y := 0
	dir := 1 // 0 = N, 1 = E...
	for _, move := range moves {
		switch move.dir {
		case N:
			y += move.steps
			break
		case S:
			y -= move.steps
			break
		case E:
			x += move.steps
			break
		case W:
			x -= move.steps
			break
		case R:
			dir = (dir + 1) % 4
			break
		case L:
			dir = (dir - 1) % 4
			break
		case F:
			forward(dir, )
			break
		default:
			log.Panic(fmt.Sprintf("Invalid direction dir: '%s'", move.dir))
		}
	}
}