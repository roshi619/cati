package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCATI(t *testing.T) {
	if _, err := exec.LookPath("cati"); err != nil {
		t.Skip("cati binary missing from PATH")
	}

	t.Run("show version", func(t *testing.T) {
		data, err := exec.Command("cati", "--version").Output()
		if err != nil {
			t.Error(err)
		}
		out := string(data)

		if !strings.Contains(out, "cati version") {
			t.Error("Missing 'cati version'")
		}
		if !strings.Contains(out, "Latest:") {
			t.Error("Missing name of latest version")
		}
		if !strings.Contains(out, "Download: https://github.com/roshi619/cati/releases") {
			t.Error("Missing latest download link")
		}
	})

	t.Run("dry run", func(t *testing.T) {
		cmd := exec.Command("cati", "--verbose", "-b=0")
		cmd.Env = []string{}

		data, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		out := string(data)

		if !strings.Contains(out, "0 catifications queued") {
			t.Error("Unexpected queued catifications")
			t.Error(out)
		}
	})
}
