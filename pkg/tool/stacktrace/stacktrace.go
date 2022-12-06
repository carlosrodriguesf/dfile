package stacktrace

import (
	"fmt"
	"runtime"
	"strings"
)

type StackEntry struct {
	function string
	file     string
	line     int
}

func (l StackEntry) String() string {
	return fmt.Sprintf("%s:%d (%s)", l.file, l.line, l.function)
}

type StackTrace []StackEntry

func (s StackTrace) WriteOnStringBuilder(builder *strings.Builder, prefix string) {
	for _, stackEntry := range s {
		builder.WriteRune('\n')
		builder.WriteString(prefix)
		builder.WriteString("at ")
		builder.WriteString(stackEntry.String())
	}
}

func getStackTrace(skipCalls int) StackTrace {
	var stackTrace []StackEntry
	for ; ; skipCalls++ {
		stackEntry := getStackEntry(skipCalls)
		if stackEntry == nil {
			break
		}
		stackTrace = append(stackTrace, *stackEntry)
	}
	return stackTrace
}

func getStackEntry(skipCalls int) *StackEntry {
	pc, file, line, ok := runtime.Caller(skipCalls)
	if !ok {
		return nil
	}
	fn := getFuncName(pc)
	return &StackEntry{
		function: fn,
		file:     file,
		line:     line,
	}
}

func getFuncName(pc uintptr) string {
	fn := runtime.FuncForPC(pc).Name()
	parts := strings.Split(fn, "/")
	return parts[len(parts)-1]
}
