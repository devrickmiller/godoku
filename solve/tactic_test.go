package solve

import (
	"testing"
	//"fmt"
)

func TestGetRowSquares(t *testing.T) {
	want := []Square{{0,0},{0,1},{0,2},{0,3},{0,4},{0,5},{0,6},{0,7},{0,8},}
	got := GetRowSquares(0)

	if len(got) != 9 {
		t.Error("returned too many Squares")
	}
	for i := 0; i < 9; i++ {
		if got[i] != want[i] {
			t.Errorf("want: %v  got: %v", want, got)
		}
	}
}

func TestGetColSquares(t *testing.T) {
	want := []Square{{0,0},{1,0},{2,0},{3,0},{4,0},{5,0},{6,0},{7,0},{8,0},}
	got := GetColSquares(0)

	if len(got) != 9 {
		t.Error("returned too many Squares")
	}
	for i := 0; i < 9; i++ {
		if got[i] != want[i] {
			t.Errorf("want: %v  got: %v", want, got)
		}
	}
}

func TestGetBlockSquares(t *testing.T) {
	want := []Square{{0,0},{0,1},{0,2},{1,0},{1,1},{1,2},{2,0},{2,1},{2,2},}
	got := GetBlockSquares(Square{0,0})

	if len(got) != 9 {
		t.Error("returned too many Squares")
	}
	for i := 0; i < 9; i++ {
		if got[i] != want[i] {
			t.Errorf("want: %v  got: %v", want, got)
		}
	}
}

func TestValsAreEqual(t *testing.T) {
	type TestCase struct {
		name    string
		input_a []int
		input_b []int
		want    bool
	}
	
	cases := []TestCase{
		{"positive",[]int{1,2,5}, []int{1,2,5}, true},
		{"unequal lens",[]int{1,2}, []int{1,2,5}, false},
		{"mismatch numbers",[]int{1,9,2}, []int{1,2,5}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if ValsAreEqual(tc.input_a, tc.input_b) != tc.want {
				t.Fail()
			}
		})
	}
}

// (sgs SquareGroupSet) ValidateGroups()
func TestValidateGroupSet(t *testing.T) {
	sgs := SquareGroupSet{}
	t.Run("Positive", func(t *testing.T) {
		sg1 := SquareGroup{}
		sg1.no = 1
		sg1.squares = append(sg1.squares, Square{0,0})
		sg1.squares = append(sg1.squares, Square{0,1})
		sg1.values = []int{5,6}
		sgs = append(sgs, sg1)
		sgs.ValidateGroups()
		if !sgs[0].valid {
			t.Fail()
		}
	})
	t.Run("Negative", func(t *testing.T) {
		sg2 := SquareGroup{}
		sg2.no = 2
		sg2.squares = append(sg2.squares, Square{2,0})
		sg2.squares = append(sg2.squares, Square{2,1})
		sg2.values = []int{1,2,9}
		sgs = append(sgs, sg2)
		sgs.ValidateGroups()
		if sgs[1].valid {
			t.Fail()
		}
	})
}

//func (sgs SquareGroupSet) EvaluateSquareForGroup(s Square, vals []int)
func TestEvaluateSquareForGroups(t *testing.T) {
	sgs := SquareGroupSet{}
	t.Run("Positive Same", func(t *testing.T) {
		sg1 := SquareGroup{}
		sg1.no = 1
		sg1.squares = append(sg1.squares, Square{0,0})
		sg1.squares = append(sg1.squares, Square{0,1})
		sg1.values = []int{5,6,9}
		sgs = append(sgs, sg1)
		newSquare := Square{0,2}
		sgs.EvaluateSquareForGroup(newSquare, []int{5,6,9})
		if len(sgs) != 1 {
			t.Errorf("SquareGroup not added")
		}
		sl := sgs[0].squares
		if len(sl) != 3 || sl[2].Row != newSquare.Row || sl[2].Col != newSquare.Col {
			t.Errorf("got: %v", sgs)
		}
	})
	t.Run("Positive New", func(t *testing.T) {
		newSquare := Square{0,2}
		sgs.EvaluateSquareForGroup(newSquare, []int{3,8})
		if len(sgs) != 2 {
			t.Errorf("SquareGroup not added")		
		}
		sl := sgs[1].squares
		if len(sl) != 1 || len(sgs) != 2 || len(sgs[1].values) != 2 || sgs[1].values[1] != 8 {
			t.Errorf("got: %v", sgs)
		}
	})
	t.Run("Negative Invalid", func(t *testing.T) {
		newSquare := Square{0,7}
		sgs.EvaluateSquareForGroup(newSquare, []int{5})
		if len(sgs) != 2 {
			t.Errorf("SquareGroup mistakenly added")		
		}
	})	
}

func TestToggleSquareGroupValues(t *testing.T) {
	t.Run("Postive Set of Two", func(t *testing.T) {
		m := Mat{}
		m.initialize()
		sgs := SquareGroupSet{}
		sg1 := SquareGroup{}
		sg1.no = 1
		sg1.squares = append(sg1.squares, Square{0,0})
		sg1.squares = append(sg1.squares, Square{0,1})
		sg1.values = []int{5,6}
		sg1.valid = true
		sgs = append(sgs, sg1)	
		sg2 := SquareGroup{}
		sg2.no = 2 
		sg2.squares = append(sg2.squares, Square{0,5})
		sg2.values = []int{1,2,3,4,5,6}
		sgs = append(sgs, sg2)
		m.ToggleSquareGroupValues(&sgs)
		vals := m.GetPossibleSquareValues(Square{0,5})
		for _, val := range vals {
			if val == 5 || val == 6 {
				t.Errorf("Values 5 and 6 not toggled off: %v", vals)
			}
		}
	})
	t.Run("Negative Set of Two", func(t *testing.T) {
		m := Mat{}
		m.initialize()
		sgs := SquareGroupSet{}
		sg1 := SquareGroup{}
		sg1.no = 1
		sg1.squares = append(sg1.squares, Square{0,0})
		sg1.squares = append(sg1.squares, Square{0,1})
		sg1.values = []int{5,6}
		sg1.valid = true
		sgs = append(sgs, sg1)	
		sg2 := SquareGroup{}
		sg2.no = 2 
		sg2.squares = append(sg2.squares, Square{0,5})
		sg2.values = []int{1,4,5,6}
		sgs = append(sgs, sg2)
		m.ToggleSquareGroupValues(&sgs)
		vals := m.GetPossibleSquareValues(Square{0,5})
		valsFound := false
		for _, val := range vals {
			if val == 1 || val == 4 {
				valsFound = true
			}
		}
		if !valsFound {
			t.Errorf("Values 1 or 4 were mistakenly toggled off: %v", vals)
		}
	})
}