package notee

type Key struct {
	Value int
}

var (
	A  = Key{0}
	AS = Key{1}
	B  = Key{2}
	C  = Key{3}
	CS = Key{4}
	D  = Key{5}
	DS = Key{6}
	E  = Key{7}
	F  = Key{8}
	FS = Key{9}
	G  = Key{10}
	GS = Key{11}
)

func (k Key) ToString() string {
	switch k {
	case A:
		return "A"
	case AS:
		return "A#"
	case B:
		return "B"
	case C:
		return "C"
	case CS:
		return "C#"
	case D:
		return "D"
	case DS:
		return "D#"
	case E:
		return "E"
	case F:
		return "F"
	case FS:
		return "F#"
	case G:
		return "G"
	case GS:
		return "G#"
	}
	return "Unknown"
}
