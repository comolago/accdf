package interfaces

import (
	"fmt"
)

type ErrHandler struct {
	Number      uintptr
	Component   string
	Function    string
	Description string
}

var errors = [...]string{
	1: "",                                                                       // EXTERNAL_LIBRARY
	2: "You must set elasticsearch URL - for example http://127.0.0.1:9200",     // NOURL
	3: "You must set elasticsearch Index",                                       // NOINDEX
	4: "The specified index does not exist or you are not allowed to access it", // WRONGINDEX
	5: "Elasticsearch query failed",
}

func (e ErrHandler) Error() string {
	var message string
	if 0 <= int(e.Number) && int(e.Number) < len(errors) {
		message = fmt.Sprintf("%s | %s", e.Component, e.Function)
		if int(e.Number) == 1 {
			message = fmt.Sprintf("%s | %s | %s | ", e.Component, e.Function, e.Description)
		} else {
			message = fmt.Sprintf("%s | %s | %s | %s", e.Component, e.Function, errors[e.Number], e.Description)
		}

	} else {
		message = fmt.Sprintf("%s | %s | Undefined error: error number is %d | %s", e.Component, e.Function, e.Number, e.Description)
	}
	return message
}
