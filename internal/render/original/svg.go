package original

import (
	"fmt"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/models"
)

func defaultSVGConfig() models.SvgConfig {
	return models.SvgConfig{
		Width:  430,
		Height: 304,

		TitleX: 33,
		TitleY: 40,

		TopLangX: 135,
		TopLangY: 80,

		TopPercentX: 330,
		TopPercentY: 70,

		RowStartY: 150,
		RowGap:    30,

		OverviewShiftY: 18,

		MaxBarWidth: 200,
	}
}

func BuildSVG(aggregate models.LangStatWithTotal, repo models.RepositoryCountAndAuthor) (string, error) {

	cfg := defaultSVGConfig()

	var (
		langItems = aggregate.Items
		totalSize = aggregate.Total
	)

	if len(langItems) == 0 {
		return "", fmt.Errorf("no data")
	}

	var b strings.Builder

	topLang := langItems[0]

	fmt.Fprintf(&b, `<svg version="1.1" width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`+"\n",
		cfg.Width, cfg.Height,
	)

	renderStyles(&b, topLang)
	renderBorder(&b, cfg)

	renderHeaderTop(&b, cfg, topLang)

	renderDivider(&b)

	overviewOffset := renderLangRows(&b, cfg, langItems)

	showOverview := len(langItems) <= 2
	if showOverview {
		renderOverview(&b, cfg, overviewOffset, totalSize, repo)
	}

	b.WriteString(`</svg>`)
	return b.String(), nil
}
