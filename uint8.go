// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-21 19:01:29.598815741 +0000 UTC

package optional

// Uint8 is an optional uint8
type Uint8 struct {
	uint8   uint8
	present bool
}

// EmptyUint8 returns an empty optional.Uint8
func EmptyUint8() Uint8 {
	return Uint8{}
}

// OfUint8 creates an optional.Uint8 from a uint8
func OfUint8(u uint8) Uint8 {
	return Uint8{uint8: u, present: true}
}

// Set sets the uint8 value
func (o *Uint8) Set(u uint8) {
	o.uint8 = u
	o.present = true
}

// Uint8 returns the uint8 value
func (o *Uint8) Uint8() uint8 {
	return o.uint8
}

// Present returns whether or not the value is present
func (o *Uint8) Present() bool {
	return o.present
}

// OrElse returns the uint8 value or a default value if the value is not present
func (o *Uint8) OrElse(u uint8) uint8 {
	if o.present {
		return o.uint8
	}
	return u
}
