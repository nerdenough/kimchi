package replace

import "strings"

// APIRequest replaces string tokens based on the content type and body.
func APIRequest(str, contentType string, body []byte) string {
	switch contentType {
	case "text/plain":
		return strings.ReplaceAll(str, "{0}", string(body))
	default:
		return str
	}
}
