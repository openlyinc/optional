// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-21 19:01:27.782945968 +0000 UTC

package optional

// String is an optional string
type String struct {
	string  string
	present bool
}

// EmptyString returns an empty optional.String
func EmptyString() String {
	return String{}
}

// OfString creates an optional.String from a string
func OfString(s string) String {
	return String{string: s, present: true}
}

// Set sets the string value
func (o *String) Set(s string) {
	o.string = s
	o.present = true
}

// String returns the string value
func (o *String) String() string {
	return o.string
}

// Present returns whether or not the value is present
func (o *String) Present() bool {
	return o.present
}

// OrElse returns the string value or a default value if the value is not present
func (o *String) OrElse(s string) string {
	if o.present {
		return o.string
	}
	return s
}
