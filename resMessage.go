// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect standard transaction response functions | v0.9.1

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
	var value interface{} = nil
	code := "unknown"
	resCode := OK
	resMessage := ""
	message := "Unknown/UnAuthorized action"
	// compose response-Message: for known/standard code/messageParam OR unknown/user-specified-code/message/value
	val, ok := StdResMessages[msgCode]
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
			code = defaultVal.Code
			resCode = defaultVal.ResCode
			resMessage = defaultVal.ResMessage
			message = defaultVal.Message
			value = defaultVal.Value
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
