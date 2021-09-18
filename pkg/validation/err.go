package validation

type errs struct {
	errs []error
}

func (e *errs) setError(err error) {
	e.errs = append(e.errs, err)
}

func (e errs) Len() int {
	return len(e.errs)
}

func (e errs) First() error {
	if e.Len() <= 0 {
		return nil
	}
	return e.errs[0]
}
