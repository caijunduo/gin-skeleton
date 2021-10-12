package validation

import "reflect"

type match struct {
	patten  string
	message string
}

func Match(patten string) *match {
	return &match{patten: patten}
}

func (m match) Validate(value interface{}) bool {
	return !regexpMatch(m.patten, reflect.ValueOf(value).String())
}

func (m *match) SetMessage(message string) ValidateRule {
	m.message = message
	return m
}

func (m *match) Message() string {
	return m.message
}
