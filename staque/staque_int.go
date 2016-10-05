package staque

type IntSlice []int

func Int() IntSlice {
	return IntSlice{}
}

func (stk *IntSlice) Push(xs ...int) {
	*stk = append(*stk, xs...)
}

func (stk *IntSlice) Peekstk() (int, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return 0, Empty("Cannot Peek() on empty staque")
	}
	return (*stk)[ilast], nil
}

func (que *IntSlice) Peekque() (int, error) {
	if len(*que) == 0 {
		return 0, Empty("Cannot Peek() on empty staque")
	}
	return (*que)[0], nil
}

func (stk *IntSlice) Popstk() (int, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return 0, Empty("Cannot Pop() on empty staque")
	}

	last := (*stk)[ilast] // save last value; it won't be available afterwards
	if ilast < cap(*stk) / 4 {
		*stk = append(make(IntSlice, 0, cap(*stk) / 2), (*stk)[:ilast]...)
	} else {
		*stk = (*stk)[:ilast]
	}
	return last, nil
}

func (que *IntSlice) Popque() (int, error) {
	len := len(*que)
	if len == 0 {
		return 0, Empty("Cannot Pop() on empty staque")
	}

	first := (*que)[0] // save first value; it won't be available afterwards
	if len > cap(*que) / 4 {
		*que = (*que)[1:]
	} else {
		*que = append(make(IntSlice, 0, cap(*que) / 2), (*que)[1:]...)
	}
	return first, nil
}