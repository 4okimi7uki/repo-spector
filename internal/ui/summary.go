package ui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/4okimi7uki/repo-spector/internal/models"
	"github.com/dustin/go-humanize"
)

func PrintSummary(aggregate models.LangStatWithTotal, excludeLang []string) error {
	var (
		langItems = aggregate.Items
		Total     = aggregate.Total
	)

	now := time.Now().Local().Format("2006-01-02 15:04 MST")

	bar := strings.Repeat("─", 50)
	_, _ = fmt.Println()
	_, _ = fmt.Fprintln(os.Stdout, bar)

	_, _ = fmt.Fprintf(os.Stdout, " %-18s : %s\n", "Date", now)
	_, _ = fmt.Fprintf(os.Stdout, " %-18s : %s byte\n", "Total", humanize.Comma(int64(Total)))
	_, _ = fmt.Fprintf(os.Stdout, " %-18s : %s\n", "Exclude Languages", excludeLang)

	_, _ = fmt.Fprintln(os.Stdout, bar)
	_, _ = fmt.Fprintln(os.Stdout, "\n - TOP LANGUEGES BY USAGE -")
	_, _ = fmt.Fprintln(os.Stdout, bar)
	for idx, item := range langItems {
		_, _ = fmt.Fprintf(
			os.Stdout,
			" #%2d  %-15s %10s byte  %6.2f%%\n",
			idx+1,
			item.Name,
			humanize.Comma(int64(item.Size)),
			item.Percent,
		)
	}
	_, _ = fmt.Fprintln(os.Stdout, bar)
	return nil
}
