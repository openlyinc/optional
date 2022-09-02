// Code generated by 'go generate'

package optional

import (
	"encoding/json"
	"errors"
)

// Rune is an optional rune.
type Rune struct {
	value *rune
}

// NewRune creates an optional.Rune from a rune.
func NewRune(v rune) Rune {
	return Rune{&v}
}

// NewRuneFromPointer creates an optional.Rune from a rune pointer.
func NewRuneFromPointer(v *rune) Rune {
	if v == nil {
		return Rune{}
	}
	return NewRune(*v)
}

// Set sets the rune value.
func (r *Rune) Set(v rune) {
	r.value = &v
}

// ToPointer returns a *rune of the value or nil if not present.
func (r Rune) ToPointer() *rune {
	if !r.Present() {
		return nil
	}
	v := *r.value
	return &v
}

// Get returns the rune value or an error if not present.
func (r Rune) Get() (rune, error) {
	if !r.Present() {
		var zero rune
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

// MustGet returns the rune value or panics if not present.
func (r Rune) MustGet() rune {
	if !r.Present() {
		panic("value not present")
	}
	return *r.value
}

// Present returns whether or not the value is present.
func (r Rune) Present() bool {
	return r.value != nil
}

// OrElse returns the rune value or a default value if the value is not present.
func (r Rune) OrElse(v rune) rune {
	if r.Present() {
		return *r.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (r Rune) If(fn func(rune)) {
	if r.Present() {
		fn(*r.value)
	}
}

func (r Rune) MarshalJSON() ([]byte, error) {
	if r.Present() {
		return json.Marshal(r.value)
	}
	return json.Marshal(nil)
}

func (r *Rune) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		r.value = nil
		return nil
	}

	var value rune

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	r.value = &value
	return nil
}
