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

type ComplexPattern struct {
	Starts       []Step
	Lengths      []Step
	Intervals    []Interval
	Scales       []int // index of ScalePalette
	ScalePalette []Scale
	Length       Step
}
