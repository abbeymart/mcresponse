// @Author: abbeymart | Abi Akindele | @Created: 2020-11-30 | @Updated: 2020-11-30
// @Company: mConnect.biz | @License: MIT
// @Description: mcconfig: standard network codes and texts

package mcresponse

const (
	// Continue /** RFC 7231, 6.2.1 */
	Continue = 100
	// SwitchingProtocols /** RFC 7231, 6.2.2 */
	SwitchingProtocols = 101
	// Processing /** RFC 2518, 10.1 */
	Processing = 102
	// EarlyHints /** RFC 8297 **/
	EarlyHints = 103
	// OK /** RFC 7231, 6.3.1 */
	OK = 200
	// Created /** RFC 7231, 6.3.2 */
	Created = 201
	// Accepted /** RFC 7231, 6.3.3 */
	Accepted = 202
	// NonAuthoritativeInfo /** RFC 7231, 6.3.4 */
	NonAuthoritativeInfo = 203
	// NoContent /** RFC 7231, 6.3.5 */
	NoContent = 204
	// ResetContent /** RFC 7231, 6.3.6 */
	ResetContent = 205
	// PartialContent /** RFC 7233, 4.1 */
	PartialContent = 206
	// MultiStatus /** RFC 4918, 11.1 */
	MultiStatus = 207
	// AlreadyReported /** RFC 5842, 7.1 */
	AlreadyReported = 208
	// IMUsed /** RFC 3229, 10.4.1 */
	IMUsed = 226
	// MultipleChoices /** RFC 7231, 6.4.1 */
	MultipleChoices = 300
	// MovedPermanently /** RFC 7231, 6.4.2 */
	MovedPermanently = 301
	// Found /** RFC 7231, 6.4.3 */
	Found = 302
	// SeeOther /** RFC 7231, 6.4.4 */
	SeeOther = 303
	// NotModified /** RFC 7232, 4.1 */
	NotModified = 304
	// UseProxy /** RFC 7231, 6.4.5 */
	UseProxy = 305
	// TemporaryRedirect /** RFC 7231, 6.4.7 */
	TemporaryRedirect = 307
	// PermanentRedirect /** RFC 7538, 3 */
	PermanentRedirect = 308
	// BadRequest /** RFC 7231, 6.5.1 */
	BadRequest = 400
	// UnAuthorized /** RFC 7235, 3.1 */
	UnAuthorized = 401
	// PaymentRequired /** RFC 7231, 6.5.2 */
	PaymentRequired = 402
	// Forbidden /** RFC 7231, 6.5.3 */
	Forbidden = 403
	// NotFound /** RFC 7231, 6.5.4 */
	NotFound = 404
	// MethodNotAllowed /** RFC 7231, 6.5.5 */
	MethodNotAllowed = 405
	// NotAcceptable /** RFC 7231, 6.5.6 */
	NotAcceptable = 406
	// ProxyAuthRequired /** RFC 7235, 3.2 */
	ProxyAuthRequired = 407
	// RequestTimeout /** RFC 7231, 6.5.7 */
	RequestTimeout = 408
	// Conflict /** RFC 7231, 6.5.8 */
	Conflict = 409
	// Gone /** RFC 7231, 6.5.9 */
	Gone = 410
	// LengthRequired /** RFC 7231, 6.5.10 */
	LengthRequired = 411
	// PreconditionFailed /** RFC 7232, 4.2 */
	PreconditionFailed = 412
	// RequestEntityTooLarge /** RFC 7231, 6.5.11 */
	RequestEntityTooLarge = 413
	// RequestURITooLong /** RFC 7231, 6.5.12 */
	RequestURITooLong = 414
	// UnsupportedMediaType /** RFC 7231, 6.5.13 */
	UnsupportedMediaType = 415
	// RequestedRangeNotSatisfiable /** RFC 7233, 4.4 */
	RequestedRangeNotSatisfiable = 416
	// ExpectationFailed /** RFC 7231, 6.5.14 */
	ExpectationFailed = 417
	// Teapot /** RFC 7168, 2.3.3 */
	Teapot = 418
	// MisdirectedRequest /** RFC 7540, 9.1.2 */
	MisdirectedRequest = 421
	// UnprocessableEntity /** RFC 4918, 11.2 */
	UnprocessableEntity = 422
	// Locked /** RFC 4918, 11.3 */
	Locked = 423
	// FailedDependency /** RFC 4918, 11.4 */
	FailedDependency = 424
	// TooEarly /** RFC 8470, 5.2 */
	TooEarly = 425
	// UpgradeRequired /** RFC 7231, 6.5.15 */
	UpgradeRequired = 426
	// PreconditionRequired /** RFC 6585, 3 */
	PreconditionRequired = 428
	// TooManyRequests /** RFC 6585, 4 */
	TooManyRequests = 429
	// RequestHeaderFieldsTooLarge /** RFC 6585, 5 */
	RequestHeaderFieldsTooLarge = 431
	// UnavailableForLegalReasons /** RFC 7725, 3 */
	UnavailableForLegalReasons = 451

	// InternalServerError /** RFC 7231, 6.6.1 */
	InternalServerError = 500
	// NotImplemented /** RFC 7231, 6.6.2 */
	NotImplemented = 501
	// BadGateway /** RFC 7231, 6.6.3 */
	BadGateway = 502
	// ServiceUnavailable /** RFC 7231, 6.6.4 */
	ServiceUnavailable = 503
	// GatewayTimeout /** RFC 7231, 6.6.5 */
	GatewayTimeout = 504
	// HTTPVersionNotSupported /** RFC 7231, 6.6.6 */
	HTTPVersionNotSupported = 505
	// VariantAlsoNegotiates /** RFC 2295, 8.1 */
	VariantAlsoNegotiates = 506
	// InsufficientStorage /** RFC 4918, 11.5 */
	InsufficientStorage = 507
	// LoopDetected /** RFC 5842, 7.2 */
	LoopDetected = 508
	// NotExtended /** RFC 2774, 7 */
	NotExtended = 510
	// NetworkAuthenticationRequired /** RFC 6585, 6 */
	NetworkAuthenticationRequired = 511
)

var StatusText = map[int]string{
	Continue:                      "Continue",
	SwitchingProtocols:            "Switching Protocols",
	Processing:                    "Processing",
	EarlyHints:                    "Early Hints",
	OK:                            "OK",
	Created:                       "Created",
	Accepted:                      "Accepted",
	NonAuthoritativeInfo:          "Non-Authoritative Information",
	NoContent:                     "No Content",
	ResetContent:                  "Reset Content",
	PartialContent:                "Partial Content",
	MultiStatus:                   "Multi-Status",
	AlreadyReported:               "Already Reported",
	IMUsed:                        "IM Used",
	MultipleChoices:               "Multiple Choices",
	MovedPermanently:              "Moved Permanently",
	Found:                         "Found",
	SeeOther:                      "See Other",
	NotModified:                   "Not Modified",
	UseProxy:                      "Use Proxy",
	TemporaryRedirect:             "Temporary Redirect",
	PermanentRedirect:             "Permanent Redirect",
	BadRequest:                    "Bad Request",
	UnAuthorized:                  "UnAuthorized",
	PaymentRequired:               "Payment Required",
	Forbidden:                     "Forbidden",
	NotFound:                      "Not Found",
	MethodNotAllowed:              "Method Not Allowed",
	NotAcceptable:                 "Not Acceptable",
	ProxyAuthRequired:             "Proxy Authentication Required",
	RequestTimeout:                "Request Timeout",
	Conflict:                      "Conflict",
	Gone:                          "Gone",
	LengthRequired:                "Length Required",
	PreconditionFailed:            "Precondition Failed",
	RequestEntityTooLarge:         "Request Entity Too Large",
	RequestURITooLong:             "Request URI Too Long",
	UnsupportedMediaType:          "Unsupported Media Type",
	RequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable",
	ExpectationFailed:             "Expectation Failed",
	Teapot:                        "I'm a teapot",
	MisdirectedRequest:            "Misdirected Request",
	UnprocessableEntity:           "Unprocessable Entity",
	Locked:                        "Locked",
	FailedDependency:              "Failed Dependency",
	TooEarly:                      "Too Early",
	UpgradeRequired:               "Upgrade Required",
	PreconditionRequired:          "Precondition Required",
	TooManyRequests:               "Too Many Requests",
	RequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	UnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	InternalServerError:           "Internal Server Error",
	NotImplemented:                "Not Implemented",
	BadGateway:                    "Bad Gateway",
	ServiceUnavailable:            "Service Unavailable",
	GatewayTimeout:                "Gateway Timeout",
	HTTPVersionNotSupported:       "HTTP Version Not Supported",
	VariantAlsoNegotiates:         "Variant Also Negotiates",
	InsufficientStorage:           "Insufficient Storage",
	LoopDetected:                  "Loop Detected",
	NotExtended:                   "Not Extended",
	NetworkAuthenticationRequired: "Network Authentication Required",
}
