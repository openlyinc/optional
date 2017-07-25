// Code generated by go generate
// This file was generated by robots at 2017-07-25 13:05:51.285857136 +0000 UTC

package optional

// Rune is an optional rune
type Rune struct {
	rune    rune
	present bool
}

// EmptyRune returns an empty optional.Rune
func EmptyRune() Rune {
	return Rune{}
}

// OfRune creates a optional.Rune from a rune
func OfRune(r rune) Rune {
	return Rune{rune: r, present: true}
}

// Set sets the rune value
func (o *Rune) Set(r rune) {
	o.rune = r
	o.present = true
}

// Get returns the rune value
func (o *Rune) Get() rune {
	return o.rune
}

// Present returns whether or not the value is present
func (o *Rune) Present() bool {
	return o.present
}

// OrElse returns the rune value or a default value if the value is not present
func (o *Rune) OrElse(r rune) rune {
	if o.present {
		return o.rune
	}
	return r
}

// If calls the function f with the value if the value is present
func (o *Rune) If(f func(r rune)) {
	if o.present {
		f(o.rune)
	}
}
