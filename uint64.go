// Code generated by go generate
// This file was generated by robots at 2017-07-25 13:05:50.403670902 +0000 UTC

package optional

// Uint64 is an optional uint64
type Uint64 struct {
	uint64  uint64
	present bool
}

// EmptyUint64 returns an empty optional.Uint64
func EmptyUint64() Uint64 {
	return Uint64{}
}

// OfUint64 creates a optional.Uint64 from a uint64
func OfUint64(u uint64) Uint64 {
	return Uint64{uint64: u, present: true}
}

// Set sets the uint64 value
func (o *Uint64) Set(u uint64) {
	o.uint64 = u
	o.present = true
}

// Get returns the uint64 value
func (o *Uint64) Get() uint64 {
	return o.uint64
}

// Present returns whether or not the value is present
func (o *Uint64) Present() bool {
	return o.present
}

// OrElse returns the uint64 value or a default value if the value is not present
func (o *Uint64) OrElse(u uint64) uint64 {
	if o.present {
		return o.uint64
	}
	return u
}

// If calls the function f with the value if the value is present
func (o *Uint64) If(f func(u uint64)) {
	if o.present {
		f(o.uint64)
	}
}
