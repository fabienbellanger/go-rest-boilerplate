package templates

import (
	"html/template"
	"strings"
)

// TemplateFuncMap lists functions
var TemplateFuncMap = template.FuncMap{
	"getHttpCodeClass":   GetHttpCodeClass,
	"getHttpMethodClass": GetHttpMethodClass,
	"addNoBreakspace":    AddNoBreakspace,
}

// GetHttpCodeClass displays right CSS class in function of HTTP code
func GetHttpCodeClass(code int) string {
	if code >= 200 && code < 300 {
		return "success"
	} else if code >= 500 && code < 600 {
		return "danger"
	}

	return "warning"
}

// GetHttpMethodClass displays right CSS class in function of HTTP method
func GetHttpMethodClass(method string) string {
	switch method {
	case "GET":
		return "success"
	case "POST":
		return "warning"
	case "PATCH":
		return "primary"
	case "PUT":
		return "info"
	case "DELETE":
		return "danger"
	default:
		return "secondary"
	}
}

// AddNoBreakspace adds &nbsp; instead of spaces
func AddNoBreakspace(s string) template.HTML {
	return template.HTML(strings.ReplaceAll(s, " ", "&nbsp;"))
}
