// Code generated by go generate
// This file was generated by robots at 2017-07-25 13:05:47.753464808 +0000 UTC

package optional

// Int is an optional int
type Int struct {
	int     int
	present bool
}

// EmptyInt returns an empty optional.Int
func EmptyInt() Int {
	return Int{}
}

// OfInt creates a optional.Int from a int
func OfInt(i int) Int {
	return Int{int: i, present: true}
}

// Set sets the int value
func (o *Int) Set(i int) {
	o.int = i
	o.present = true
}

// Get returns the int value
func (o *Int) Get() int {
	return o.int
}

// Present returns whether or not the value is present
func (o *Int) Present() bool {
	return o.present
}

// OrElse returns the int value or a default value if the value is not present
func (o *Int) OrElse(i int) int {
	if o.present {
		return o.int
	}
	return i
}

// If calls the function f with the value if the value is present
func (o *Int) If(f func(i int)) {
	if o.present {
		f(o.int)
	}
}
