package response

import (
	"Toolkit-GO/constants"
)

// 错误定义步骤
//  1. 先在下方定义自己服务的错误码code区间， 各自的服务请用不同的 const() 段区分
//  2. 请按格式写清楚注释 格式如下 （不要用 /**/的注释)
//     // 错误代码 注释
//     用上面格式可以保证IDE在代码提示时显示注释
//  3. 在errorMap中定义具体的错误对象, 包括code，desc，msg

// 通用区间错误码
const (
	// Success 成功
	Success constants.ErrorCodeType = "0000"
	// Error 读取body体数据错误
	Failed             constants.ErrorCodeType = "9999"
	ReadBoydError      constants.ErrorCodeType = "0001"
	JsonBoydError      constants.ErrorCodeType = "0002"
	NotEnough          constants.ErrorCodeType = "0003"
	DBWgError          constants.ErrorCodeType = "0004"
	UpdateAccountError constants.ErrorCodeType = "0005"
	// flow bill response code
	CreateTagSuccess constants.ErrorCodeType = "1000"
	CreateTagError   constants.ErrorCodeType = "1001"
)

////////////////////////////////

// 然后定义具体错误对象
var errorMap = map[constants.ErrorCodeType]*Error{
	// 通用错误
	Success:            NewErrorDefault(Success, "成功"),
	Failed:             NewErrorDefault(Failed, "未知失败,系统异常"),
	ReadBoydError:      NewErrorDefault(ReadBoydError, "读取boyd数据错误"),
	JsonBoydError:      NewErrorDefault(ReadBoydError, "解析boyd数据错误"),
	NotEnough:          NewErrorDefault(NotEnough, "可用资金不足"),
	DBWgError:          NewErrorDefault(DBWgError, "数据库网关错误"),
	UpdateAccountError: NewErrorDefault(UpdateAccountError, "更新账户信息失败"),
	// flow bill response code
	CreateTagSuccess: NewErrorDefault(CreateTagSuccess, "创建tag成功"),
	CreateTagError:   NewErrorDefault(CreateTagError, "创建Tag失败,请联系管理员"),
}

// GetErrorWithErrorCode 通过错误代码获取对应的错误
func GetErrorWithErrorCode(code constants.ErrorCodeType) *Error {
	if err, ok := errorMap[code]; ok {
		return err
	}
	return errorMap[Failed]
}

func GetError(result bool, code constants.ErrorCodeType) *Response {
	resp := NewResponseDefault()
	resp.Result = result
	resp.AppendError(NewErrorWithErrorCode(code))
	return resp
}

func GetRowsResponse(result bool, code constants.ErrorCodeType, total int, rows interface{}) *Response {
	resp := NewResponseDefault()
	resp.Result = result
	resp.AppendError(NewErrorWithErrorCode(code))
	resp.Total = total
	resp.WithRows(rows)
	return resp

}
