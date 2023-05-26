package slog

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Sys = log.Logger
)

func init() {
	file := "./log/sys.log"
	fi, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Sys = zerolog.New(fi)
	Sys = Sys.With().
		Timestamp().
		Logger()
}

func Warn(str string, err error) {
	_, file, line, _ := runtime.Caller(1)
	Sys.Warn().
		Str("file", file).
		Int("line", line).
		Str("msg", str).
		Msg(err.Error())
}

func Error(str string, err error, args ...interface{}) {
	_, file1, line1, _ := runtime.Caller(1)
	fs := strings.Split(file1, "/")
	e := Sys.Error().
		Str("file", strings.Join(fs[len(fs)-3:], "/")).
		Int("line", line1)

	e.Str("msg", str)
	e.Str("msg", err.Error())
	e.Msg(fmt.Sprint(args...))
}
