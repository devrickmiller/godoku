package solve

import (
	//"errors"
)

type Mat [9][9][9]bool

func Solve(m Mat) (Mat, error){
	
	return m, nil
}

func (m Mat) SetSquare(row int, col int, n int) {
	for i := 0; i < 9; i++ {
		if i != n {
			m[row][col][i] = false
		}
		if i != col {
			m[row][i][n] = false
		}
		if i != row {
			m[i][col][n] = false
		}
	}
	m.setOtherTriBlockSquares(row, col, n)
}

func (m Mat) setOtherTriBlockSquares(row int, col int, n int) {
	frow, fcol := 0, 0
	if row > 5 {
		frow = 6
	} else if row > 3 {
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
				m[r][c][n] = false
			}
		}
	}
}

func NewMat() *Mat {
	var m Mat
	m.initialize()
	return &m
}

func (m Mat) initialize() {
	for ri := 0; ri < 9; ri++ {
		for ci := 0; ci < 9; ci++ {
			for n := 0; n < 9; n++ {
				m[ri][ci][n] = true
			}
		}
	}
}

