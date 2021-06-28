package gogenerate

// ErrCode2 表示错误码
type ErrCode2 int

// https://pkg.go.dev/golang.org/x/tools@v0.1.4/cmd/stringer
// https://darjun.github.io/2019/08/21/golang-generate/
//go:generate stringer -type ErrCode2 -linecomment
// 定义错误码
const (
	ErrCode2Ok            ErrCode2 = 0 // OK
	ErrCode2InvalidParams ErrCode2 = 1 // 无效参数
	ErrCode2Timeout       ErrCode2 = 2 // 超时
)
