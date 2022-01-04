// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect standard transaction response functions

package mcresponse

func msgFunc(code string, resCode int, resMessage string, msg string, value interface{}) ResponseMessage {
	return ResponseMessage{
		Code:       code,
		ResCode:    resCode,
		ResMessage: resMessage,
		Value:      value,
		Message:    msg,
	}
}

func GetResMessage(msgCode string, options ResponseMessageOptions) ResponseMessage {
	var (
		value      interface{} = nil
		code                   = "unknown"
		resCode                = UnprocessableEntity
		resMessage             = ""
		message                = "Unknown/Unspecified response message"
	)
	// compose response-Message
	if val, ok := StdResMessages[msgCode]; ok {
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
			// append optional message
			//message += " | "
			//message += options.Message
		}
	} else {
		if val, ok := StdResMessages["unknown"]; ok {
			code = val.Code
			resCode = val.ResCode
			resMessage = val.ResMessage
			message = val.Message
			value = val.Value
			// update msgCode and option-values: Message && Value
			if msgCode != "" {
				// set value to msgCode, as specified
				code = msgCode
			}
			if options.Value != nil {
				// set value to optional value
				value = options.Value
			}
			if options.Message != "" {
				// set message to optional message
				message = options.Message
			}
		}
	}

	return msgFunc(code, resCode, resMessage, message, value)
}
