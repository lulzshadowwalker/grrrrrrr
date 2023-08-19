package enum

import (
	"errors"
	"fmt"
)

type Enumer interface {
	// TODO would probably be more efficient to write a generalized 
	// validation method where you pass in an array of valid enums
	
	Validate() error 
}

var ErrInvalidEnum = errors.New("invalid enum value")

type EnumError struct {
	Enum interface{}
	Err error
}

func (e *EnumError) Error() string {
	return	fmt.Sprintf("enum error (%d) %q", e.Enum, ErrInvalidEnum)
}

func (e *EnumError) Unwrap() error {
	return e.Err
}

func IsIn[T comparable](v T, enums[]T) error {
	for _, e := range enums {
		if v == e {
			return nil
		}
	}

	return &EnumError{
		Enum: v,
		Err: ErrInvalidEnum,
	}
}
