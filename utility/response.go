package utility

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// JsonResponse 统一响应结构
type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(r *ghttp.Request, data ...interface{}) {
	resp := JsonResponse{
		Code:    0,
		Message: "success",
	}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	r.Response.WriteJsonExit(resp)
}

// Error 错误响应
func Error(r *ghttp.Request, code int, message string) {
	r.Response.WriteJsonExit(JsonResponse{
		Code:    code,
		Message: message,
	})
}

// ErrorMsg 简单错误响应
func ErrorMsg(r *ghttp.Request, message string) {
	Error(r, -1, message)
}

// PageResult 分页结果
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// SuccessPage 分页成功响应
func SuccessPage(r *ghttp.Request, list interface{}, total, page, pageSize int) {
	Success(r, PageResult{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
