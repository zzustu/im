package view

const (
	successCode = 0 // 成功业务状态码
	failedCode  = 1 // 失败业务状态码
)

// 前端响应结果集
type result struct {
	Code    int         `json:"code"`    // 业务状态码
	Message string      `json:"message"` // 错误消息
	Data    interface{} `json:"data"`    // 响应数据
}

// 新建一个成功的消息报文
//
// 参数:
// 		- message: 错误提示信息
func NewFailedResult(message string) *result {
	return &result{
		Code:    failedCode,
		Message: message,
		Data:    nil,
	}
}

// 新建一个成功的消息报文
//
// 参数:
// 		- data: 响应数据
func NewSuccessResult(data interface{}) *result {
	return &result{
		Code:    successCode,
		Message: "",
		Data:    data,
	}
}
