package validation

type field struct {
	value interface{}
	rules []ValidateRule
}

func Field(value interface{}) *field {
	return &field{value: value}
}

func (f *field) Rule(rule ValidateRule) *field {
	f.rules = append(f.rules, rule)
	return f
}
