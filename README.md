# ezlog

A simple log mapping module with Linux log level:

- (-1 Disable)
0. Emerg
1. Alert
2. Crit
3. Err
4. Warning
5. Notice
6. Info
7. Debug
8. Trace

## FUNCTIONS

### Send Log
```go
func Alert(msg string)
```
Send ALERT

```go
func AlertP(msg *string)
```
Send ALERT

```go
func Crit(msg string)
```
Send CRIT

```go
func CritP(msg *string)
```
Send CRIT

```go
func Debug(msg string)
```
Send DEBUG

```go
func DebugP(msg *string)
```
Send DEBUG

```go
func Emerg(msg string)
```
Send EMERG

```go
func EmergP(msg *string)
```
Send EMERG

```go
func Err(msg string)
```
Send ERR

```go
func ErrP(msg *string)
```
Send ERR

```go
func Info(msg string)
```
Send INFO

```go
func InfoP(msg *string)
```
Send Info

```go
func Msg(msg string)
```
Send Message without log level control

```go
func MsgP(msg *string)
```
Send Message without log level control

```go
func Notice(msg string)
```
Send NOTICE

```go
func NoticeP(msg *string)
```
Send Notice

### Set Log Functions

```go
func SetAlert(f LogFunc)
```
SetALERT logging func

```go
func SetAllPrintln()
```
Set all log func to use fmt.Println()

```go
func SetCrit(f LogFunc)
```
SetCRIT logging func

```go
func SetDebug(f LogFunc)
```
SetDEBUG logging func

```go
func SetEmerg(f LogFunc)
```
SetEMERG logging func

```go
func SetErr(f LogFunc)
```
SetERR logging func

```go
func SetInfo(f LogFunc)
```
SetINFO logging func

```go
func SetLogLevel(level Level)
```
Setlog level

```go
func SetMsg(f LogFunc)
```
SetMsg func

```go
func SetNotice(f LogFunc)
```
SetNOTICE logging func

```go
func SetTrace(f LogFunc)
```
SetTRACE logging func

```go
func SetWarning(f LogFunc)
```
SetWARNING logging func

```go
func Trace(msg string)
```
Send TRACE

```go
func TraceP(msg *string)
```
Send TRACE

```go
func Warning(msg string)
```
Send WARNING

```go
func WarningP(msg *string)
```
Send WARNING

## TYPES

```go
type Level int8

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
```
ezlog log level

```go
func LogLevel() Level
```
Get log level

```go
type LogFunc func(msg *string)
```

### Change Log

- v1.0.0
  - Initial commit

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
