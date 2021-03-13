package console

import (
	"testing"

	"github.com/caioeverest/vulcan/infra/prompter"
)

func beforeEach() Console {
	prompterMock := new(prompter.MockPrompter)
	return GetWithParams(prompterMock, LevelDebug)
}

func TestConsole_Debug(t *testing.T) {
	target := beforeEach()
	target.Debug("Test")
	target.Debugf("Test %d", 1234)
}

func TestConsole_Info(t *testing.T) {
	target := beforeEach()
	target.Info("Test")
	target.Infof("Test %d", 1234)
}

func TestConsole_Error(t *testing.T) {
	target := beforeEach()
	target.Error("Test")
	target.Errorf("Test %d", 1234)
}

func TestConsole_Warn(t *testing.T) {
	target := beforeEach()
	target.Warn("Test")
	target.Warnf("Test %d", 1234)
}

func TestConsole_Success(t *testing.T) {
	target := beforeEach()
	target.Success("Test")
	target.Successf("Test %d", 1234)
}
