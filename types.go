// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect standard transaction response types

package mcresponse

type ResponseMessage struct {
	Code       string      `json:"code"`
	ResCode    int         `json:"res_code"`    // Standard Http Code
	ResMessage string      `json:"res_message"` // Standard Http-Code Text
	Message    string      `json:"message"`
	Value      interface{} `json:"value"`
}

type ResponseMessageOptions struct {
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

type MessageParam map[string]ResponseMessage
