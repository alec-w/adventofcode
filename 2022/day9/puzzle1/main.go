package main

import (
	"fmt"
	"log"

	"github.com/alec-w/adventofcode/2022/day9/internal"
)

func main() {
	rope := internal.NewRope(2)
	tail := rope.Tail()
	motions, err := internal.MotionsFromFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	visited := map[internal.Position]struct{}{tail.Position: {}}

	move := func(direction func(), amount int) {
		for i := 0; i < amount; i++ {
			direction()
			visited[tail.Position] = struct{}{}
		}
	}
	for _, motion := range motions {
		switch motion.Direction {
		case "U":
			move(rope.MoveUp, motion.Amount)
		case "D":
			move(rope.MoveDown, motion.Amount)
		case "L":
			move(rope.MoveLeft, motion.Amount)
		case "R":
			move(rope.MoveRight, motion.Amount)
		}
	}

	fmt.Printf("positions visited: %d\n", len(visited))
}
