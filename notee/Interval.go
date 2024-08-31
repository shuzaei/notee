package notee

type Accidental struct {
	Value int
}

type Interval struct {
	Value      int        // 1st, 2nd, 3rd, 4th, 5th, 6th, 7th
	Accidental Accidental // Natural, Sharp, Flat
}

var (
	Sharp   = Accidental{1}
	Flat    = Accidental{-1}
	Natural = Accidental{0}
)

var (
	I1 = Interval{1, Natural}
	I2 = Interval{2, Natural}
	I3 = Interval{3, Natural}
	I4 = Interval{4, Natural}
	I5 = Interval{5, Natural}
	I6 = Interval{6, Natural}
	I7 = Interval{7, Natural}

	IS1 = Interval{1, Sharp}
	IS2 = Interval{2, Sharp}
	IS3 = Interval{3, Sharp}
	IS4 = Interval{4, Sharp}
	IS5 = Interval{5, Sharp}
	IS6 = Interval{6, Sharp}
	IS7 = Interval{7, Sharp}

	IF1 = Interval{1, Flat}
	IF2 = Interval{2, Flat}
	IF3 = Interval{3, Flat}
	IF4 = Interval{4, Flat}
	IF5 = Interval{5, Flat}
	IF6 = Interval{6, Flat}
	IF7 = Interval{7, Flat}

	DefaultIntervals = []Interval{
		I1, I2, I3, I4, I5, I6, I7,
		IS1, IS2, IS3, IS4, IS5, IS6, IS7,
		IF1, IF2, IF3, IF4, IF5, IF6, IF7,
	}
)

func (i Interval) ToString() string {
	switch i {
	case I1:
		return "1"
	case I2:
		return "2"
	case I3:
		return "3"
	case I4:
		return "4"
	case I5:
		return "5"
	case I6:
		return "6"
	case I7:
		return "7"
	case IS1:
		return "1#"
	case IS2:
		return "2#"
	case IS3:
		return "3#"
	case IS4:
		return "4#"
	case IS5:
		return "5#"
	case IS6:
		return "6#"
	case IS7:
		return "7#"
	case IF1:
		return "1b"
	case IF2:
		return "2b"
	case IF3:
		return "3b"
	case IF4:
		return "4b"
	case IF5:
		return "5b"
	case IF6:
		return "6b"
	case IF7:
		return "7b"
	}

	return "Unknown"
}
