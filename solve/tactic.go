package solve

import (
	//"fmt"
)

type SquareGroup struct {
	no      int
	squares []Square
	values  []int
	valid   bool
}

type SquareGroupSet []SquareGroup

func (m *Mat) EliminateGroupedRowValues() {
	for r := 0; r < 9; r++ {
		squares := GetRowSquares(r)
		m.EliminateGroupedValuesBySquares(squares)
	}
}

func (m *Mat) EliminateGroupedColValues() {
	for c := 0; c < 9; c++ {
		squares := GetColSquares(c)
		m.EliminateGroupedValuesBySquares(squares)
	}
}

func (m *Mat) EliminateGroupedBlockValues() {
	for b := 0; b < 9; b++ {
		squares := GetBlockSquares(b)
		m.EliminateGroupedValuesBySquares(squares)
	}
}

func (m *Mat) EliminateGroupedValuesBySquares(squares []Square) {
	sgs := SquareGroupSet{}
	for _, s := range squares {
		vals := m.GetPossibleSquareValues(s)
		sgs.EvaluateSquareForGroup(s, vals)
	}
	sgs.ValidateGroups()
	m.ToggleSquareGroupValues(&sgs)
}

func (m *Mat) ToggleSquareGroupValues(sgs *SquareGroupSet) {
	for _, sg := range *sgs {
		groupNo := sg.no
		if sg.valid {
			for _, sgo := range *sgs {
				if sgo.no != groupNo {
					for _, v := range sg.values {
						for _, s := range sgo.squares {
							m.toggleOffSquare(s, v)
						}
					}
				}
			}
		}
	}
}

func (sgs SquareGroupSet) ValidateGroups() {
	for i, sg := range sgs {
		if len(sg.values) > 1 && len(sg.values) == len(sg.squares) {
			sgs[i].valid = true
		}
	}
}

func (sgs *SquareGroupSet) EvaluateSquareForGroup(s Square, vals []int) {
	if len(vals) < 1 {
		panic("EvaluateSquare had less than one possible value")
	} else if len(vals) == 1 {
		return
	}
	added := false
	groupNoMax := 0
	for i, v := range *sgs {
		if v.no > groupNoMax {
			groupNoMax = v.no
		}
		if ValsAreEqual(v.values, vals) {
			(*sgs)[i].squares = append((*sgs)[i].squares, s)
			added = true
			break
		}
	}
	if !added {
		sg := SquareGroup{}
		sg.no = groupNoMax + 1
		sg.squares = append(sg.squares, s)
		sg.values = vals
		*sgs = append(*sgs, sg)
	}
}

func ValsAreEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func GetRowSquares(r int) []Square {
	squares := []Square{}
	for i := 0; i < 9; i++ {
		squares = append(squares, Square{r,i})
	}
	return squares
}

func GetColSquares(c int) []Square {
	squares := []Square{}
	for i := 0; i < 9; i++ {
		squares = append(squares, Square{i,c})
	}
	return squares
}

func GetBlockSquares(b int) []Square {
	s := []Square{}
	for i := 0; i < 9; i++ {
		s = append(s, Square{((b / 3) * 3) + (i / 3), ((b % 3) * 3) + (i % 3)})
	}

	return s
}




// eliminate grouped possible numbers for a given set
// find groups in set
// for each set eliminate grouped numbers for squares not in given group
	//get numbers for given group
	//eliminate those numbers for all other squares in group