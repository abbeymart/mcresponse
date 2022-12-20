// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect standard transaction response functions | v0.9.1

package mcresponse

import "fmt"

func msgFunc(code string, resCode int, resMessage string, msg string, value interface{}) ResponseMessage {
	fmt.Printf("msgCode: %v\n", code)
	return ResponseMessage{
		Code:       code,
		ResCode:    resCode,
		ResMessage: resMessage,
		Value:      value,
		Message:    msg,
	}
}

func GetResMessage(msgCode string, options ResponseMessageOptions) ResponseMessage {
	var value interface{} = nil
	code := msgCode
	resCode := OK
	resMessage := ""
	message := options.Message
	// compose response-Message: for known/standard code/messageParam OR unknown/user-specified-code/message/value
	val, ok := StdResMessages[code]
	if ok {
		code = val.Code
		resCode = val.ResCode
		resMessage = val.ResMessage
		message = val.Message
		value = val.Value
		// update option-values: Message && Value
		if options.Value != nil {
			// set value to optional value
			value = options.Value
		}
		if options.Message != "" {
			// set message to optional message
			message = options.Message
		}
	} else {
		defaultVal, dOk := StdResMessages["unknown"]
		if dOk {
			resCode = defaultVal.ResCode
			resMessage = defaultVal.ResMessage
			// update Code, Value and Message
			if msgCode != "" {
				// set value to msgCode, as specified
				code = msgCode
			} else {
				code = defaultVal.Code
			}
			if options.Value != nil {
				// set value to optional value
				value = options.Value
			} else {
				value = defaultVal.Value
			}
			if options.Message != "" {
				// set message to optional message
				message = options.Message
			} else {
				message = defaultVal.Message
			}
		} else {
			code = msgCode
			resCode = 404
			resMessage = options.Message
			message = options.Message
			value = options.Value
		}
	}

	return msgFunc(code, resCode, resMessage, message, value)
}
