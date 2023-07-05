// Code generated by 'go generate'

package optional

import (
	"encoding/json"
	"errors"
)

// Int64 is an optional int64.
type Int64 struct {
	value *int64
}

// NewInt64 creates an optional.Int64 from a int64.
func NewInt64(v int64) Int64 {
	return Int64{&v}
}

// NewInt64FromPtr creates an optional.Int64 from a int64 pointer.
func NewInt64FromPtr(v *int64) Int64 {
	if v == nil {
		return Int64{}
	}
	return NewInt64(*v)
}

// Set sets the int64 value.
func (i *Int64) Set(v int64) {
	i.value = &v
}

// ToPtr returns a *int64 of the value or nil if not present.
func (i Int64) ToPtr() *int64 {
	if !i.Present() {
		return nil
	}
	v := *i.value
	return &v
}

// Get returns the int64 value or an error if not present.
func (i Int64) Get() (int64, error) {
	if !i.Present() {
		var zero int64
		return zero, errors.New("value not present")
	}
	return *i.value, nil
}

// MustGet returns the int64 value or panics if not present.
func (i Int64) MustGet() int64 {
	if !i.Present() {
		panic("value not present")
	}
	return *i.value
}

// Present returns whether or not the value is present.
func (i Int64) Present() bool {
	return i.value != nil
}

// OrElse returns the int64 value or a default value if the value is not present.
func (i Int64) OrElse(v int64) int64 {
	if i.Present() {
		return *i.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (i Int64) If(fn func(int64)) {
	if i.Present() {
		fn(*i.value)
	}
}

func (i Int64) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.value)
	}
	return json.Marshal(nil)
}

func (i *Int64) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		i.value = nil
		return nil
	}

	var value int64

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.value = &value
	return nil
}
