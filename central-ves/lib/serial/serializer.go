package serial

import "fmt"

type CodeType = int

type ErrorSerializer struct {
	Code CodeType `json:"code"`
	Err  string   `json:"error"`
}

func (e ErrorSerializer) Error() string {
	return fmt.Sprintf("<code:%v,err:%v>", e.Code, e.Err)
}

type Response struct {
	Code CodeType `json:"code"`
}
