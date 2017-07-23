// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-21 19:01:30.648954797 +0000 UTC

package optional

// Uintptr is an optional uintptr
type Uintptr struct {
	uintptr uintptr
	present bool
}

// EmptyUintptr returns an empty optional.Uintptr
func EmptyUintptr() Uintptr {
	return Uintptr{}
}

// OfUintptr creates an optional.Uintptr from a uintptr
func OfUintptr(u uintptr) Uintptr {
	return Uintptr{uintptr: u, present: true}
}

// Set sets the uintptr value
func (o *Uintptr) Set(u uintptr) {
	o.uintptr = u
	o.present = true
}

// Uintptr returns the uintptr value
func (o *Uintptr) Uintptr() uintptr {
	return o.uintptr
}

// Present returns whether or not the value is present
func (o *Uintptr) Present() bool {
	return o.present
}

// OrElse returns the uintptr value or a default value if the value is not present
func (o *Uintptr) OrElse(u uintptr) uintptr {
	if o.present {
		return o.uintptr
	}
	return u
}
