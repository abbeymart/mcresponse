// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect standard transaction response types

package mcresponse

type ResponseMessage struct {
	Code       string
	ResCode    uint32		// Standard Http Code
	ResMessage string		// Standard Http-Code Text
	Message    string
	Value      interface{}
}

type ResponseMessageOptions struct {
	Message string
	Value   interface{}
}

type MessageParam map[string]ResponseMessage
