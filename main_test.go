package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()
	w.Close()
	os.Stdout = stdout
	return <-outC
}

func TestMain(t *testing.T) {
	output := captureOutput(main)
	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}
