package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/caioeverest/vulcan/infra/prompter"
	"github.com/fatih/color"
)

type Console struct {
	Prompt   prompter.Prompter
	logLevel Level
}

var _global Console

// Console constants character to messages.
const (
	fearfulMark = "\U0001f628"
	debugMark   = "\U0001f50d"
	checkMark   = "✔"
	infoMark    = "i"
	warnMark    = "!"
	errorMark   = "✖"
)

func Get() Console { return GetWithParams(new(prompter.CliPrompter), LevelInfo) }

func GetWithParams(p prompter.Prompter, level Level) Console { return Console{p, level} }

func (c *Console) SetDebugLevel() {
	c.logLevel.set(LevelDebug)
}

// Debug logs the given message as a debug message.
func (c *Console) Debug(msg string) {
	if !c.logLevel.permits(LevelDebug) {
		return
	}

	coloredPrintMsg(debugMark, msg, color.FgYellow, color.FgYellow)
}

// Debugf logs the given message as a debug message.
func (c *Console) Debugf(msg string, params ...interface{}) {
	msgf := fmt.Sprintf(msg, params...)
	c.Debug(msgf)
}

// Info logs the given message as a info message.
func (c *Console) Info(msg string) {
	if !c.logLevel.permits(LevelInfo) {
		return
	}

	coloredPrintMsg(infoMark, msg, color.FgWhite, color.FgBlue)
}

// Infof logs the given message as a info message.
func (c *Console) Infof(msg string, params ...interface{}) {
	msgf := fmt.Sprintf(msg, params...)
	c.Info(msgf)
}

// Error logs the given message as a error message.
func (c *Console) Error(msg string) {
	if !c.logLevel.permits(LevelError) {
		return
	}

	coloredPrintMsg(errorMark, msg, color.FgWhite, color.FgRed)
}

// Errorf logs the given message as a error message.
func (c *Console) Errorf(msg string, params ...interface{}) {
	msgf := fmt.Sprintf(msg, params...)
	c.Error(msgf)
}

// Warn logs the given message as a warn message.
func (c *Console) Warn(msg string) {
	if !c.logLevel.permits(LevelWarn) {
		return
	}

	coloredPrintMsg(warnMark, msg, color.FgMagenta, color.FgMagenta)
}

// Warnf logs the given message as a warn message.
func (c *Console) Warnf(msg string, params ...interface{}) {
	msgf := fmt.Sprintf(msg, params...)
	c.Warn(msgf)
}

// Success logs the given message as a success message.
func (c *Console) Success(msg string) {
	if !c.logLevel.permits(LevelSuccess) {
		return
	}

	coloredPrintMsg(checkMark, msg, color.FgWhite, color.FgGreen)
}

// Successf logs the given message as a success message.
func (c *Console) Successf(msg string, params ...interface{}) {
	msgf := fmt.Sprintf(msg, params...)
	c.Success(msgf)
}

// Fatal logs the given message as a fatal message.
func (c *Console) Fatal(msg string) {
	if !c.logLevel.permits(LevelFatal) {
		return
	}

	coloredPrintMsg(fearfulMark, msg, color.FgWhite, color.FgRed)
	os.Exit(1)
}

// Fatalf logs the given message as a fatal message.
func (c *Console) Fatalf(msg string, params ...interface{}) {
	msgf := fmt.Sprintf(msg, params...)
	c.Fatal(msgf)
}

// WaitForEnter expect any character in to continue
func (c *Console) WaitForEnter() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Err()
}

func coloredPrintMsg(icon string, msg string, iC color.Attribute, mC color.Attribute) {
	fmt.Println(
		color.New(color.Bold, mC).SprintFunc()(icon),
		color.New(color.Bold, iC).SprintFunc()(msg))
}
