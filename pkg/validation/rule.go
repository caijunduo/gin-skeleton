package validation

type ValidateRule interface {
	Validate(value interface{}) bool
}

type GetMessageRule interface {
	Message() string
}

type SetMessageRule interface {
	SetMessage(message string) ValidateRule
}
