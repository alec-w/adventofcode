package internal

type Option = int

const (
	OPTION_ROCK Option = iota
	OPTION_PAPER
	OPTION_SCISSORS
)

type Round struct {
	Opponent Option
	You      Option
}

func (r Round) Score() int {
	return ((((((r.You - r.Opponent) + 2) % 3) + 2) % 3) * 3) + (r.You + 1)
}

/**
((((you - opponent) + 2) % 3) + 2) % 3
R R => D =>  0 => 2 => 2 => 4 => 1
R P => L => -1 => 1 => 1 => 3 => 0
R S => W => -2 => 0 => 0 => 2 => 2
P R => W =>  1 => 3 => 0 => 2 => 2
P P => D =>  0 => 2 => 2 => 4 => 1
P S => L => -1 => 1 => 1 => 3 => 0
S R => L =>  2 => 4 => 1 => 3 => 0
S P => W =>  1 => 3 => 0 => 2 => 2
S S => D =>  0 => 2 => 2 => 4 => 1
*/
