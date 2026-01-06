package aggregate

import (
	"math"
	"sort"

	"github.com/4okimi7uki/repo-spector/internal/models"
)

func AggregateLanguages(resp *models.GraphQLResponse, agg map[string]*models.LangAgg, excludeLang map[string]struct{}) {
	nodes := resp.Data.Viewer.Repositories.Nodes
	for _, repo := range nodes {
		for _, e := range repo.Languages.Edges {
			name := e.Node.Name
			if name == "" {
				continue
			}
			if IsExcludeLang(name, excludeLang) {
				continue
			}

			a, ok := agg[name]
			if !ok {
				color := e.Node.Color
				agg[name] = &models.LangAgg{Size: e.Size, Color: &color}
				continue
			}

			a.Size += e.Size
			if a.Color == nil && e.Node.Color != "" {
				a.Color = &e.Node.Color
			}
		}
	}
}

func BuildSortedAgg(agg map[string]*models.LangAgg) models.LangStatWithTotal {
	out := []models.LangStat{}
	var total = 0
	for name, a := range agg {
		out = append(out, models.LangStat{
			Name:  name,
			Size:  a.Size,
			Color: a.Color,
		})
		total += a.Size
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Size > out[j].Size
	})

	if total > 0 {
		for i := range out {
			raw := float64(out[i].Size) / float64(total) * 100
			out[i].Percent = math.Round(raw*100) / 100
		}
	}

	return models.LangStatWithTotal{
		Items: out,
		Total: total,
	}
}
