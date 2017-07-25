// Code generated by go generate
// This file was generated by robots at 2017-07-25 13:05:48.63557518 +0000 UTC

package optional

// Int32 is an optional int32
type Int32 struct {
	int32   int32
	present bool
}

// EmptyInt32 returns an empty optional.Int32
func EmptyInt32() Int32 {
	return Int32{}
}

// OfInt32 creates a optional.Int32 from a int32
func OfInt32(i int32) Int32 {
	return Int32{int32: i, present: true}
}

// Set sets the int32 value
func (o *Int32) Set(i int32) {
	o.int32 = i
	o.present = true
}

// Get returns the int32 value
func (o *Int32) Get() int32 {
	return o.int32
}

// Present returns whether or not the value is present
func (o *Int32) Present() bool {
	return o.present
}

// OrElse returns the int32 value or a default value if the value is not present
func (o *Int32) OrElse(i int32) int32 {
	if o.present {
		return o.int32
	}
	return i
}

// If calls the function f with the value if the value is present
func (o *Int32) If(f func(i int32)) {
	if o.present {
		f(o.int32)
	}
}
