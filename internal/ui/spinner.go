package ui

import (
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

func padSuffix(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}

func WithSpinner(initialMsg string, fn func(update func(string)) error) error {
	const spinnerSuffixWidth = 30
	const interval = 140 * time.Millisecond
	s := spinner.New(spinner.CharSets[29], 100*time.Millisecond)
	s.Writer = os.Stderr
	s.Suffix = padSuffix(" "+initialMsg, spinnerSuffixWidth)

	s.Start()
	defer func() {
		time.Sleep(600 * time.Millisecond)
		s.Stop()
	}()

	update := func(msg string) {
		s.Lock()
		s.Suffix = padSuffix(" "+msg, spinnerSuffixWidth)
		s.Unlock()
	}

	return fn(update)
}
