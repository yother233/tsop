package logger

import (
	"testing"
	"tsop/internal/pkg/conf"
)

func TestZapOut(t *testing.T) {
	l, f, err := NewZap(&conf.Config{
		Name: "tsop",
		Mode: "debug",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer f()

	l.Info("test input")
}
