package validation

import (
	"errors"
)

func Verify(fields ...*field) error {
	var err errs
	for _, f := range fields {
		for _, r := range f.rules {
			if r.Validate(f.value) {
				if msg, ok := r.(GetMessageRule); ok {
					err.setError(errors.New(msg.Message()))
				}
				break
			}
		}
	}
	if err.Len() <= 0 {
		return nil
	}
	return err.First()
}
