package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	accumulator := initial
	for _, v := range s {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accumulator := initial
	for i := len(s) - 1; i >= 0; i-- {
		accumulator = fn(s[i], accumulator)
	}
	return accumulator
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			out = append(out, v)
		}
	}

	return out
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	out := make(IntList, len(s))
	for i, v := range s {
		out[i] = fn(v)
	}
	return out
}

func (s IntList) Reverse() IntList {
	out := make(IntList, len(s))
	i := 0
	j := len(s) - 1
	for i < j {
		out[i], out[j] = s[j], s[i]
		i++
		j--
	}
	return out
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	out := make(IntList, len(s))
	copy(out, s)
	for _, list := range lists {
		out = append(out, list...)
	}

	return out
}
