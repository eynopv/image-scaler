package data

type Size struct {
	Width  int
	Height int
}

func (s Size) IsNonNull() bool {
	return s.Width > 0 || s.Height > 0
}
