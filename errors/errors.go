package errors

import "errors"

// MissingRequiredArgumentErr missing required argument
var MissingRequiredArgumentErr = errors.New("missing required argument")

// AccountNameExistsErr account name exists
var AccountNameExistsErr = errors.New("account name exists")

// AccountNameNotExistsErr account name not exists
var AccountNameNotExistsErr = errors.New("account name not exists")

// GenerateSecretErr generate account secret error
var GenerateSecretErr = errors.New("generate account secret error")

// InvalidExportPathErr invalid export path
var InvalidExportPathErr = errors.New("invalid export path")
