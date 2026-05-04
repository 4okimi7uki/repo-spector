package classic

import (
	"fmt"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/models"
)

func defaultSVGConfig() models.SvgConfig {
	return models.SvgConfig{
		BarWidth:  370,
		BarHeight: 10,
		Width:     410,
		Height:    255,
	}
}

func BuildSVG(aggregate models.LangStatWithTotal, repo models.RepositoryCountAndAuthor) (string, error) {
	var (
		langItems = aggregate.Items
	)

	if len(langItems) == 0 {
		return "", fmt.Errorf("no data")
	}

	cfg := defaultSVGConfig()

	var sb strings.Builder
	fmt.Fprintf(&sb, `<svg version="1.1" width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`+"\n",
		cfg.Width, cfg.Height,
	)
	// style
	renderStyles(&sb)

	//card
	fmt.Fprintf(&sb, `<rect id="border" x="0" y="0" width="%d" height="%d" fill='#3D444D' rx="15" ry="15" />`+"\n", cfg.Width, cfg.Height)
	fmt.Fprintf(&sb, `<rect id="bg" x="1" y="1" width="%d" height="%d" rx="14.5" ry="14.5" />`+"\n\n", cfg.Width-2, cfg.Height-2)

	// title
	sb.WriteString(`<text id="title" x="20" y="40" font-size="18" font-weight="bold" fill='#fff' font-family="system-ui, -apple-system, sans-serif">Top Languages</text>` + "\n")

	fmt.Fprintf(&sb, `<defs>
	<clipPath id="roundedClip">
    	<rect id="bar_back" x="20" y="60" width="%d" height="%d" rx="5" ry="5">
     	</rect>
    </clipPath>
</defs>`+"\n\n", cfg.BarWidth, cfg.BarHeight)

	fmt.Fprintf(&sb, `<rect width="%d" height="%d" x="20" y="60" rx="5" ry="5" fill="#3D444D" />`+"\n", cfg.BarWidth, cfg.BarHeight)

	fmt.Fprintf(&sb, `<g clip-path="url(#roundedClip)">`+"\n")
	renderLangRows(&sb, cfg, langItems)
	sb.WriteString(`</g>` + "\n\n")

	renderCircleAndLang(&sb, cfg, langItems)
	sb.WriteString("\n" + `</svg>`)

	return sb.String(), nil
}
