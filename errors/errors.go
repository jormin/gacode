package errors

import "errors"

// MissingRequiredArgumentErr 缺少必传参数
var MissingRequiredArgumentErr = errors.New("缺少必传参数")

// AccountNameExistsErr 账户名称已存在
var AccountNameExistsErr = errors.New("账户名称已存在")

// AccountNameNotExistsErr 账户名称不存在
var AccountNameNotExistsErr = errors.New("账户名称不存在")

// GenerateSecretErr 生成账户秘钥失败
var GenerateSecretErr = errors.New("生成账户秘钥失败")

// FlagContentValidateErr flag content validate error
var FlagContentValidateErr = errors.New("flag content validate error")

// TodoNotExistsErr todo not exists
var TodoNotExistsErr = errors.New("todo not exists")
