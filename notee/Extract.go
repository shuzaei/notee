package notee

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

// Score - []Scale -> ComplexPattern

func (s Score) ExtractComplex(scales []int, scalePalette []Scale) ComplexPattern {
	trans := make([]Interval, len(s.Notes))
	for i, n := range s.Notes {
		trans[i] = scalePalette[scales[i]].KeyToInterval(n.Key)
	}

	return ComplexPattern{
		Starts:       s.Starts,
		Lengths:      s.Lengths,
		Intervals:    trans,
		Scales:       scales,
		ScalePalette: scalePalette,
		Length:       s.Length,
	}
}
