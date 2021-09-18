package validation

import (
	"fmt"
	"reflect"
)

type min struct {
	value   interface{}
	message string
}

func Min(value interface{}) *min {
	return &min{value: value}
}

func (m min) Validate(value interface{}) bool {
	return compareVerify(reflect.ValueOf(value), fmt.Sprintf("lt=%d", m.value))
}

func (m *min) SetMessage(message string) ValidateRule {
	m.message = message
	return m
}

func (m *min) Message() string {
	return m.message
}
