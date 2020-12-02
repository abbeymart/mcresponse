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

func GetResMessage(msgType string, options ResponseMessageOptions) ResponseMessage {
	var (
		value      interface{}
		code       string
		resCode    uint32
		resMessage string
		message    string
	)
	// compose response-Message
	if val, ok := StdResMessages[msgType]; ok {
		code = val.Code
		resCode = val.ResCode
		resMessage = val.ResMessage
		message = val.Message
		value = val.Value
		// update option-values: Message && Value
		if options.Value != nil {
			value = options.Value
		}
		if options.Message != "" {
			message += " | "
			message += options.Message
		}
	} else {
		if val, ok := StdResMessages["stdRes"]; ok {
			code = val.Code
			resCode = val.ResCode
			resMessage = val.ResMessage
			message = val.Message
			value = val.Value
			// update option-values: Message && Value
			if options.Value != nil {
				value = options.Value
			}
			if options.Message != "" {
				message += " | "
				message += options.Message
			}
		}
	}

	return msgFunc(code, resCode, resMessage, message, value)
}
