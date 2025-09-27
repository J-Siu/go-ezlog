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
//  0. Disable)
//  1. Emerg
//  2. Alert
//  3. Crit
//  4. Err
//  5. Warning
//  6. Notice
//  7. Info
//  8. Debug
//  9. Trace
package ezlog

import (
	"fmt"
	"unicode/utf8"

	"github.com/J-Siu/go-strany"
)

type Level int8

// ezlog log level
const (
	LogLevel Level = iota - 2 // Not exactly a log level. It is for logging regardless of log level
	Disabled
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
	StrAny = strany.New()
)

type OutFunc func(msg *string)

type ezlog struct {
	logLevel    Level
	msgLogLevel Level
	outFunc     OutFunc
	strBuf      []string
}

func (e *ezlog) New() *ezlog {
	e.SetOutPrintLn()
	return e
}

func (e *ezlog) SetOutFunc(f OutFunc) *ezlog {
	e.outFunc = f
	return e
}

// Set out
func (e *ezlog) SetOutPrint() *ezlog {
	e.SetOutFunc(func(str *string) { fmt.Print(*str) })
	return e
}

// Set out
func (e *ezlog) SetOutPrintLn() *ezlog {
	e.SetOutFunc(func(str *string) { fmt.Println(*str) })
	return e
}

// Set log level
func (e *ezlog) SetLogLevel(level Level) *ezlog {
	e.logLevel = level
	return e
}

// Get log level
func (e *ezlog) GetLogLevel() Level { return e.logLevel }

// Clear message
func (e *ezlog) Clear() *ezlog {
	e.strBuf = nil
	return e
}

// --- Output

func (e *ezlog) Out() *ezlog {
	if e.msgLogLevel <= e.logLevel {
		e.outFunc(e.StringP())
	}
	return e
}

func (e *ezlog) String() string { return *e.StringP() }

func (e *ezlog) StringP() *string {
	str := ""
	if e.msgLogLevel <= e.logLevel {
		if e.strBuf != nil {
			// str = strings.Join(l.strBuf, " ")
			for _, s := range e.strBuf {
				_, size := utf8.DecodeLastRuneInString(str)
				if size > 0 && str[len(str)-size] != '\n' {
					str += " "
				}
				str += s
			}
		}
	}
	return &str
}

// --- Build log message

// Add msg to log
func (e *ezlog) Msg(data any) *ezlog {
	if e.msgLogLevel <= e.logLevel {
		e.strBuf = append(e.strBuf, StrAny.Str(data))
	}
	return e
}

// Add separator to message
func (e *ezlog) Sp(data any) *ezlog {
	if e.msgLogLevel <= e.logLevel {
		e.strBuf[len(e.strBuf)-1] = e.strBuf[len(e.strBuf)-1] + StrAny.Str(data)
	}
	return e
}

// Add newline to message
func (e *ezlog) Ln() *ezlog { return e.Msg("\n") }

// Log on new line
func (e *ezlog) MsgLn(data any) *ezlog { return e.Msg(data).Ln() }

// Add : after data
func (e *ezlog) Name(data any) *ezlog { return e.Msg(data).Sp(":") }

// Add : and newline after data
func (e *ezlog) NameLn(data any) *ezlog { return e.Name(data).Ln() }

// -- Set log message level

func (e *ezlog) Log() *ezlog {
	e.Clear().msgLogLevel = LogLevel
	return e
}
func (e *ezlog) Emerg() *ezlog {
	e.Clear().msgLogLevel = EmergLevel
	return e
}
func (e *ezlog) Alert() *ezlog {
	e.Clear().msgLogLevel = AlertLevel
	return e
}
func (e *ezlog) Crit() *ezlog {
	e.Clear().msgLogLevel = CritLevel
	return e
}
func (e *ezlog) Err() *ezlog {
	e.Clear().msgLogLevel = ErrLevel
	return e
}
func (e *ezlog) Warning() *ezlog {
	e.Clear().msgLogLevel = WarningLevel
	return e
}
func (e *ezlog) Notice() *ezlog {
	e.Clear().msgLogLevel = NoticeLevel
	return e
}
func (e *ezlog) Info() *ezlog {
	e.Clear().msgLogLevel = InfoLevel
	return e
}
func (e *ezlog) Debug() *ezlog {
	e.Clear().msgLogLevel = DebugLevel
	return e
}
func (e *ezlog) Trace() *ezlog {
	e.Clear().msgLogLevel = TraceLevel
	return e
}

// ---

var log = New()

func New() *ezlog {
	StrAny.IndentEnable(true)
	return new(ezlog).New()
}

// Set all log func to use fmt.Print()
func SetOutPrint() *ezlog { return log.SetOutPrint() }

// Set all log func to use fmt.Println()
func SetOutPrintLn() *ezlog { return log.SetOutPrintLn() }

// Get log level
func GetLogLevel() Level { return log.GetLogLevel() }

// Set log level
func SetLogLevel(level Level) *ezlog { return log.SetLogLevel(level) }

func Log() *ezlog {
	log.Clear().msgLogLevel = LogLevel
	return log
}
func Emerg() *ezlog {
	log.Clear().msgLogLevel = EmergLevel
	return log
}
func Alert() *ezlog {
	log.Clear().msgLogLevel = AlertLevel
	return log
}
func Crit() *ezlog {
	log.Clear().msgLogLevel = CritLevel
	return log
}
func Err() *ezlog {
	log.Clear().msgLogLevel = ErrLevel
	return log
}
func Warning() *ezlog {
	log.Clear().msgLogLevel = WarningLevel
	return log
}
func Notice() *ezlog {
	log.Clear().msgLogLevel = NoticeLevel
	return log
}
func Info() *ezlog {
	log.Clear().msgLogLevel = InfoLevel
	return log
}
func Debug() *ezlog {
	log.Clear().msgLogLevel = DebugLevel
	return log
}
func Trace() *ezlog {
	log.Clear().msgLogLevel = TraceLevel
	return log
}
