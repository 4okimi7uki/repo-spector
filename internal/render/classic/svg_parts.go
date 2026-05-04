package classic

import (
	"fmt"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/models"
	"github.com/4okimi7uki/repo-spector/internal/render/color"
	"github.com/dustin/go-humanize"
)

const maxDisplayLangs = 11

func renderStyles(b *strings.Builder) {
	fmt.Fprint(b, `<style>`+"\n")
	fmt.Fprint(b, ` text { color: #9198a1; }`+"\n")
	fmt.Fprint(b, ` tspan { fill: #f3f3f3; font-weight: 600 }`+"\n")
	fmt.Fprint(b, `@keyframes slideIn { from { width: 0; } to { width: 660; } }`)
	fmt.Fprint(b, `#bar_back { animation: slideIn 1.7s ease-in forwards;}`+"\n")
	fmt.Fprint(b, `</style>`+"\n")
}

func renderLangRows(sb *strings.Builder, cfg models.SvgConfig, langItems []models.LangStat) {
	if len(langItems) <= 1 {
		return
	}

	const langRowTemplate = `	<rect x="%f" y="60" width="%f" height="10" fill="%s"> <title>%s %.2f%%</title> </rect>` + "\n"
	const otherTemplate = `		<rect x="%f" y="60" width="%f" height="10" fill="#ededed"> <title>Other %.2f%%</title> </rect>` + "\n"

	currentX := 20.0
	var percentSum float64

	for i, item := range langItems {
		if i >= maxDisplayLangs {
			restPercent := 100.0 - percentSum
			remainingWidth := float64(cfg.BarWidth) - (currentX - 20.0)
			fmt.Fprintf(sb, otherTemplate, currentX, remainingWidth, restPercent)
			return
		}

		ratio := item.Percent / 100.0
		barWidth := ratio * float64(cfg.BarWidth)
		c := color.Safe(item.Color, "#ffffff")
		fmt.Fprintf(sb, langRowTemplate, currentX, barWidth-2.0, c, item.Name, item.Percent)
		currentX += barWidth
		percentSum += item.Percent
	}
}

func renderCircleAndLang(sb *strings.Builder, cfg models.SvgConfig, langItems []models.LangStat) {
	const circleAndLang = `<circle cx="%d" cy="%d" r="%d" fill="%s" />
<text x="%d" y="%d" font-size="13" font-family="system-ui, -apple-system, sans-serif" fill="currentColor"><tspan>%s</tspan> %s%%</text>` + "\n"
	const circleAndOther = `<circle cx="%d" cy="%d" r="%d" fill="#ededed" />
<text x="%d" y="%d" font-size="13" font-family="system-ui, -apple-system, sans-serif" fill="currentColor"><tspan>Other</tspan> %s%%</text>` + "\n"

	const (
		legendStartY     = 105
		legendLineHeight = 24
		legendDotRadius  = 6
		legendColumns    = 2
	)

	columnWidth := cfg.Width/legendColumns - 10

	var percentSum float64
	for i, item := range langItems {
		col := i % legendColumns
		row := i / legendColumns
		legendX := 20 + col*columnWidth
		legendY := legendStartY + row*legendLineHeight

		if i >= maxDisplayLangs {
			restPercent := 100.0 - percentSum
			fmt.Fprintf(sb, circleAndOther,
				legendX+10,
				legendY,
				legendDotRadius,
				legendX+legendDotRadius+20,
				legendY+4,
				humanize.CommafWithDigits(restPercent, 2))
			return
		}

		c := color.Safe(item.Color, "#ffffff")
		fmt.Fprintf(sb, circleAndLang,
			legendX+10,
			legendY,
			legendDotRadius,
			c,
			legendX+legendDotRadius+20,
			legendY+4,
			item.Name,
			humanize.CommafWithDigits(item.Percent, 2),
		)
		percentSum += item.Percent
	}
}
