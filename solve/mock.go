package solve

type Difficulty uint8

const (
	EASY Difficulty = iota
	MEDIUM
	HARD
	HARDEST
)

func (m *Mat) MockMat(d Difficulty, n int) {
	switch d {
	case EASY:
		switch n {
		case 1:
			m.SetRow(0,[]int{2,-1,-1,3,-1,-1,8,-1,1})
			m.SetRow(1,[]int{-1,-1,-1,-1,-1,2,-1,-1,-1})
			m.SetRow(2,[]int{-1,4,-1,-1,-1,9,-1,-1,-1})
			m.SetRow(3,[]int{-1,-1,8,-1,-1,1,3,6,-1})
			m.SetRow(4,[]int{1,7,-1,-1,-1,-1,-1,-1,-1})
			m.SetRow(5,[]int{3,6,-1,8,-1,-1,-1,-1,-1})
			m.SetRow(6,[]int{-1,9,-1,1,4,7,-1,2,-1})
			m.SetRow(7,[]int{6,3,-1,-1,2,-1,-1,4,1})
			m.SetRow(8,[]int{-1,-1,-1,-1,-1,-1,9,-1,-1})
		default:
						
		}
	case MEDIUM:
		switch n {
		case 1:
			m.SetRow(0,[]int{-1,-1,-1,5,1,-1,8,-1,2})
			m.SetRow(1,[]int{-1,-1,-1,-1,-1,3,7,-1,-1})
			m.SetRow(2,[]int{6,5,-1,-1,-1,-1,-1,-1,9})
			m.SetRow(3,[]int{-1,-1,1,3,-1,8,-1,-1,-1})
			m.SetRow(4,[]int{9,8,4,-1,7,-1,1,5,3})
			m.SetRow(5,[]int{-1,-1,-1,4,-1,1,6,-1,-1})
			m.SetRow(6,[]int{2,-1,-1,-1,-1,-1,-1,8,4})
			m.SetRow(7,[]int{-1,-1,6,9,-1,-1,-1,-1,-1})
			m.SetRow(8,[]int{3,-1,5,-1,8,2,-1,-1,-1})
		case 2:
			m.SetRow(0,[]int{5,3,-1,-1,-1,4,-1,-1,-1})
			m.SetRow(1,[]int{-1,-1,4,1,-1,-1,-1,-1,-1})
			m.SetRow(2,[]int{1,-1,-1,9,-1,6,-1,8,3})
			m.SetRow(3,[]int{9,7,-1,3,2,-1,6,-1,-1})
			m.SetRow(4,[]int{-1,-1,-1,-1,4,-1,-1,-1,-1})
			m.SetRow(5,[]int{-1,-1,3,-1,6,9,-1,7,2})
			m.SetRow(6,[]int{3,5,-1,6,-1,7,-1,-1,4})
			m.SetRow(7,[]int{-1,-1,-1,-1,-1,2,7,-1,-1})
			m.SetRow(8,[]int{-1,-1,-1,8,-1,-1,-1,3,1})						
		}
	case HARD:
		switch n {
		case 1:
			m.SetRow(0,[]int{-1,-1,-1,-1,-1,-1,8,-1,6})
			m.SetRow(1,[]int{4,-1,5,6,9,-1,-1,1,-1})
			m.SetRow(2,[]int{-1,-1,9,-1,-1,2,4,-1,-1})
			m.SetRow(3,[]int{5,-1,-1,-1,-1,3,-1,8,-1})
			m.SetRow(4,[]int{-1,-1,7,8,-1,9,6,-1,-1})
			m.SetRow(5,[]int{-1,9,-1,2,-1,-1,-1,-1,3})
			m.SetRow(6,[]int{-1,-1,4,7,-1,-1,1,-1,-1})
			m.SetRow(7,[]int{-1,6,-1,-1,4,1,7,-1,8})
			m.SetRow(8,[]int{7,-1,3,-1,-1,-1,-1,-1,-1})
		case 2:
			m.SetRow(0,[]int{9,-1,5,-1,-1,1,-1,-1,-1})
			m.SetRow(1,[]int{4,-1,3,9,-1,-1,-1,5,-1})
			m.SetRow(2,[]int{-1,8,-1,7,5,-1,-1,-1,-1})
			m.SetRow(3,[]int{-1,5,1,-1,-1,-1,-1,-1,3})
			m.SetRow(4,[]int{8,4,-1,-1,-1,-1,-1,7,6})
			m.SetRow(5,[]int{6,-1,-1,-1,-1,-1,1,8,-1})
			m.SetRow(6,[]int{-1,-1,-1,-1,9,6,-1,1,-1})
			m.SetRow(7,[]int{-1,9,-1,-1,-1,3,8,-1,7})
			m.SetRow(8,[]int{-1,-1,-1,1,-1,-1,5,-1,9})						
		}
	case HARDEST:
		switch n {
		case 1:
			m.SetRow(0,[]int{9,-1,3,7,-1,8,6,-1,2})
			m.SetRow(1,[]int{7,-1,-1,-1,6,-1,-1,-1,-1})
			m.SetRow(2,[]int{-1,-1,8,-1,-1,-1,-1,7,-1})
			m.SetRow(3,[]int{6,-1,9,2,8,-1,-1,-1,-1})
			m.SetRow(4,[]int{-1,8,-1,-1,-1,-1,-1,1,-1})
			m.SetRow(5,[]int{-1,-1,-1,-1,7,4,8,-1,5})
			m.SetRow(6,[]int{-1,3,-1,-1,-1,-1,9,-1,-1})
			m.SetRow(7,[]int{-1,-1,-1,-1,9,-1,-1,-1,4})
			m.SetRow(8,[]int{1,-1,6,8,-1,5,7,-1,3})
		}
	}





		
}


// m.SetRow(0,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})
// m.SetRow(1,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})
// m.SetRow(2,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})
// m.SetRow(3,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})
// m.SetRow(4,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})
// m.SetRow(5,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})
// m.SetRow(6,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})
// m.SetRow(7,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})
// m.SetRow(8,[]int{-1,-1,-1,-1,-1,-1,-1,-1,-1})	