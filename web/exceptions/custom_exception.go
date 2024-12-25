package exceptions

import (
	"fmt"
	"runtime"
	"strings"
)

type ErrorCode struct {
	Code    int
	Message string
}

// CustomError 定义自定义错误类型
type CustomError struct {
	Code    int
	Message string
	Stack   string
}

// 实现 error 接口
func (e *CustomError) Error() string {
	return fmt.Sprintf("%s\nStack trace:\n%s", e.Message, e.Stack)
}

const DefaultCode int = 0

// Of 创建一个新的自定义错误
func NewErrorCode(code int, message string) *ErrorCode {
	return &ErrorCode{
		Code:    code,
		Message: message,
	}
}

func OfCode(errCode *ErrorCode) *CustomError {
	return &CustomError{
		Code:    errCode.Code,
		Message: errCode.Message,
		Stack:   captureStackTrace(),
	}
}

// Of 创建一个新的自定义错误
func Of(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Stack:   captureStackTrace(),
	}
}

func OfMessage(message string) *CustomError {
	return &CustomError{
		Code:    DefaultCode,
		Message: message,
		Stack:   captureStackTrace(),
	}
}

func captureStackTrace() string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	count := 0

	var sb strings.Builder
	for _, pc := range pcs[:n] {
		count++
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		sb.WriteString(fmt.Sprintf("%s:%d\n", file, line))
		if count > 10 {
			// 堆栈信息只收集10行，多了没必要
			break
		}
	}
	return sb.String()
}

func (e *CustomError) GetI18nMsgTemplate() string {
	if e.Code == DefaultCode {
		return e.Message
	}
	if e.Message == "" {
		return fmt.Sprintf("%d", e.Code)
	}
	return fmt.Sprintf("${err.%d:%s}", e.Code, e.Message)
}

func TryCatch(try func() (interface{}, error), catch func(error), finally func()) interface{} {
	// 默认 catch 和 finally 为无操作函数
	if catch == nil {
		catch = func(err error) {}
	}
	if finally == nil {
		finally = func() {}
	}

	// 使用 defer 确保 finally 被执行
	defer finally()

	// 使用 defer 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			// 如果发生 panic，调用 catch 来处理 panic
			catch(fmt.Errorf("panic caught: %v", r))
		}
	}()

	// 执行 try 代码块
	data, err := try()
	if err != nil {
		// 如果有错误，调用 catch 块处理错误
		catch(err)
		return nil
	}

	// 没有错误时，返回结果
	return data
}
