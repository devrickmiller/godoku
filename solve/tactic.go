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

func (m *Mat) EliminateGroupedValues() {
	m.AdjustMatFromSolvedSquares()
	m.EliminateGroupedRowValues()
	m.AdjustMatFromSolvedSquares()
	m.EliminateGroupedColValues()
	m.AdjustMatFromSolvedSquares()
	m.EliminateGroupedBlockValues()		
	m.AdjustMatFromSolvedSquares()	
}

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


func (m *Mat) SetUniqueValues() {
	m.AdjustMatFromSolvedSquares()
	m.SetUniqueRowValues()
	m.AdjustMatFromSolvedSquares()
	m.SetUniqueColValues()
	m.AdjustMatFromSolvedSquares()
	m.SetUniqueBlockValues()	
	m.AdjustMatFromSolvedSquares()		
}

// SetUniqueRowValues
func (m *Mat) SetUniqueRowValues() {
	for r := 0; r < 9; r++ {
		squares := GetRowSquares(r)
		m.SetUniqueValuesForSquares(squares)
	}
}

// SetUniqueColValues
func (m *Mat) SetUniqueColValues() {
	for c := 0; c < 9; c++ {
		squares := GetColSquares(c)
		m.SetUniqueValuesForSquares(squares)
	}
}

// SetUniqueBlockValues
func (m *Mat) SetUniqueBlockValues() {
	for b := 0; b < 9; b++ {
		squares := GetBlockSquares(b)
		m.SetUniqueValuesForSquares(squares)
	}
}
// SetUniqueValuesForSquares

func (m *Mat) SetUniqueValuesForSquares(squares []Square) {
	for n := 0; n < 9; n++ {
		foundIndex := -1
		occurs := 0
		for i, s := range squares {
			if m[s.Row][s.Col][n] {
				foundIndex = i
				occurs++
			}
		}
		if occurs == 0 {
			panic("this should never happen")
		}
		if occurs == 1 {
			if len(m.GetPossibleSquareValues(squares[foundIndex])) > 1 {
				m.SetSquare(squares[foundIndex], n)
			}
		}
	}

}