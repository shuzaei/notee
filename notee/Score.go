package notee

import (
	"errors"
	"strconv"
	"unicode"
)

type Semitone struct {
	Value int
}

type Scale struct {
	Semitones []Semitone
}

func NewScale(codes []string) Scale {
	semitones := []Semitone{}
	for _, code := range codes {
		semitone, _ := StringToSemitone(code)
		semitones = append(semitones, semitone)
	}
	return Scale{semitones}
}

type Chord struct {
	Scale   Scale
	Degrees []int
}

func (c Chord) GetDegrees() []Semitone {
	degrees := []Semitone{}
	for _, degree := range c.Degrees {
		degrees = append(degrees, c.Scale.Semitones[degree])
	}
	return degrees
}

// for internal use
var (
	aToSemitone = map[rune]int{
		'A': 0,
		'B': 2,
		'C': 3,
		'D': 5,
		'E': 7,
		'F': 8,
		'G': 10,
	}
	aToAccidental = map[rune]int{
		'#': 1,
		'b': -1,
		' ': 0,
	}

	semitoneToCode = map[int]string{
		0:  "A",
		1:  "A#",
		2:  "B",
		3:  "C",
		4:  "C#",
		5:  "D",
		6:  "D#",
		7:  "E",
		8:  "F",
		9:  "F#",
		10: "G",
		11: "G#",
	}
)

// for external use
var (
	MajorScale = NewScale([]string{"C", "D", "E", "F", "G", "A", "B"})
	MinorScale = NewScale([]string{"C", "D", "D#", "F", "G", "G#", "A#"})

	MajorChord = Chord{MajorScale, []int{0, 2, 4}}
	MinorChord = Chord{MinorScale, []int{0, 2, 4}}

	Major7Chord = Chord{MajorScale, []int{0, 2, 4, 6}}
	Minor7Chord = Chord{MinorScale, []int{0, 2, 4, 6}}
)

// @return Semitone which corresponds to the given code (e.g. "A0" -> 0)
func CodeToSemitone(scale rune, accidental rune, octave int) Semitone {
	semitone := 0
	semitone += aToSemitone[scale]
	semitone += aToAccidental[accidental]
	semitone += 12 * octave
	return Semitone{semitone}
}

// @return Semitone which corresponds to the given code (e.g. "A0" -> 0)
func StringToSemitone(code string) (Semitone, error) {
	if unicode.IsDigit(rune(code[len(code)-1])) {
		switch len(code) {
		case 2:
			return CodeToSemitone(rune(code[0]), ' ', int(code[1]-'0')), nil
		case 3:
			return CodeToSemitone(rune(code[0]), rune(code[1]), int(code[2]-'0')), nil
		default:
			return Semitone{}, errors.New("invalid code")
		}
	} else {
		switch len(code) {
		case 1:
			return CodeToSemitone(rune(code[0]), ' ', 0), nil
		case 2:
			return CodeToSemitone(rune(code[0]), rune(code[1]), 0), nil
		default:
			return Semitone{}, errors.New("invalid code")
		}
	}
}

// @return code which corresponds to the given semitone (e.g. 0 -> "A0")
func (s Semitone) ToCode() string {
	scale := semitoneToCode[s.Value%12]
	octave := s.Value / 12
	return scale + strconv.Itoa(octave)
}
