package service

import (
	"convy/internal/errorext"
	"regexp"
)

type Validation interface {
	SetEmail(email string) *ValidationBuilder
	SetPassword(password string) *ValidationBuilder
	SetUsername(username string) *ValidationBuilder
	SetIds(ids ...uint) *ValidationBuilder
	Validate() (bool, error)
}

type valueAndValidation struct {
	value     any
	validator func(any) bool
}

type ValidationBuilder struct {
	email    valueAndValidation
	password valueAndValidation
	username valueAndValidation
	id       valueAndValidation
}

func NewValidation() Validation {
	return ValidationBuilder{}
}

func (v ValidationBuilder) SetEmail(email string) *ValidationBuilder {
	v.email.value = email
	v.email.validator = func(a any) bool {
		emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		return emailRegex.MatchString(a.(string))
	}
	return &v
}

func (v ValidationBuilder) SetPassword(password string) *ValidationBuilder {
	v.password.value = password
	v.password.validator = func(a any) bool {
		passwordRegex := regexp.MustCompile(`^[0-9a-zA-Z]{6,}$`)
		return passwordRegex.MatchString(a.(string))
	}
	return &v
}

func (v ValidationBuilder) SetUsername(username string) *ValidationBuilder {
	v.username.value = username
	v.username.validator = func(a any) bool {
		return len(a.(string)) > 6
	}
	return &v
}

func (v ValidationBuilder) SetIds(ids ...uint) *ValidationBuilder {
	v.id.value = ids
	v.id.validator = func(a any) bool {
		set := a.([]uint)

		for _, item := range set {
			if item <= 0 {
				return false
			}
		}
		return true
	}

	return &v
}

func (v ValidationBuilder) Validate() (bool, error) {
	if v.email.value != nil {
		if ok := v.email.validator(v.email.value); !ok {
			return false, errorext.NewValidationError("email is invalid")
		}
	}

	if v.password.value != nil {
		if ok := v.password.validator(v.password.value); !ok {
			return false, errorext.NewValidationError("password is invalid")
		}
	}

	if v.username.value != nil {
		if ok := v.username.validator(v.username.value); !ok {
			return false, errorext.NewValidationError("username is invalid")
		}
	}

	if v.id.value != nil {
		if ok := v.id.validator(v.id.validator); !ok {
			return false, errorext.NewValidationError("id is invalid")
		}
	}

	return true, nil
}
