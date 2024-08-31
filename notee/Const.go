package notee

var (
	CMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			C, D, E, F, G, A, B,
			CS, DS, F, FS, GS, AS,
			B, CS, DS, E, FS, GS, AS,
		},
	}

	CMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			C, D, DS, F, G, GS, AS,
			CS, DS, F, FS, GS, AS, B,
			B, CS, D, E, FS, G, A,
		},
	}
)
