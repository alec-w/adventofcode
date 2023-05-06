package internal

import (
	"strconv"
	"strings"

	"github.com/alec-w/adventofcode/2022/helpers"
)

type Position struct {
	X int
	Y int
}

type Knot struct {
	Position
	Next *Knot
}

func (k *Knot) Move(preceding *Knot) {
	defer func() {
		if k.Next != nil {
			k.Next.Move(k)
		}
	}()
	xDiff := preceding.X - k.X
	yDiff := preceding.Y - k.Y
	if xDiff >= -1 && xDiff <= 1 && yDiff >= -1 && yDiff <= 1 {
		return
	}
	if preceding.X == k.X {
		if k.Y < preceding.Y {
			k.Y++
		} else {
			k.Y--
		}
		return
	}
	if preceding.Y == k.Y {
		if k.X < preceding.X {
			k.X++
		} else {
			k.X--
		}
		return
	}
	if preceding.X > k.X {
		k.X++
	} else {
		k.X--
	}
	if preceding.Y > k.Y {
		k.Y++
	} else {
		k.Y--
	}
}

type Rope struct {
	Head Knot
}

func NewRope(length int) *Rope {
	rope := &Rope{Head: Knot{}}
	tail := &rope.Head
	for i := 0; i < length-1; i++ {
		tail.Next = &Knot{}
		tail = tail.Next
	}
	return rope
}

func (r *Rope) MoveUp() {
	r.Head.Y++
	if r.Head.Next != nil {
		r.Head.Next.Move(&r.Head)
	}
}
func (r *Rope) MoveDown() {
	r.Head.Y--
	if r.Head.Next != nil {
		r.Head.Next.Move(&r.Head)
	}
}
func (r *Rope) MoveLeft() {
	r.Head.X--
	if r.Head.Next != nil {
		r.Head.Next.Move(&r.Head)
	}
}
func (r *Rope) MoveRight() {
	r.Head.X++
	if r.Head.Next != nil {
		r.Head.Next.Move(&r.Head)
	}
}

func (r *Rope) Tail() *Knot {
	knot := &r.Head
	for ; knot.Next != nil; knot = knot.Next {
	}
	return knot
}

type Motion struct {
	Direction string
	Amount    int
}

func MotionsFromFile(fileName string) ([]Motion, error) {
	lines, err := helpers.FileToStrings(fileName)
	if err != nil {
		return nil, err
	}

	motions := []Motion{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		motions = append(motions, Motion{Direction: direction, Amount: amount})
	}

	return motions, nil
}
