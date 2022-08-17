package main

type Packing interface {
	Pack() string
}

type Wrapper struct {
}

func (w *Wrapper) Pack() string {
	return "Wrapper Pck"
}

type Bottle struct {
}

func (w *Bottle) Pack() string {
	return "Bottle Pck"
}
