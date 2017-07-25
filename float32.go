// Code generated by go generate
// This file was generated by robots at 2017-07-25 13:05:51.639320872 +0000 UTC

package optional

// Float32 is an optional float32
type Float32 struct {
	float32 float32
	present bool
}

// EmptyFloat32 returns an empty optional.Float32
func EmptyFloat32() Float32 {
	return Float32{}
}

// OfFloat32 creates a optional.Float32 from a float32
func OfFloat32(f float32) Float32 {
	return Float32{float32: f, present: true}
}

// Set sets the float32 value
func (o *Float32) Set(f float32) {
	o.float32 = f
	o.present = true
}

// Get returns the float32 value
func (o *Float32) Get() float32 {
	return o.float32
}

// Present returns whether or not the value is present
func (o *Float32) Present() bool {
	return o.present
}

// OrElse returns the float32 value or a default value if the value is not present
func (o *Float32) OrElse(f float32) float32 {
	if o.present {
		return o.float32
	}
	return f
}

// If calls the function f with the value if the value is present
func (o *Float32) If(f func(f float32)) {
	if o.present {
		f(o.float32)
	}
}
