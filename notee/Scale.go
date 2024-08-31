package notee

// a scale is a series of mappings from a key to a interval.
// interval is a relative position but also has an accidental to fill in the gaps.
// it is not 1:1 because there are 12 keys and 7 * 3 = 21 interval.
// thus, set primary keys in the first 12 interval.
type Scale struct {
	Interval []Interval // contains only 1 of each interval
	Keys     []Key      // contains multiple of the same key (the primary keys appear first)
}

func (s Scale) KeyToInterval(key Key) Interval {
	for i, k := range s.Keys {
		if k == key {
			return s.Interval[i]
		}
	}

	return Interval{}
}

func (s Scale) IntervalToKey(interval Interval) Key {
	for i, t := range s.Interval {
		if t == interval {
			return s.Keys[i]
		}
	}

	return Key{}
}

func (s Scale) ToString() string {
	str := ""
	for i, k := range s.Keys {
		str += k.ToString() + " -> " + s.Interval[i].ToString() + "\n"
	}

	return str
}
