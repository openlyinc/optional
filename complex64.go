// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-21 19:01:31.927247438 +0000 UTC

package optional

// Complex64 is an optional complex64
type Complex64 struct {
	complex64 complex64
	present   bool
}

// EmptyComplex64 returns an empty optional.Complex64
func EmptyComplex64() Complex64 {
	return Complex64{}
}

// OfComplex64 creates an optional.Complex64 from a complex64
func OfComplex64(c complex64) Complex64 {
	return Complex64{complex64: c, present: true}
}

// Set sets the complex64 value
func (o *Complex64) Set(c complex64) {
	o.complex64 = c
	o.present = true
}

// Complex64 returns the complex64 value
func (o *Complex64) Complex64() complex64 {
	return o.complex64
}

// Present returns whether or not the value is present
func (o *Complex64) Present() bool {
	return o.present
}

// OrElse returns the complex64 value or a default value if the value is not present
func (o *Complex64) OrElse(c complex64) complex64 {
	if o.present {
		return o.complex64
	}
	return c
}
