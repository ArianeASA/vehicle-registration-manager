package logger

import (
	"runtime"
	"strconv"
	"strings"
)

func getCallerInfo() string {
	pc, _, line, ok := runtime.Caller(3)
	if !ok {
		return "unknown:0"
	}

	functionCaller := runtime.FuncForPC(pc)
	if functionCaller == nil {
		return "unknown:0"
	}

	packageFunc := extractNamesCaller(functionCaller.Name())
	return packageFunc + ":" + strconv.Itoa(line)
}

func extractNamesCaller(fullName string) string {
	packageFunc := fullName
	if lastSlash := strings.LastIndex(fullName, "/"); lastSlash != -1 {
		packageFunc = fullName[lastSlash+1:]
	}
	if dotIndex := strings.LastIndex(packageFunc, "."); dotIndex != -1 {
		return packageFunc[:dotIndex] + "." + packageFunc[dotIndex+1:]
	}
	return packageFunc
}
