package errors

import "fmt"

func MultiError(errs []error, label string) error {
	if len(errs) > 0 {
		err := errs[0]
		for _, err_ := range errs[1:] {
			err = fmt.Errorf("%w\n%s", err, err_)
		}
		err = fmt.Errorf("%s: \n%w", label, err)
		return err
	} else {
		return nil
	}
}
