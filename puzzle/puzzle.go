package puzzle

import (
	"fmt"
	"math/rand"
)

type Puzzle struct {
	pieces []uint
	width  uint
}

func (p *Puzzle) Len() int {
	return len(p.pieces)
}
func (p *Puzzle) Width() int {
	return int(p.width)
}
func (p *Puzzle) Height() int {
	return len(p.pieces) / int(p.width)
}
func (p *Puzzle) Pieces() []uint {
	return p.pieces
}
func (p *Puzzle) PiecesCopied() []uint {
	pieces := make([]uint, len(p.pieces))
	copy(pieces, p.pieces)
	return pieces
}

func (self *Puzzle) Print() {
	fmt.Printf("hi, i'm Puzzle{width=%v,%v}\n", self.width, self.pieces)
}

func Solved(width, height uint) Puzzle {
	length := width * height
	puzzle := Puzzle{make([]uint, length), width}
	pieces := puzzle.pieces[:]

	for i := range puzzle.pieces[:length-1] {
		pieces[i] = uint(i) + 1
	}

	return puzzle
}

func Gen(width, height uint, seed int64) Puzzle {
	length := width * height
	puzzle := Solved(width, height)
	pieces := puzzle.pieces[:]

	rng := rand.New(rand.NewSource(seed))
	rng.Shuffle(int(length-2), func(i, j int) {
		pieces[i], pieces[j] = pieces[j], pieces[i]
	})

	empty_idx := rng.Intn(int(length))
	empty_x, empty_y := empty_idx%int(width), empty_idx/int(width)

	d := (int(width) - 1 - empty_x) + (int(height) - 1 - empty_y)
	solvable := puzzle.Solvable()
	even := d%2 == 0
	switch {
	case solvable && d == 0:
	case solvable && even:
		pieces[empty_idx], pieces[length-2], pieces[length-1] = 0, pieces[empty_idx], pieces[length-2]
	case solvable /* && odd */ :
		pieces[empty_idx], pieces[length-1] = 0, pieces[empty_idx]
	case /* !solvable && */ d == 0:
		pieces[length-3], pieces[length-2] = pieces[length-2], pieces[length-3]
	case /* !solvable && */ even:
		pieces[empty_idx], pieces[length-1] = 0, pieces[empty_idx]
	default /* !solvable && odd */ :
		pieces[empty_idx], pieces[length-3], pieces[length-1] = 0, pieces[empty_idx], pieces[length-3]
	}

	return puzzle
}

func findIndex(s []uint, el uint) int {
	for i, v := range s {
		if v == el {
			return i
		}
	}
	return -1
}

func (p *Puzzle) Clone() Puzzle {
	pieces := make([]uint, len(p.pieces))
	copy(pieces, p.pieces)
	return Puzzle{pieces, p.width}
}

func (p *Puzzle) Solvable() bool {
	length := len(p.pieces)
	width := int(p.width)
	height := length / int(p.width)

	pieces := make([]uint, len(p.pieces))
	copy(pieces, p.pieces)

	empty_idx := findIndex(pieces, 0)
	empty_x, empty_y := empty_idx%width, empty_idx/width

	d := (width - 1 - empty_x) + (height - 1 - empty_y)
	switch {
	case d == 0:
	case d%2 == 0:
		pieces[empty_idx], pieces[length-1] = pieces[length-1], pieces[empty_idx]
	default:
		pieces[empty_idx], pieces[length-2], pieces[length-1] = pieces[length-2], pieces[length-1], pieces[empty_idx]
	}

	swaps := 0
	for i := range pieces[0 : length-2] {
		for {
			j := pieces[i] - 1
			if uint(i) == j {
				break
			}

			pieces[i], pieces[j] = pieces[j], pieces[i]
			swaps += 1
		}
	}

	return swaps%2 == 0
}

func toInt(b bool) (result int8) {
	if b {
		result = 1
	}
	return
}

func (p *Puzzle) SlideFrom(from_x, from_y uint) bool {
	width := p.Width()
	height := p.Height()
	from_idx := int(from_y)*width + int(from_x)

	if from_x >= uint(width) || from_y >= uint(height) {
		return false
	}

	empty_idx := findIndex(p.pieces, 0)
	empty_x, empty_y := empty_idx%width, empty_idx/width

	//            v  v  v v v v v empty_x
	// ord_x := |-1|-1|-1|0|1|1|1|
	//                    ^ from_x
	ord_x, ord_y := -toInt(empty_x < int(from_x))+toInt(int(from_x) < empty_x), -toInt(empty_y < int(from_y))+toInt(int(from_y) < empty_y)
	ord_x_eq, ord_y_eq := ord_x == 0, ord_y == 0

	switch {
	case ord_x_eq == ord_y_eq:
		return false

	case ord_y_eq:
		// y (outer index) is aligned; `copy_within`-optimized swapping
		row := p.pieces[int(from_y)*width : int(from_y+1)*width]

		switch ord_x {
		case -1:
			copy(row[empty_x:from_x], row[empty_x+1:from_x+1])
		case 1:
			copy(row[from_x+1:empty_x+1], row[from_x:empty_x])
		default:
			panic("unreachable")
		}

		row[from_x] = 0

	case ord_x_eq:
		// x (inner index) is aligned; swapping using loop
		delta := -int(ord_y) * width
		start := empty_y*width + empty_x
		end := int(from_y)*width + int(from_x)

		for cursor := start; cursor != end; cursor += delta {
			p.pieces[cursor] = p.pieces[cursor+delta]
		}
		p.pieces[from_idx] = 0
	}

	return true
}
