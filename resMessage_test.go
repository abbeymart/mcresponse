// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01, 2026-06-04
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect standard transaction response testing | v0.9.1

package mcresponse

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/abbeymart/mcresponse/messagecodes"
)
import "github.com/abbeymart/mctest"

func TestResMessage(t *testing.T) {
	// test-data
	msgType := messagecodes.Success
	msgType2 := messagecodes.CheckError
	msgType3 := "custom"
	msg3 := "Custom Message"
	options := ResponseMessageOptions{
		Message: "",
		Value:   []string{"a", "b", "c"},
	}
	options2 := ResponseMessageOptions{
		Message: "",
		Value:   "",
	}
	options3 := ResponseMessageOptions{
		Message: msg3,
		Value:   "Custom",
	}
	res := ResponseMessage{
		Code:       messagecodes.Success,
		ResCode:    200,
		ResMessage: "OK",
		Value:      "",
		Message:    "Request completed successfully",
	}
	res2 := ResponseMessage{
		Code:       messagecodes.ParamsError,
		ResCode:    406,
		ResMessage: "Not Acceptable",
		Value:      "",
		Message:    "Parameters checking error",
	}
	res3 := ResponseMessage{
		Code:       "custom",
		ResCode:    200,
		ResMessage: "OK",
		Value:      "Custom",
		Message:    "Custom Message",
	}
	// Test cases

	var results []mctest.UnitTestResult

	test1 := mctest.NewTest(mctest.ParamsType{
		Name: "should return success code for success-message",
	})
	test1.SetTestFunction(func() {
		req := GetResMessage(msgType, options)
		test1.AssertEquals(req.Code, res.Code, "response-code should be: "+res.Code)
		test1.AssertEquals(req.Message, res.Message, "response-message should be: "+res.Message)
	})
	test1Result := test1.RunTest()
	results = append(results, test1Result)

	fmt.Println("")
	test2 := mctest.NewTest(mctest.ParamsType{
		Name: "should return ok/200 resCode for success-message",
	})
	test2.SetTestFunction(func() {
		req := GetResMessage(msgType, ResponseMessageOptions{})
		test2.AssertEquals(req.ResCode, res.ResCode, "response-code should be: "+strconv.Itoa(res.ResCode))
		test2.AssertEquals(req.Message, res.Message, "response-message should be: "+res.Message)
	})
	test2Result := test2.RunTest()
	results = append(results, test2Result)

	fmt.Println("")
	test3 := mctest.NewTest(mctest.ParamsType{
		Name: "should return Completed successfully message for success-message",
	})
	test3.SetTestFunction(func() {
		req := GetResMessage(msgType, ResponseMessageOptions{})
		test3.AssertEquals(req.Message, res.Message, "response-message should be: "+res.Message)
	})
	test3Result := test3.RunTest()
	results = append(results, test3Result)

	fmt.Println("")
	test4 := mctest.NewTest(mctest.ParamsType{
		Name: "should return correct default message",
	})
	test4.SetTestFunction(func() {
		options := ResponseMessageOptions{
			Value:   []string{"a", "b", "c"},
			Message: "Successful",
		}
		req := GetResMessage(msgType, options)
		test4.AssertEquals(strings.Contains(req.Message, options.Message), true, "response-message should contains: "+options.Message)
	})
	test4Result := test4.RunTest()
	results = append(results, test4Result)

	fmt.Println("")
	test5 := mctest.NewTest(mctest.ParamsType{
		Name: "should return correct custom message",
	})
	test5.SetTestFunction(func() {
		req := GetResMessage(msgType3, options3)
		test5.AssertEquals(req.Code, res3.Code, "response-code should be: "+res3.Code)
		test5.AssertEquals(strings.Contains(req.Message, options3.Message), true, "response-message should contains: "+options3.Message)
	})
	test5Result := test5.RunTest()
	results = append(results, test5Result)

	// check-error test-cases

	fmt.Println("")
	test6 := mctest.NewTest(mctest.ParamsType{
		Name: "should return paramsError code for checkError-message",
	})
	test6.SetTestFunction(func() {
		req := GetResMessage(msgType2, options2)
		test6.AssertEquals(req.Code, res2.Code, "response-code should be: "+res2.Code)
		test6.AssertNotEquals(req.Code, "unAuthorized", "response-code"+req.Code+"should not be: unAuthorized")
	})
	test6Result := test6.RunTest()
	results = append(results, test6Result)

	fmt.Println("")
	test7 := mctest.NewTest(mctest.ParamsType{
		Name: "should return NOT_ACCEPTABLE/406 resCode",
	})
	test7.SetTestFunction(func() {
		req := GetResMessage(msgType2, ResponseMessageOptions{})
		test7.AssertEquals(req.ResCode, res2.ResCode, "response-code should be: "+strconv.Itoa(res2.ResCode))
		test7.AssertEquals(req.ResMessage, res2.ResMessage, "response-message should be: "+res2.ResMessage)
	})
	test7Result := test7.RunTest()
	results = append(results, test7Result)

	fmt.Println("")
	test8 := mctest.NewTest(mctest.ParamsType{
		Name: "should return Parameters checking error message",
	})
	test8.SetTestFunction(func() {
		req := GetResMessage(msgType2, options2)
		test8.AssertEquals(req.Message, res2.Message, "response-code should be: "+res2.Message)
	})
	test8Result := test8.RunTest()
	results = append(results, test8Result)

	fmt.Println("")
	test9 := mctest.NewTest(mctest.ParamsType{
		Name: "should return authCode and auth-code message",
	})
	test9.SetTestFunction(func() {
		req := GetResMessage("authCode", ResponseMessageOptions{
			Message: "auth-code",
			Value:   nil,
		})
		test9.AssertEquals(req.Code, "authCode", "response-code should be: authCode")
		test9.AssertEquals(req.Message, "auth-code", "response-message should be: auth-code")
	})
	test9Result := test9.RunTest()
	results = append(results, test9Result)

	// Summary result for all the test-cases
	fmt.Println("")
	mctest.TestResult(results)
}
