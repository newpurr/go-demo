package gogenerate

// ErrCode 表示错误码
type ErrCode int

// https://pkg.go.dev/golang.org/x/tools@v0.1.4/cmd/stringer
// https://darjun.github.io/2019/08/21/golang-generate/
//go:generate stringer -type ErrCode -linecomment
// 定义错误码
const (
	ErrCodeOk            ErrCode = 0 // OK
	ErrCodeInvalidParams ErrCode = 1 // 无效参数
	ErrCodeTimeout       ErrCode = 2 // 超时
)
