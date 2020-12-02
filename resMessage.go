// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcresponse

func msgFunc(code string, resCode uint32, resMessage string, msg string, value interface{}) ResponseMessage {
	return ResponseMessage{
		Code:       code,
		ResCode:    resCode,
		ResMessage: resMessage,
		Value:      value,
		Message:    msg,
	}
}

