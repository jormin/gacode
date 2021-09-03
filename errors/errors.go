package errors

import "errors"

// ErrMissingRequiredArgument missing required argument
var ErrMissingRequiredArgument = errors.New("missing required argument")

// ErrAccountNameExists account name exists
var ErrAccountNameExists = errors.New("account name exists")

// ErrAccountNameNotExists account name not exists
var ErrAccountNameNotExists = errors.New("account name not exists")

// ErrGenerateSecret generate account secret error
var ErrGenerateSecret = errors.New("generate account secret error")

// ErrInvalidExportPath invalid export path
var ErrInvalidExportPath = errors.New("invalid export path")
