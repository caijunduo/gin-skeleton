package validation

import (
	"fmt"
	"reflect"
)

type max struct {
	value   interface{}
	message string
}

func Max(value interface{}) *max {
	return &max{value: value}
}

func (m max) Validate(value interface{}) bool {
	return compareVerify(reflect.ValueOf(value), fmt.Sprintf("gt=%d", m.value))
}

func (m *max) SetMessage(message string) ValidateRule {
	m.message = message
	return m
}

func (m *max) Message() string {
	return m.message
}
