package render

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/4okimi7uki/repo-spector/internal/models"
)

func safeColor(p *string, fallback string) string {
	if p == nil || *p == "" {
		return fallback
	}
	return *p
}

func renderStyles(b *strings.Builder, topLang models.LangStat) {
	fmt.Fprint(b, `  <linearGradient id="gradient" x1="0" x2="1" y1="0" y2="0">
      <stop offset="0%" stop-color="#fff" stop-opacity="0.15" />
      <stop offset="100%" stop-color="#fff" />
</linearGradient>`+"\n\n",
	)

	fmt.Fprint(b, `  <defs>
  <linearGradient id="animeGrad" x1="0" y1="0" x2="1" y2="0" gradientUnits="objectBoundingBox">
    <stop offset="0%" stop-color="#ffffff" stop-opacity="1" />
    <stop offset="50%" stop-color="#ffffff" stop-opacity="0.3" />
    <stop offset="100%" stop-color="#ffffff" stop-opacity="1" />
    <animateTransform
      attributeName="gradientTransform"
      type="translate"
      from="-1 0"
      to="1 0"
      dur="3s"
      repeatCount="indefinite"
    />
  </linearGradient>
</defs>`+"\n\n",
	)

	topColor := safeColor(topLang.Color, "#ffffff")

	//style
	fmt.Fprint(b, `  <style>`+"\n")
	fmt.Fprintf(b, `  .top {animation: fadeIn 1.2s ease-in forwards; filter: drop-shadow(0 0 5px %s);}`+"\n", topColor)
	fmt.Fprint(b, `  .bar {animation: slideIn 1.3s 0.6s cubic-bezier(0.47, 0, 0.745, 0.715) forwards; width: 0}`+"\n")
	fmt.Fprint(b, `  .langRow {animation: fadeIn 1s ease-in forwards;}`+"\n\n")
	fmt.Fprint(b, `  @keyframes fadeIn { from { opacity: 0;} to { opacity: 1;} }`+"\n")
	fmt.Fprint(b, `  @keyframes slideIn { from { width: 0; opacity: 1 } to { width: var(--w); opacity: 1}}`+"\n")
	fmt.Fprint(b, `  </style>`+"\n")

}

func renderBorder(b *strings.Builder, cfg models.SvgConfig) {
	// border
	fmt.Fprintf(b, `  <rect id="border" x="0.5" y="0.5" width="%d" height="%d" fill="#3D444D" rx="5" ry="5" />`+"\n"+` <rect x="0.5" y="0.5" width="429" height="303" rx="4.5" ry="4.5" stroke="#3D444D"/>`+"\n",
		cfg.Width-2, cfg.Height-2)
}

func renderHeaderTop(b *strings.Builder, cfg models.SvgConfig, topLang models.LangStat) {
	// head text
	fmt.Fprintf(b, `  <text id="title" x="%d" y="%d" font-size="14" font-weight="bold" fill="#fff" font-family="system-ui, -apple-system, sans-serif">%s</text>`+"\n", cfg.TitleX, cfg.TitleY, "Most Used Languages")

	// Top lang
	fmt.Fprintf(b, `  <text id="topLang" class="top" x="135" y="80" font-size="38" dominant-baseline="middle" text-anchor="middle" font-weight="bold" fill='#fff' font-family="system-ui, -apple-system, sans-serif">%s</text>`+"\n", topLang.Name)

	// Top Percent
	fmt.Fprintf(b, `  <text id="topPercent" class="top" x="330" y="70" font-size="28" font-weight="bold" fill='#fff' font-family="system-ui, -apple-system, sans-serif" dominant-baseline="middle" text-anchor="middle" >%.2f%%</text>`+"\n", topLang.Percent)
	// Top bytes
	fmt.Fprintf(b, `  <text id="topByte" class="top" x="330" y="95" font-size="16" font-weight="normal" fill='#fff' font-family="system-ui, -apple-system, sans-serif" dominant-baseline="middle" text-anchor="middle" >%d bytes</text>`+"\n", topLang.Size)
}

func renderDivider(b *strings.Builder) {
	b.WriteString("\n" + `  <rect x="14" y="115.8" width="400" height="1" fill="url(#animeGrad)" />` + "\n\n")
}

func renderLangRows(b *strings.Builder, cfg models.SvgConfig, langItems []models.LangStat) int {
	if len(langItems) <= 1 {
		return 0
	}

	const langRowTemplate = `<g transform="translate(0, %[1]d)" class="langRow">
  <text x="33" y="150" font-size="16" font-weight="bold" fill="#fff" font-family="system-ui, -apple-system, sans-serif">%[2]s  <tspan font-size="13" fill="#B6B6B6" font-weight="normal">%.2f%%</tspan></text>
  <rect class="bar" style="--w: %[4]fpx" x="195" y="140" width="%[4]f" height="12" fill="url(#gradient)" />
</g>
`
	rowGap := 30
	if len(langItems) < 6 {
		rowGap = 40
	}
	second := langItems[1]
	denom := float64(second.Size)
	if denom <= 0 {
		denom = 1
	}

	const maxRows = 5
	end := min(1+maxRows, len(langItems)) // slice end
	for i, item := range langItems[1:end] {
		offsetY := rowGap * i

		var rectWidth = math.Round(cfg.MaxBarWidth * float64(item.Size) / denom)

		fmt.Fprintf(b, langRowTemplate+"\n", offsetY, item.Name, item.Percent, rectWidth)
	}

	return cfg.OverviewShiftY
}

func renderOverview(b *strings.Builder, cfg models.SvgConfig, overviewShiftY int, totalSize int, repo models.RepositoryCountAndAuthor) {
	fmt.Fprintf(b, `<g transform="translate(0, %[1]d)">`+"\n", overviewShiftY)
	fmt.Fprint(b, ` <text class="" x="215" y="155" font-size="13" dominant-baseline="middle" text-anchor="middle" font-weight="bold" fill='#fff' font-family="system-ui, -apple-system, sans-serif">Overview</text>`+"\n")

	now := time.Now().Local().Format("2006-01-02 15:04 MST")

	overviewLabels := []string{"Date", "Author", "Repositories scanned", "Total bytes"}
	overviewVals := []string{now, repo.Author, strconv.Itoa(repo.Count), strconv.Itoa(totalSize)}

	const baseOffsetY = 25

	const labelTemplate = `  <text x="%[1]d" y="%[2]d" font-size="13" font-weight="bold" fill='#fff' font-family="system-ui, -apple-system, sans-serif">%s:</text>`
	const valTemplate = `  <text x="265" y="%[1]d" font-size="13" font-weight="normal" fill='#fff' font-family="system-ui, -apple-system, sans-serif">%s</text>`

	for i, label := range overviewLabels {
		offsetY := 180 + baseOffsetY*i
		fmt.Fprintf(b, labelTemplate+"\n", cfg.TitleX, offsetY, label)
		fmt.Fprintf(b, valTemplate+"\n", offsetY, overviewVals[i])
	}

	fmt.Fprint(b, `</g>`+"\n")
}
