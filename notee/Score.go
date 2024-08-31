package notee

type Step struct {
	Value int
}

type Score struct {
	Starts  []Step
	Lengths []Step
	Notes   []Note
	Length  Step
}

type Pattern struct {
	Starts    []Step
	Lengths   []Step
	Intervals []Interval
	Scale     Scale
	Length    Step
}

// Score - Scale -> Pattern

func (s Score) Extract(scale Scale) Pattern {
	trans := make([]Interval, len(s.Notes))
	for i, n := range s.Notes {
		trans[i] = scale.KeyToInterval(n.Key)
	}

	return Pattern{
		Starts:    s.Starts,
		Lengths:   s.Lengths,
		Intervals: trans,
		Scale:     scale,
		Length:    s.Length,
	}
}

// Pattern * []Scale -> Pattern
