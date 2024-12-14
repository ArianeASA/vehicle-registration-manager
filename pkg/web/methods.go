package web

import (
	"net/http"
	"strings"
)

func VerifyMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if !strings.EqualFold(method, r.Method) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return true
	}
	return false
}
