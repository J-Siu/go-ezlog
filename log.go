/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// ezlog - A simple log mapping module
//
//	(-1 Disable)
//	 0. Emerg
//	 1. Alert
//	 2. Crit
//	 3. Err
//	 4. Warning
//	 5. Notice
//	 6. Info
//	 7. Debug
//	 8. Trace
package ezlog

import "fmt"

type Level int8

// ezlog log level
const (
	Disabled Level = iota - 1
	EmergLevel
	AlertLevel
	CritLevel
	ErrLevel
	WarningLevel
	NoticeLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

var (
	logLevel = Disabled
	logger   = new(log)
)

type LogFunc func(msg *string)

type log struct {
	// log level func
	emerg   LogFunc
	alert   LogFunc
	crit    LogFunc
	err     LogFunc
	warning LogFunc
	notice  LogFunc
	info    LogFunc
	debug   LogFunc
	trace   LogFunc
	// ---
	msg LogFunc
}

// Get log level
func LogLevel() Level { return logLevel }

// Set log level
func SetLogLevel(level Level) { logLevel = level }

// --- Set func

// Set EMERG logging func
func SetEmerg(f LogFunc) { logger.emerg = f }

// Set ALERT logging func
func SetAlert(f LogFunc) { logger.alert = f }

// Set CRIT logging func
func SetCrit(f LogFunc) { logger.crit = f }

// Set ERR logging func
func SetErr(f LogFunc) { logger.err = f }

// Set WARNING logging func
func SetWarning(f LogFunc) { logger.warning = f }

// Set NOTICE logging func
func SetNotice(f LogFunc) { logger.notice = f }

// Set INFO logging func
func SetInfo(f LogFunc) { logger.info = f }

// Set TRACE logging func
func SetTrace(f LogFunc) { logger.trace = f }

// Set DEBUG logging func
func SetDebug(f LogFunc) { logger.debug = f }

// Set Msg func
func SetMsg(f LogFunc) { logger.msg = f }

// Set all log func to use fmt.Println()
func SetAllPrintln() {
	SetEmerg(func(msg *string) { fmt.Println(*msg) })
	SetAlert(func(msg *string) { fmt.Println(*msg) })
	SetCrit(func(msg *string) { fmt.Println(*msg) })
	SetErr(func(msg *string) { fmt.Println(*msg) })
	SetWarning(func(msg *string) { fmt.Println(*msg) })
	SetNotice(func(msg *string) { fmt.Println(*msg) })
	SetInfo(func(msg *string) { fmt.Println(*msg) })
	SetTrace(func(msg *string) { fmt.Println(*msg) })
	SetDebug(func(msg *string) { fmt.Println(*msg) })
	SetMsg(func(msg *string) { fmt.Println(*msg) })
}

// --- string version func

// Send EMERG
func Emerg(msg string) { EmergP(&msg) }

// Send ALERT
func Alert(msg string) { AlertP(&msg) }

// Send CRIT
func Crit(msg string) { CritP(&msg) }

// Send ERR
func Err(msg string) { ErrP(&msg) }

// Send WARNING
func Warning(msg string) { WarningP(&msg) }

// Send NOTICE
func Notice(msg string) { NoticeP(&msg) }

// Send INFO
func Info(msg string) { InfoP(&msg) }

// Send DEBUG
func Debug(msg string) { DebugP(&msg) }

// Send TRACE
func Trace(msg string) { TraceP(&msg) }

// Send Message without log level control
func Msg(msg string) {
	MsgP(&msg)
}

// --- string pointer version func

// Send EMERG
func EmergP(msg *string) {
	if logLevel >= EmergLevel && logger.emerg != nil {
		logger.emerg(msg)
	}
}

// Send ALERT
func AlertP(msg *string) {
	if logLevel >= AlertLevel && logger.alert != nil {
		logger.alert(msg)
	}
}

// Send CRIT
func CritP(msg *string) {
	if logLevel >= CritLevel && logger.crit != nil {
		logger.crit(msg)
	}
}

// Send ERR
func ErrP(msg *string) {
	if logLevel >= ErrLevel && logger.err != nil {
		logger.err(msg)
	}
}

// Send WARNING
func WarningP(msg *string) {
	if logLevel >= WarningLevel && logger.warning != nil {
		logger.warning(msg)
	}
}

// Send Notice
func NoticeP(msg *string) {
	if logLevel >= NoticeLevel && logger.notice != nil {
		logger.notice(msg)
	}
}

// Send Info
func InfoP(msg *string) {
	if logLevel >= InfoLevel && logger.info != nil {
		logger.info(msg)
	}
}

// Send DEBUG
func DebugP(msg *string) {
	if logLevel >= DebugLevel && logger.debug != nil {
		logger.debug(msg)
	}
}

// Send TRACE
func TraceP(msg *string) {
	if logLevel >= TraceLevel && logger.trace != nil {
		logger.trace(msg)
	}
}

// Send Message without log level control
func MsgP(msg *string) {
	if logger.msg != nil {
		logger.msg(msg)
	}
}
