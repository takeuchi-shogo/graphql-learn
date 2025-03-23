package errors

import "errors"

// エラーメッセージ
const (
	ErrMessageUserNotFoundText = "ユーザーが見つかりません"
)

// エラー変数
var (
	ErrMessageUserNotFound = errors.New(ErrMessageUserNotFoundText)
)
