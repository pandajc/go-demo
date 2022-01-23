package errs

import (
	"fmt"
	pe "github.com/pkg/errors"
)

type BizErr struct {
	code int
	msg  string
}

func (e *BizErr) Error() string {
	return fmt.Sprintf("biz err: code: %d, msg: %s", e.code, e.msg)
}

func (e *BizErr) GetCode() int {
	return e.code
}

func (e *BizErr) GetMsg() string {
	return e.msg
}

func NewBizErrWithStack(code int, msg string) error {
	return pe.WithStack(&BizErr{
		code: code,
		msg:  msg,
	})
}

func NewBizErr(code int, msg string) error {
	return &BizErr{
		code: code,
		msg:  msg,
	}
}

func IsBizErr(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(*BizErr); ok {
		return true
	}
	err = pe.Unwrap(err)
	if err == nil {
		return false
	}
	if _, ok := err.(*BizErr); ok {
		return true
	}
	return false
}
