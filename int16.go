// Code generated by 'go generate'

package optional

import (
	"encoding/json"
	"errors"
)

// Int16 is an optional int16.
type Int16 struct {
	value *int16
}

// NewInt16 creates an optional.Int16 from a int16.
func NewInt16(v int16) Int16 {
	return Int16{&v}
}

// NewInt16FromPtr creates an optional.Int16 from a int16 pointer.
func NewInt16FromPtr(v *int16) Int16 {
	if v == nil {
		return Int16{}
	}
	return NewInt16(*v)
}

// Set sets the int16 value.
func (i *Int16) Set(v int16) {
	i.value = &v
}

// ToPtr returns a *int16 of the value or nil if not present.
func (i Int16) ToPtr() *int16 {
	if !i.Present() {
		return nil
	}
	v := *i.value
	return &v
}

// Get returns the int16 value or an error if not present.
func (i Int16) Get() (int16, error) {
	if !i.Present() {
		var zero int16
		return zero, errors.New("value not present")
	}
	return *i.value, nil
}

// MustGet returns the int16 value or panics if not present.
func (i Int16) MustGet() int16 {
	if !i.Present() {
		panic("value not present")
	}
	return *i.value
}

// Present returns whether or not the value is present.
func (i Int16) Present() bool {
	return i.value != nil
}

// OrElse returns the int16 value or a default value if the value is not present.
func (i Int16) OrElse(v int16) int16 {
	if i.Present() {
		return *i.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (i Int16) If(fn func(int16)) {
	if i.Present() {
		fn(*i.value)
	}
}

func (i Int16) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.value)
	}
	return json.Marshal(nil)
}

func (i *Int16) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		i.value = nil
		return nil
	}

	var value int16

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.value = &value
	return nil
}
