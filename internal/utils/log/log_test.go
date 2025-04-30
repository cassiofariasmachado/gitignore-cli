package log

import (
	"bytes"
	"log"
	"testing"
)

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	debug := true
	Configure(&debug)

	Debug("Debug message: %s", "test")

	if !bytes.Contains(buf.Bytes(), []byte("Debug message: test")) {
		t.Errorf("Expected debug message in log, got %s", buf.String())
	}
}

func TestFatal(t *testing.T) {
	// Skipping Fatal test as it calls os.Exit
	t.Skip("Skipping Fatal test as it calls os.Exit")
}

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	Print("Print message: %s", "test")
	if !bytes.Contains(buf.Bytes(), []byte("Print message: test")) {
		t.Errorf("Expected print message in log, got %s", buf.String())
	}
}
