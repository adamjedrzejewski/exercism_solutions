package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}

	out := make(Ints, 0)
	for _, v := range i {
		if filter(v) {
			out = append(out, v)
		}
	}
	return out
}

func (i Ints) Discard(filter func(int) bool) Ints {
	negate := func(i int) bool {
		return !filter(i)
	}
	return i.Keep(negate)
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	if l == nil {
		return nil
	}

	out := make(Lists, 0)
	for _, v := range l {
		if filter(v) {
			out = append(out, v)
		}
	}
	return out
}

func (s Strings) Keep(filter func(string) bool) Strings {
	if s == nil {
		return nil
	}

	out := make(Strings, 0)
	for _, v := range s {
		if filter(v) {
			out = append(out, v)
		}
	}
	return out
}
