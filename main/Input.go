package main

import (
	"notee/notee"
)

// type Pattern struct {
// 	Starts       []Step
// 	Lengths      []Step
// 	Intervals []Interval
// 	Scale        Scale
// 	Length       Step
// }

var (
	RoyalRoadC = []notee.Scale{
		notee.FMajor,
		notee.GMajor,
		notee.EMinor,
		notee.AMinor,
		notee.DMinor,
		notee.GMajor,
		notee.CMajor,
		notee.CMajor,
		notee.FMajor,
		notee.GMajor,
		notee.EMinor,
		notee.AMinor,
		notee.DMinor,
		notee.GMajor,
		notee.CMajor,
		notee.CMajor,
	}

	basicChordPattern = notee.Pattern{
		Starts:    []notee.Step{{0}, {0}, {0}},
		Lengths:   []notee.Step{{1}, {1}, {1}},
		Intervals: []notee.Interval{notee.I1, notee.I3, notee.I5},
		Scale:     notee.CMajor,
		Length:    notee.Step{1},
	}

	piano3R2Pattern = notee.Pattern{
		Starts: []notee.Step{
			{0}, {0}, {0},
			{4}, {4}, {4},
			{7}, {7}, {7},
			{9}, {9}, {9},
			{12}, {12}, {12},
		},
		Lengths: []notee.Step{
			{1}, {1}, {1},
			{1}, {1}, {1},
			{1}, {1}, {1},
			{1}, {1}, {1},
			{1}, {1}, {1},
		},
		Intervals: []notee.Interval{
			notee.I1, notee.I3, notee.I5,
			notee.I1, notee.I3, notee.I5,
			notee.I1, notee.I3, notee.I5,
			notee.I1, notee.I3, notee.I5,
			notee.I1, notee.I3, notee.I5,
		},
		Scale:  notee.CMajor,
		Length: notee.Step{16},
	}

	piano3R2Pattern2x = notee.ComplexPattern{
		Starts: []notee.Step{
			{0}, {0}, {0},
			{4}, {4}, {4},
			{7}, {7}, {7},
			{9}, {9}, {9},
			{12}, {12}, {12},
		},
		Lengths: []notee.Step{
			{1}, {1}, {1},
			{1}, {1}, {1},
			{1}, {1}, {1},
			{1}, {1}, {1},
			{1}, {1}, {1},
		},
		Intervals: []notee.Interval{
			notee.I1, notee.I3, notee.I5,
			notee.I1, notee.I3, notee.I5,
			notee.I1, notee.I3, notee.I5,
			notee.I1, notee.I3, notee.I5,
			notee.I1, notee.I3, notee.I5,
		},
		Scales:       []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		ScalePalette: []notee.Scale{notee.CMajor, notee.GMajor},
		Length:       notee.Step{16},
	}
)
