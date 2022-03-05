package gerr

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	traceStart = 4
	traceEnd   = 64 // let's hope that it will never reach this lol.
)

var separator = ":\n\t"

type Error struct {
	err        error
	grpcCodes  codes.Code
	stack      []uintptr
	stackPrint string // probably put this to err
	query      string
	args       []interface{}
}

func (e *Error) Error() string {
	sb := new(strings.Builder)
	e.writeTraceToSb(sb)
	pad(sb, ": ", e.err.Error())
	if sb.Len() == 0 {
		return "no error"
	}
	e.stackPrint = sb.String()
	fmt.Println(e.stackPrint) // TODO: log this in middleware.
	return e.err.Error()
}

func (e *Error) GRPCStatus() *status.Status {
	return status.New(e.grpcCodes, e.err.Error())
}

func LogErr(err error) string {
	var customErr *Error
	if ok := errors.As(err, &customErr); !ok {
		return err.Error()
	}

	if customErr.query == "" {
		return customErr.Error()
	}

	return fmt.Errorf("query %s failed with args %v, err: %w", customErr.query, customErr.args, customErr.err).Error()
}

// callers returns a list of ptr to the functions that called it.
func callers() []uintptr {
	pc := make([]uintptr, traceEnd)

	n := runtime.Callers(traceStart, pc)
	return pc[:n]
}

// frame returns the nth frame, with frame 0 as the initial function call.
func frame(callers []uintptr, n int) *runtime.Frame {
	frames := runtime.CallersFrames(callers)
	var f runtime.Frame
	var ok bool
	for i := len(callers) - 1; i >= n; i-- {
		f, ok = frames.Next()
		if !ok {
			break
		}
	}
	return &f
}

func populateStack() []uintptr {
	return callers()
}

func (e *Error) writeTraceToSb(sb *strings.Builder) {
	if e.stack == nil {
		return
	}

	printCallers := callers()
	// iterate through e.stack and skip shared frames with printCallers.
	// print function names and lines numbers.
	var prev string
	var diff bool // print and error stack is different
	for i := 0; i < len(e.stack); i++ {
		curFrame := frame(e.stack, i)
		name := curFrame.Func.Name()

		if !diff && i < len(e.stack) {
			if name == frame(printCallers, i).Func.Name() {
				continue
			}
			// don't consider printCallers again.
			diff = true
		}

		// prevent duplication.
		if name == prev {
			continue
		}

		// find uncommon prefix between this and previous name, separation by dots and slashes.
		trim := 0
		for {
			j := strings.IndexAny(name[trim:], "./")
			if j < 0 {
				break
			}
			if !strings.HasPrefix(prev, name[:j+trim]) {
				break
			}
			trim += j + 1
		}

		pad(sb, separator)
		fmt.Fprintf(sb, "%v:%d: ", curFrame.File, curFrame.Line)
		if trim > 0 {
			sb.WriteString("...")
		}
		sb.WriteString(name[trim:])

		prev = name
	}
}

func pad(sb *strings.Builder, strs ...string) {
	if sb.Len() == 0 {
		return
	}
	for _, str := range strs {
		sb.WriteString(str)
	}
}
