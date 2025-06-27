// Define Global Constants HTTP header name
package constants

import "net/textproto"

var (
	XServiceName  = textproto.CanonicalMIMEHeaderKey("x-service-name")
	XAPIKey       = textproto.CanonicalMIMEHeaderKey("x-api-key")
	XRequestAt    = textproto.CanonicalMIMEHeaderKey("x-request-at")
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")
)
