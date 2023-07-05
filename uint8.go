// Code generated by 'go generate'

package optional

import (
	"encoding/json"
	"errors"
)

// Uint8 is an optional uint8.
type Uint8 struct {
	value *uint8
}

// NewUint8 creates an optional.Uint8 from a uint8.
func NewUint8(v uint8) Uint8 {
	return Uint8{&v}
}

// NewUint8FromPtr creates an optional.Uint8 from a uint8 pointer.
func NewUint8FromPtr(v *uint8) Uint8 {
	if v == nil {
		return Uint8{}
	}
	return NewUint8(*v)
}

// Set sets the uint8 value.
func (u *Uint8) Set(v uint8) {
	u.value = &v
}

// ToPtr returns a *uint8 of the value or nil if not present.
func (u Uint8) ToPtr() *uint8 {
	if !u.Present() {
		return nil
	}
	v := *u.value
	return &v
}

// Get returns the uint8 value or an error if not present.
func (u Uint8) Get() (uint8, error) {
	if !u.Present() {
		var zero uint8
		return zero, errors.New("value not present")
	}
	return *u.value, nil
}

// MustGet returns the uint8 value or panics if not present.
func (u Uint8) MustGet() uint8 {
	if !u.Present() {
		panic("value not present")
	}
	return *u.value
}

// Present returns whether or not the value is present.
func (u Uint8) Present() bool {
	return u.value != nil
}

// OrElse returns the uint8 value or a default value if the value is not present.
func (u Uint8) OrElse(v uint8) uint8 {
	if u.Present() {
		return *u.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (u Uint8) If(fn func(uint8)) {
	if u.Present() {
		fn(*u.value)
	}
}

func (u Uint8) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.value)
	}
	return json.Marshal(nil)
}

func (u *Uint8) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		u.value = nil
		return nil
	}

	var value uint8

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.value = &value
	return nil
}
