package slog

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Info(businessid int64, uname string, mt string, ft string, other ...interface{}) string {
	str := ""
	if len(other) != 0 {
		str = fmt.Sprintf("%s %s %s了%s %s\n", time.Now().Format("2006-01-02 15:04:05"), uname, mt, ft, other)
	} else {
		str = fmt.Sprintf("%s %s %s了%s\n", time.Now().Format("2006-01-02 15:04:05"), uname, mt, ft)
	}
	_, file, line, _ := runtime.Caller(1)
	Cp.Info().
		Str("file", file).
		Int("line", line).
		Int64("business_id", businessid).
		Str("name", uname).
		Str("method", mt).
		Str("func", ft).
		Send()
	return str
}

var (
	Cp    = log.Logger
	Login = log.Logger
)

func init() {
	file := "./log/cp.log"
	fi, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Cp = zerolog.New(fi).With().Timestamp().Logger()

	fiLogin := "./log/login.log"
	fLogin, err := os.OpenFile(fiLogin, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Login = zerolog.New(fLogin).With().Timestamp().Logger()
}

func CpInfo(str string, data interface{}) {
	buf, _ := json.Marshal(data)
	_, file, line, _ := runtime.Caller(1)
	Cp.Info().
		Str("file", file).
		Int("line", line).
		Str("msg", str).
		Msg(string(buf))
}

func CpWarn(str string, err error) {
	_, file, line, _ := runtime.Caller(1)
	Cp.Warn().
		Str("file", file).
		Int("line", line).
		Str("msg", str).
		Msg(err.Error())
}

func CpError(str string, err error, args ...interface{}) {
	_, file1, line1, _ := runtime.Caller(1)
	fs := strings.Split(file1, "/")
	e := Cp.Error().
		Str("file", strings.Join(fs[len(fs)-3:], "/")).
		Int("line", line1)

	e.Str("msg", str)
	e.Str("msg", err.Error())
	e.Msg(fmt.Sprint(args...))
}
