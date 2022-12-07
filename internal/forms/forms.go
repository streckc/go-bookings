package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

// New initialize the form struct
func New(data url.Values) *Form {
	f := Form{
		data,
		errors(map[string][]string{}),
	}
	return &f
}

// Has checks if form field in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required takes a list of forme field names and makes them required.
func (f *Form) Required(fs ...string) {
	for _, field := range fs {
		v := f.Get(field)
		if strings.TrimSpace(v) == "" {
			f.Errors.Add(field, "This field can not be blank")
		}
	}
}

// MinLength checks for string mininum length
func (f *Form) MinLength(field string, l int) bool {
	v := f.Get(field)
	if len(strings.TrimSpace(v)) < l {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters", l))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

// Valid returns true if no errors exist
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
