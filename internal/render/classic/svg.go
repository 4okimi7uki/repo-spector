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

		Width:  410,
		Height: 255,

		// TitleX: 33,
		// TitleY: 40,

		// TopLangX: 135,
		// TopLangY: 80,

		// TopPercentX: 330,
		// TopPercentY: 70,

		// RowStartY: 150,
		// RowGap:    30,

		// OverviewShiftY: 18,

		// MaxBarWidth: 200,
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

	fmt.Fprintf(&sb, `<rect id="border" x="0" y="0" width="%d" height="%d" fill='#3D444D' rx="15" ry="15" />`+"\n", cfg.Width, cfg.Height)
	fmt.Fprintf(&sb, `<rect id="bg" x="1" y="1" width="%d" height="%d" rx="14.5" ry="14.5" />`+"\n\n", cfg.Width-2, cfg.Height-2)

	// title
	sb.WriteString(`<text id="title" x="20" y="40" font-size="18" font-weight="bold" fill='#fff' font-family="system-ui, -apple-system, sans-serif">Top Languages</text>` + "\n")

	// avatar
	if repo.AvatarUrl != "" {
		fmt.Fprint(&sb, `  <defs><clipPath id="avatarClip"><circle cx="19" cy="19" r="14" /></clipPath></defs>`+"\n")
		fmt.Fprintf(&sb, `  <image href="%s" x="5" y="5" width="28" height="28" clip-path="url(#avatarClip)" />`+"\n", repo.AvatarUrl)
	}
	sb.WriteString("\n")

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
