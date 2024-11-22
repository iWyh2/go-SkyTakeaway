package model

// Result 后端统一返回结果
type Result[T any] struct {
	// 编码，1成功，0和其它数字为失败
	Code int `json:"code"`
	// 错误信息
	Msg string `json:"msg"`
	// 数据，任意类型
	Data T `json:"data"`
}

// Success 返回成功结果（不带数据）
func (r *Result[T]) Success() *Result[T] {
	r.Code = 1
	return r
}

// SuccessByData 返回成功结果（带数据）
func (r *Result[T]) SuccessByData(data T) *Result[T] {
	r.Code = 1
	r.Data = data
	return r
}

// Error 返回失败结果（带错误信息）
func (r *Result[T]) Error(msg string) *Result[T] {
	r.Code = 0
	r.Msg = msg
	return r
}
