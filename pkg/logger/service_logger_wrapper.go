//go:build !darwin

package logger

func NewNezhaServiceLogger(s interface{}, errs chan<- error) (interface{}, error) {
	// 返回一个空实现或控制台记录器
	return struct{}{}, nil
}
