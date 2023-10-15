package solve

import (
	"fmt"
	"strconv"
)

type Mat [9][9][9]bool

type Square struct {
	Row    int
	Col    int
}

var ToggleCount int = 0

func (m *Mat) Solve() error {
	tc := 0
	for tc < ToggleCount {
		tc = ToggleCount
		m.AdjustMatFromSolvedSquares()
		m.EliminateGroupedRowValues()
		m.AdjustMatFromSolvedSquares()
		m.EliminateGroupedColValues()
		m.AdjustMatFromSolvedSquares()
		m.EliminateGroupedBlockValues()		
		m.AdjustMatFromSolvedSquares()
		fmt.Println("ToggleCount = ", ToggleCount)
	}

	return nil
}

func (m *Mat) MockMat() {
	// m.SetRow(0,[]int{2,-1,-1,3,-1,-1,8,-1,1})
	// m.SetRow(1,[]int{-1,-1,-1,-1,-1,2,-1,-1,-1})
	// m.SetRow(2,[]int{-1,4,-1,-1,-1,9,-1,-1,-1})
	// m.SetRow(3,[]int{-1,-1,8,-1,-1,1,3,6,-1})
	// m.SetRow(4,[]int{1,7,-1,-1,-1,-1,-1,-1,-1})
	// m.SetRow(5,[]int{3,6,-1,8,-1,-1,-1,-1,-1})
	// m.SetRow(6,[]int{-1,9,-1,1,4,7,-1,2,-1})
	// m.SetRow(7,[]int{6,3,-1,-1,2,-1,-1,4,1})
	// m.SetRow(8,[]int{-1,-1,-1,-1,-1,-1,9,-1,-1})

	m.SetRow(0,[]int{2,-1,-1,3,-1,-1,8,-1,-1})
	m.SetRow(1,[]int{-1,-1,-1,-1,-1,2,-1,-1,-1})
	m.SetRow(2,[]int{-1,4,-1,-1,-1,9,-1,-1,-1})
	m.SetRow(3,[]int{-1,-1,8,-1,-1,1,3,6,-1})
	m.SetRow(4,[]int{1,7,-1,-1,-1,-1,-1,-1,-1})
	m.SetRow(5,[]int{3,6,-1,8,-1,-1,-1,-1,-1})
	m.SetRow(6,[]int{-1,9,-1,1,4,7,-1,2,-1})
	m.SetRow(7,[]int{6,3,-1,-1,2,-1,-1,4,1})
	m.SetRow(8,[]int{-1,-1,-1,-1,-1,-1,9,-1,-1})	

	// m.SetRow(0,[]int{-1,6,4,-1,9,5,-1,-1,-1})
	// m.SetRow(1,[]int{5,-1,8,-1,-1,1,4,7,-1})
	// m.SetRow(2,[]int{9,1,2,-1,3,4,5,6,-1})
	// m.SetRow(3,[]int{-1,7,3,-1,6,-1,2,9,-1})
	// m.SetRow(4,[]int{6,-1,-1,-1,5,2,8,-1,7})
	// m.SetRow(5,[]int{-1,2,5,-1,-1,3,-1,4,1})
	// m.SetRow(6,[]int{-1,4,-1,2,8,9,7,-1,-1})
	// m.SetRow(7,[]int{3,-1,-1,5,-1,-1,-1,-1,6})
	// m.SetRow(8,[]int{2,-1,7,-1,1,-1,-1,8,-1})		
}

func (m *Mat) GetPossibleSquareValues(s Square) [] int {
	ints := []int{}
	for i, v := range m[s.Row][s.Col] {
		if v {
			ints = append(ints, i)
		}
	}
	return ints
}

func (m *Mat) SetRow(row int, ints []int) {
	for col := 0; col < 9; col++ {
		if ints[col] > 0 && ints[col] < 10 {
			m.SetSquare(Square{row, col}, ints[col]-1)
		}
	}
}

func (m *Mat) SetSquare(s Square, n int) {
	for i := 0; i < 9; i++ {
		if i != n {
			m.performToggle(s, i)
		}
	}
}

func (m *Mat) AdjustMatFromSolvedSquares() {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			s := Square{r,c}
			valCount := len(m.GetPossibleSquareValues(s))
			if valCount < 1 {
				fmt.Println("AdjustMatFromSolvedSquares() : s = ", s)
				panic("AdjustMatFromSolvedSquares() : valCount < 1")
			} else if valCount == 1 {
				m.toggleOffRelatedSquares(s)
			}
		}
	} 
}

func (m *Mat) toggleOffRelatedSquares(s Square) {
	vals := m.GetPossibleSquareValues(s)
	n := vals[0]
	if len(vals) == 1 {
		for i := 0; i < 9; i++ {
			if i != s.Col {
				m.toggleOffSquare(Square{s.Row,i},n)
			}
			if i != s.Row {
				m.toggleOffSquare(Square{i,s.Col},n)
			}
		}
		m.setOtherTriBlockSquares(s, n)
	} else if len(vals) < 1 {
		fmt.Println("AdjustMatFromSolvedSquares() : s = ", s)
		panic("AdjustMatFromSolvedSquares() : valCount < 1")
	} else {
		fmt.Println("toggleOffRelatedSquares(): exception.. Should be unreachable")
	}
}

func (m *Mat) setOtherTriBlockSquares(s Square, n int) {
	frow, fcol := 0, 0
	if s.Row > 5 {
		frow = 6
	} else if s.Row > 2 {
		frow = 3
	}
	if s.Col > 5 {
		fcol = 6
	} else if s.Col > 2 {
		fcol = 3
	}
	for r := frow; r < frow + 3; r++ {
		for c := fcol; c < fcol + 3; c++ {
			if r != s.Row && c != s.Col {
				m.toggleOffSquare(Square{r,c},n)
			}
		}
	}
}

func (m *Mat) toggleOffSquare(s Square, n int) {
	toggle := m[s.Row][s.Col][n]
	m.performToggle(s, n)
	valuesLeft := m.GetPossibleSquareValues(s)
	if len(valuesLeft) == 0 {
		pmessage := fmt.Sprintf("[[   row: %d col: %d  No values left!   ]]",s.Row,s.Col)
		panic(pmessage)
	} else if len(valuesLeft) == 1 && toggle {
		m.toggleOffRelatedSquares(s)
	}	
}

func (m *Mat) performToggle(s Square, n int) {
	if m[s.Row][s.Col][n] {
		m[s.Row][s.Col][n] = false
		ToggleCount++
	}
}

func NewMat() *Mat {
	var m Mat
	m.initialize()
	return &m
}

func (m *Mat) initialize() {
	for ri := 0; ri < 9; ri++ {
		for ci := 0; ci < 9; ci++ {
			for n := 0; n < 9; n++ {
				m[ri][ci][n] = true
			}
		}
	}
}

func (m *Mat) String() string {
	s := "╔═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╗\n"
	for c := 0; c < 9; c++ {
		s += "║"
		for r := 0; r < 9; r++ {
			vals := m.GetPossibleSquareValues(Square{c,r})
			if len(vals) == 1 {
				s += " " + strconv.Itoa(vals[0]+1) + " "
			} else {
				s += "   "
			}
			if r != 8 {
				s += "│"
			}
		}
		s += "║\n"
		if c != 8 {
			s += "╟───┼───┼───┼───┼───┼───┼───┼───┼───╢\n"
		}
	}
	s += "╚═══╧═══╧═══╧═══╧═══╧═══╧═══╧═══╧═══╝"
	return s
}

