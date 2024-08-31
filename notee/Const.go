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

	CSMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			CS, DS, FS, GS, AS, B, CS,
			D, E, FS, G, A, B, CS,
			C, D, E, FS, G, A, B,
		},
	}
	CSMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			CS, DS, E, FS, GS, A, B,
			D, E, FS, G, A, B, CS,
			C, D, DS, F, G, GS, AS,
		},
	}

	DMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			D, E, FS, G, A, B, CS,
			DS, FS, G, GS, AS, CS, DS,
			C, D, E, FS, G, A, B,
		},
	}
	DMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			D, E, F, G, A, AS, C,
			DS, F, G, GS, AS, C, D,
			C, D, DS, F, G, GS, AS,
		},
	}

	DSMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			DS, F, G, GS, AS, C, D,
			E, G, GS, A, B, D, E,
			D, E, F, G, A, B, C,
		},
	}
	DSMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			DS, F, FS, G, AS, B, C,
			E, FS, G, A, B, C, DS,
			D, DS, E, F, G, A, AS,
		},
	}

	EMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			E, FS, GS, A, B, CS, DS,
			F, GS, A, AS, C, DS, E,
			D, E, FS, GS, A, B, CS,
		},
	}
	EMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			E, FS, G, A, B, C, D,
			F, G, A, AS, C, D, E,
			D, E, F, G, A, AS, B,
		},
	}

	FMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			F, G, A, AS, C, D, E,
			FS, A, AS, B, DS, E, FS,
			E, FS, G, A, AS, C, D,
		},
	}
	FMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			F, G, GS, A, AS, C, D,
			FS, GS, A, B, C, D, E,
			E, F, FS, G, A, B, CS,
		},
	}

	FSMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			FS, GS, AS, B, CS, DS, E,
			G, AS, B, C, DS, E, FS,
			F, FS, GS, AS, B, CS, D,
		},
	}
	FSMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			FS, GS, A, B, C, D, E,
			G, A, B, CS, D, E, FS,
			F, FS, G, A, B, CS, DS,
		},
	}

	GMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			G, A, B, C, D, E, FS,
			GS, B, C, CS, E, FS, G,
			F, G, A, B, C, D, E,
		},
	}
	GMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			G, A, AS, C, D, DS, F,
			GS, AS, C, CS, DS, F, G,
			F, G, GS, A, C, CS, DS,
		},
	}

	GSMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			GS, AS, C, CS, DS, F, G,
			A, C, CS, D, FS, G, GS,
			G, GS, AS, C, CS, DS, E,
		},
	}
	GSMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			GS, AS, B, C, CS, D, E,
			A, B, C, D, E, FS, GS,
			G, GS, A, B, C, D, DS,
		},
	}

	AMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			A, B, CS, D, E, FS, GS,
			AS, CS, D, DS, FS, GS, A,
			GS, A, B, CS, D, E, FS,
		},
	}
	AMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			A, B, C, D, E, F, G,
			AS, C, D, DS, F, G, A,
			G, A, AS, C, D, DS, E,
		},
	}

	ASMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			AS, C, D, DS, F, G, A,
			B, D, DS, E, G, A, AS,
			A, AS, C, D, DS, F, G,
		},
	}
	ASMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			AS, C, CS, D, F, FS, G,
			B, CS, D, E, FS, G, AS,
			A, AS, B, C, D, E, F,
		},
	}

	BMajor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			B, CS, DS, E, FS, GS, AS,
			C, DS, E, F, GS, AS, B,
			AS, B, CS, DS, E, FS, GS,
		},
	}
	BMinor = Scale{
		Interval: DefaultIntervals,
		Keys: []Key{
			B, CS, D, E, FS, G, A,
			C, D, E, F, G, A, B,
			AS, B, C, D, E, F, FS,
		},
	}
)
