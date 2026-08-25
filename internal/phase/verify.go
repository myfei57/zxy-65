package phase

import "errors"

func (v *Verifier) Verify() error {
	if !v.IsForward() {
		return errors.New("reversed phase sequence")
	}
	return nil
}
