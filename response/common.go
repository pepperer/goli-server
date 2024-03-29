package response

import "fmt"

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// BuildListResponse 列表构建器

func BuildListResponse(items []interface{}, total uint) Response {
	return Response{
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}


func BuildErrorResponse( err error, ) Response {
	return Response{
		Error:  fmt.Sprint(err),
	}
}

func BuildResponse(item interface{}) Response {
	return Response{
		Data: item,
	}
}
