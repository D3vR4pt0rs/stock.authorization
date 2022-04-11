package storage

import "errors"

var AccountExistError = errors.New("account exist")
var AccountNotFoundError = errors.New("account didn't exist")
var WrongPasswordError = errors.New("password is wrong")
var InternalError = errors.New("internal error")
