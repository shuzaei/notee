package notee

type Translation struct {
	Value      int        // I, II, III, IV, V, VI, VII
	Accidental Accidental // Natural, Sharp, Flat
}

var (
	I   = Translation{0, Natural}
	II  = Translation{1, Natural}
	III = Translation{2, Natural}
	IV  = Translation{3, Natural}
	V   = Translation{4, Natural}
	VI  = Translation{5, Natural}
	VII = Translation{6, Natural}

	IS   = Translation{0, Sharp}
	IIS  = Translation{1, Sharp}
	IIIS = Translation{2, Sharp}
	IVS  = Translation{3, Sharp}
	VS   = Translation{4, Sharp}
	VIS  = Translation{5, Sharp}
	VIIS = Translation{6, Sharp}

	IF   = Translation{0, Flat}
	IIF  = Translation{1, Flat}
	IIIF = Translation{2, Flat}
	IVF  = Translation{3, Flat}
	VF   = Translation{4, Flat}
	VIF  = Translation{5, Flat}
	VIIF = Translation{6, Flat}

	DefaultTranslations = []Translation{
		I, II, III, IV, V, VI, VII,
		IS, IIS, IIIS, IVS, VS, VIS, VIIS,
		IF, IIF, IIIF, IVF, VF, VIF, VIIF,
	}
)

type Accidental struct {
	Value int
}

var (
	Sharp   = Accidental{1}
	Flat    = Accidental{-1}
	Natural = Accidental{0}
)

// a scale is a series of mappings from a key to a translation.
// translation is a relative position but also has an accidental to fill in the gaps.
// it is not 1:1 because there are 12 keys and 7 * 3 = 21 translations.
// thus, set primary keys in the first 12 translations.
type Scale struct {
	Translations []Translation // contains only 1 of each translation
	Keys         []Key         // contains multiple of the same key (the primary keys appear first)
}

func (s Scale) KeyToTranslation(key Key) Translation {
	for i, k := range s.Keys {
		if k == key {
			return s.Translations[i]
		}
	}

	return Translation{}
}

func (s Scale) TranslationToKey(translation Translation) Key {
	for i, t := range s.Translations {
		if t == translation {
			return s.Keys[i]
		}
	}

	return Key{}
}
