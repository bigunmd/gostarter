package heroes

import "errors"

// ErrAlreadyExists returned if requested resource is already exists.
var ErrAlreadyExists error = errors.New("already exists")
