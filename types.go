// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcresponse

type ResponseMessage struct {
	Code       string
	ResCode    uint32
	ResMessage string
	Message    string
	Value      interface{}
}

type ResponseMessageOptions struct {
	code string
	resCode uint32
	resMessage string
	message string
	value interface{}
}

type MessageParam map[string]ResponseMessage
