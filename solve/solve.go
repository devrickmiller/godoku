package solve

import (
	"fmt"
	"strconv"
)

type Mat [9][9][9]bool

func Solve(m Mat) (Mat, error){
	
	return m, nil
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

func (m *Mat) GetPossibleSquareValues(row int, col int) [] int {
	ints := []int{}
	for i, v := range m[row][col] {
		if v {
			ints = append(ints, i)
		}
	}
	return ints
}

func (m *Mat) SetRow(row int, ints []int) {
	for col := 0; col < 9; col++ {
		if ints[col] > 0 && ints[col] < 10 {
			m.SetSquare(row, col, ints[col]-1)
		}
	}
}

func (m *Mat) SetSquare(row int, col int, n int) {
	for i := 0; i < 9; i++ {
		if i != n {
			m[row][col][i] = false
		}
		if i != col {
			m.toggleOffSquare(row,i,n)
		}
		if i != row {
			m.toggleOffSquare(i,col,n)
		}
	}
	m.setOtherTriBlockSquares(row, col, n)
}

func (m *Mat) setOtherTriBlockSquares(row int, col int, n int) {
	frow, fcol := 0, 0
	if row > 5 {
		frow = 6
	} else if row > 2 {
		frow = 3
	}
	if col > 5 {
		fcol = 6
	} else if col > 2 {
		fcol = 3
	}
	for r := frow; r < frow + 3; r++ {
		for c := fcol; c < fcol + 3; c++ {
			if r != row && c != col {
				//m[r][c][n] = false
				m.toggleOffSquare(r,c,n)
			}
		}
	}
}

func (m *Mat) toggleOffSquare(row int, col int, n int) {
	toggle := m[row][col][n]
	m[row][col][n] = false
	valuesLeft := m.GetPossibleSquareValues(row,col)
	if len(valuesLeft) == 0 {
		pmessage := fmt.Sprintf("[[   row: %d col: %d  No values left!   ]]",row,col)
		panic(pmessage)
	} else if len(valuesLeft) == 1 && toggle {
		m.SetSquare(row, col, valuesLeft[0])
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
			vals := m.GetPossibleSquareValues(c,r)
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

