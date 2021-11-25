//go:generate stringer -type ErroCode -linecomment

package code

type ErrCode int

const (
	OK ErrCode = 0 // 成功
)
