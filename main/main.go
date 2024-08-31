package main

import (
	"fmt"

	"notee/notee"
)

func main() {
	// Create a scale from a list of semitone codes
	scale := notee.NewScale([]string{"C", "D", "E", "F", "G", "A", "B"})
	fmt.Println(scale)

	// Create a chord from a scale and a list of degrees
	chord := notee.Chord{
		Scale:   scale,
		Degrees: []int{0, 2, 4},
	}
	fmt.Println(chord)

	// Get the semitones of the chord
	degrees := chord.GetDegrees()
	fmt.Println(degrees)

	// Get the codes of the semitones
	codes := []string{}
	for _, degree := range degrees {
		codes = append(codes, degree.ToCode())
	}

	fmt.Println(codes)
}
