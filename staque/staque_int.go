package staque

type IntStaque []int

func Int() IntStaque {
	return IntStaque{}
}

func (stk *IntStaque) Push(xs ...int) {
	*stk = append(*stk, xs...)
}

func (stk *IntStaque) Peekstk() (int, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return 0, Empty("Cannot Peek() on empty staque")
	}
	return (*stk)[ilast], nil
}

func (que *IntStaque) Peekque() (int, error) {
	if len(*que) == 0 {
		return 0, Empty("Cannot Peek() on empty staque")
	}
	return (*que)[0], nil
}

func (stk *IntStaque) Popstk() (int, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return 0, Empty("Cannot Pop() on empty staque")
	}

	last := (*stk)[ilast] // save last value; it won't be available afterwards
	if ilast < cap(*stk) / 4 {
		*stk = append(make(IntStaque, 0, cap(*stk) / 2), (*stk)[:ilast]...)
	} else {
		*stk = (*stk)[:ilast]
	}
	return last, nil
}

func (que *IntStaque) Popque() (int, error) {
	len := len(*que)
	if len == 0 {
		return 0, Empty("Cannot Pop() on empty staque")
	}

	first := (*que)[0] // save first value; it won't be available afterwards
	if len > cap(*que) / 4 {
		*que = (*que)[1:]
	} else {
		*que = append(make(IntStaque, 0, cap(*que) / 2), (*que)[1:]...)
	}
	return first, nil
}