package classic

import (
	"fmt"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/models"
	"github.com/4okimi7uki/repo-spector/internal/render/color"
	"github.com/dustin/go-humanize"
)

func renderStyles(b *strings.Builder) {
	//style
	fmt.Fprint(b, `  <style>`+"\n")
	fmt.Fprint(b, `  text { color: #9198a1; }`+"\n")
	fmt.Fprint(b, `  tspan { fill: #f3f3f3; font-weight: 600 }`+"\n")
	fmt.Fprint(b, `  </style>`+"\n")
}

func renderLangRows(sb *strings.Builder, cfg models.SvgConfig, langItems []models.LangStat) {
	if len(langItems) <= 1 {
		return
	}

	const langRowTemplate = `	<rect x="%f" y="60" width="%f" height="10" fill="%s"> <title>%s %f%%</title> </rect>` + "\n"
	const otherTemplate = `	<rect x="%f" y="60" width="%d" height="10" fill="#ededed"> <title>Other %f%%</title> </rect>` + "\n"

	var currentX = 20.0
	const displayNum = 11
	var lastLangs = min(len(langItems), displayNum)

	for i, item := range langItems[:lastLangs+1] {
		var percent = item.Percent / 100
		var barWidth = percent * float64(cfg.BarWidth)
		var color = color.Safe(item.Color, "#ffffff")
		if i == displayNum {
			restPercent := 100.0 - percent
			fmt.Fprintf(sb, otherTemplate,
				currentX, cfg.BarWidth, restPercent)
			return
		}

		fmt.Fprintf(sb, langRowTemplate, currentX, barWidth-2.0, color, item.Name, percent)
		currentX += barWidth
	}
}

func renderCircleAndLang(sb *strings.Builder, cfg models.SvgConfig, langItems []models.LangStat) {
	const circleAndLang = `<circle cx="%d" cy="%d" r="%d" fill="%s" />
<text x="%d" y="%d" font-size="13" font-family="system-ui, -apple-system, sans-serif" fill="currentColor"><tspan>%s</tspan> %s%%</text>`
	const circleAndOther = `<circle cx="%d" cy="%d" r="%d" fill="#ededed" />
<text x="%d" y="%d" font-size="13" font-family="system-ui, -apple-system, sans-serif" fill="#ededed"><tspan fill="currentColor">Other</tspan> %s%%</text>`

	legendStartY := 105
	legendLineHeight := 24
	legendDotRadius := 6

	// legend (2 columns)
	var legendColumns = 2
	var columnWidth = cfg.Width/legendColumns - 10
	const displayNum = 11
	var lastLangs = min(len(langItems), displayNum)

	var percentSum float64
	for i, item := range langItems[:lastLangs+1] {
		percentSum += item.Percent
		var percent = humanize.CommafWithDigits(item.Percent, 2)
		var col = i % legendColumns
		var row = i / legendColumns
		var legendX = 20 + col*columnWidth
		var legendY = legendStartY + row*legendLineHeight
		var color = color.Safe(item.Color, "#ffffff")

		if i == displayNum {
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

		fmt.Fprintf(sb, circleAndLang,
			legendX+10,
			legendY,
			legendDotRadius,
			color,
			legendX+legendDotRadius+20,
			legendY+4,
			item.Name,
			percent,
		)
	}

}

// #ededed other
