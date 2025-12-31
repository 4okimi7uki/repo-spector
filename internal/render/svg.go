package render

import (
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/models"
)

func WriteSVG(path string, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() { _ = f.Close() }()
	_, err = io.WriteString(f, content)
	return err
}

func BuildSVG(aggregate models.LangStatWithTotal) (string, error) {
	const (
		width  = 430
		height = 304
	)
	var langItems = aggregate.Items

	if len(langItems) == 0 {
		return "", fmt.Errorf("no data")
	}

	var b strings.Builder

	topLang := langItems[0]

	fmt.Fprintf(&b, `<svg version="1.1" width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`+"\n",
		width, height,
	)
	fmt.Fprint(&b, `  <linearGradient id="gradient" x1="0" x2="1" y1="0" y2="0">
      <stop offset="0%" stop-color="#fff" stop-opacity="0.15" />
      <stop offset="100%" stop-color="#fff" />
</linearGradient>`+"\n\n",
	)

	fmt.Fprint(&b, `  <defs>
  <linearGradient id="animeGrad" x1="0" y1="0" x2="1" y2="0"
                  gradientUnits="objectBoundingBox">
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

	//style
	fmt.Fprint(&b, `  <style>`+"\n")
	fmt.Fprintf(&b, `  .top {animation: fadeIn 1.2s ease-in forwards; filter: drop-shadow(0 0 5px %s);}`+"\n", *topLang.Color)
	fmt.Fprint(&b, `  .bar {animation: slideIn 1.3s 0.6s cubic-bezier(0.47, 0, 0.745, 0.715) forwards; width: 0}`+"\n")
	fmt.Fprint(&b, `  .langRow {animation: fadeIn 1s ease-in forwards;}`+"\n\n")
	fmt.Fprint(&b, `  @keyframes fadeIn { from { opacity: 0;} to { opacity: 1;} }`+"\n")
	fmt.Fprint(&b, `  @keyframes slideIn { from { width: 0; opacity: 1 } to { width: var(--w); opacity: 1}}`+"\n")
	fmt.Fprint(&b, `  </style>`+"\n")

	// border
	fmt.Fprintf(&b, `  <rect id="border" x="0.5" y="0.5" width="%d" height="%d" fill="#3D444D" rx="5" ry="5" />`+"\n"+` <rect x="0.5" y="0.5" width="429" height="303" rx="4.5" ry="4.5" stroke="#3D444D"/>`+"\n",
		width-2, height-2)

	// head text
	b.WriteString(`  <text id="title" x="33" y="40" font-size="14" font-weight="bold" fill="#fff" font-family="system-ui, -apple-system, sans-serif">Most Used Languages</text>` + "\n")

	// Top lang
	fmt.Fprintf(&b, `  <text id="topLang" class="top" x="135" y="80" font-size="38" dominant-baseline="middle" text-anchor="middle" font-weight="bold" fill='#fff' font-family="system-ui, -apple-system, sans-serif">%s</text>`+"\n", topLang.Name)

	// Top Percent
	fmt.Fprintf(&b, `  <text id="topPercent" class="top" x="330" y="70" font-size="28" font-weight="bold" fill='#fff' font-family="system-ui, -apple-system, sans-serif" dominant-baseline="middle" text-anchor="middle" >%.2f%%</text>`+"\n", topLang.Percent)
	// Top bytes
	fmt.Fprintf(&b, `  <text id="topByte" class="top" x="330" y="95" font-size="16" font-weight="normal" fill='#fff' font-family="system-ui, -apple-system, sans-serif" dominant-baseline="middle" text-anchor="middle" >%d bytes</text>`+"\n", topLang.Size)

	// line
	b.WriteString("\n" + `  <rect x="14" y="115.8" width="400" height="1" fill="url(#animeGrad)" />` + "\n\n")

	const langRowTemplate = `<g transform="translate(0, %[1]d)" class="langRow">
  <text id="lang-name" x="33" y="150" font-size="16" font-weight="bold" fill="#fff" font-family="system-ui, -apple-system, sans-serif">%[2]s  <tspan font-size="13" fill="#B6B6B6" font-weight="normal">%.2f%%</tspan></text>
  <rect class="bar" style="--w: %[4]fpx" x="195" y="140" width="%[4]f" height="12" fill="url(#gradient)" />
</g>
`
	secondLang := langItems[1]
	for i, item := range langItems[1:] {
		if i == 5 {
			break
		}

		rowOffsetY := 30 * i

		const maxRectWidth = 200
		var rectWidth = math.Round(maxRectWidth * float64(item.Size) / float64(secondLang.Size))

		fmt.Fprintf(&b, langRowTemplate+"\n", rowOffsetY, item.Name, item.Percent, rectWidth)
	}

	b.WriteString(`</svg>`)
	return b.String(), nil
}
