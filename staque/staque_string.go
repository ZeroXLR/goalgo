package staque

type StringStaque []string

func String() StringStaque {
	return StringStaque{}
}

func (stk *StringStaque) Push(xs ...string) {
	*stk = append(*stk, xs...)
}

func (stk *StringStaque) Peekstk() (string, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return "", Empty("Cannot Peek() on empty staque")
	}
	return (*stk)[ilast], nil
}

func (que *StringStaque) Peekque() (string, error) {
	if len(*que) == 0 {
		return "", Empty("Cannot Peek() on empty staque")
	}
	return (*que)[0], nil
}

func (stk *StringStaque) Popstk() (string, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return "", Empty("Cannot Pop() on empty staque")
	}

	last := (*stk)[ilast] // save last value; it won't be available afterwards
	if ilast < cap(*stk) / 4 {
		*stk = append(make(StringStaque, 0, cap(*stk) / 2), (*stk)[:ilast]...)
	} else {
		*stk = (*stk)[:ilast]
	}
	return last, nil
}

func (que *StringStaque) Popque() (string, error) {
	len := len(*que)
	if len == 0 {
		return "", Empty("Cannot Pop() on empty staque")
	}

	first := (*que)[0] // save first value; it won't be available afterwards
	if len > cap(*que) / 4 {
		*que = (*que)[1:]
	} else {
		*que = append(make(StringStaque, 0, cap(*que) / 2), (*que)[1:]...)
	}
	return first, nil
}