package Variant

type colorVariant uint

const (
	rgbVariant colorVariant = iota
	vhlVariant
)

type Color interface {
	isVariant(variant colorVariant) bool
}

type RGB struct {
	R uint
	G uint
	B uint
}

func (R RGB) isVariant(variant colorVariant) bool {
	switch variant {
	case rgbVariant:
		return true
	default:
		return false
	}
}
