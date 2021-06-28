package gogenerate

// ErrCode 表示错误码
type ErrCode int

//go:generate stringer -type ErrCode -linecomment
// 定义错误码
const (
	ErrCodeOk            ErrCode = 0 // OK
	ErrCodeInvalidParams ErrCode = 1 // 无效参数
	ErrCodeTimeout       ErrCode = 2 // 超时
)
