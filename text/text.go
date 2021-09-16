package text

import (
	"fmt"
)

type text struct {
	value string
	field []interface{}
	temp  map[string]interface{}
}

func (t *text) Fields(field []interface{}) *text {
	t.setTemp("field", append(t.field, field...))
	return t
}

func (t *text) EOF() *text {
	t.setTemp("eof", "\n")
	return t
}

func (t *text) setTemp(key string, value interface{}) {
	if t.temp == nil {
		t.resetTemp()
		t.temp = make(map[string]interface{})
	}
	t.temp[key] = value
}

func (t *text) resetTemp() {
	t.temp = make(map[string]interface{})
}

func (t text) ToValue() string {
	return t.value
}

func (t text) ToString() string {
	value := t.value
	if val, ok := t.temp["eof"].(string); ok && val != "" {
		value += val
	}
	fields := t.field
	if val, ok := t.temp["field"].([]interface{}); ok && val != nil {
		fields = val
	}
	return fmt.Sprintf(value, fields...)
}
