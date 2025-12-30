package cmd

import (
	"fmt"
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

func WithSpinner(initial string, fn func(update func(string)) error) error {
	const spinnerSuffixWidth = 30
	const interval = 100 * time.Millisecond
	s := spinner.New(spinner.CharSets[29], 100*time.Millisecond)
	s.Writer = os.Stderr
	s.Suffix = padSuffix(initial, spinnerSuffixWidth)

	s.Start()
	defer func() {
		fmt.Fprint(os.Stderr, "\r\033[K") // 行をリセット
		s.Stop()
	}()

	update := func(msg string) {
		s.Suffix = padSuffix(" "+msg, spinnerSuffixWidth)
		time.Sleep(interval + 10*time.Millisecond)
	}

	return fn(update)
}
