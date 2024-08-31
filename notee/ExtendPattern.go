package notee

func (p Pattern) Extend(scales []Scale) Score {
	extended := Score{
		Starts:  []Step{},
		Lengths: []Step{},
		Notes:   []Note{},
		Length:  Step{0},
	}

	for _, scale := range scales {
		for i, start := range p.Starts {
			extended.Starts = append(extended.Starts,
				Step{extended.Length.Value + start.Value})
			extended.Lengths = append(extended.Lengths, p.Lengths[i])

			for j := 0; j < int(p.Lengths[i].Value); j++ {
				key := scale.IntervalToKey(p.Intervals[i])
				extended.Notes = append(extended.Notes, Note{Key: key})
			}
		}

		extended.Length = Step{extended.Length.Value + p.Length.Value}
	}

	return extended
}
