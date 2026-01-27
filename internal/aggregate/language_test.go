package aggregate

import (
	"testing"

	"github.com/4okimi7uki/repo-spector/internal/models"
)

func strPtr(value string) *string {
	return &value
}

func TestAggregateLanguages(t *testing.T) {
	resp := &models.GraphQLResponse{
		Data: models.Data{
			Viewer: models.Viewer{
				Repositories: models.Repositories{
					Nodes: []models.RepositoryNode{
						{
							Languages: models.Languages{
								Edges: []models.LanguageEdge{
									{Size: 10, Node: models.LanguageNode{Name: "Go", Color: "#00ADD8"}},
									{Size: 5, Node: models.LanguageNode{Name: "JavaScript", Color: ""}},
									{Size: 3, Node: models.LanguageNode{Name: "", Color: "#fff"}},
								},
							},
						},
						{
							Languages: models.Languages{
								Edges: []models.LanguageEdge{
									{Size: 15, Node: models.LanguageNode{Name: "Go", Color: ""}},
									{Size: 5, Node: models.LanguageNode{Name: "JavaScript", Color: "#f1e05a"}},
									{Size: 20, Node: models.LanguageNode{Name: "Python", Color: "#3572A5"}},
								},
							},
						},
					},
				},
			},
		},
	}

	agg := map[string]*models.LangAgg{
		"JavaScript": {Size: 0, Color: nil},
	}
	exclude := map[string]struct{}{"python": {}}
	repoCount := 0

	AggregateLanguages(resp, agg, exclude, &repoCount)

	if repoCount != 2 {
		t.Fatalf("expected repoCount 2, got %d", repoCount)
	}

	goLang := agg["Go"]
	if goLang == nil {
		t.Fatalf("expected Go to be present")
	}
	if goLang.Size != 25 {
		t.Fatalf("expected Go size 25, got %d", goLang.Size)
	}
	if goLang.Color == nil || *goLang.Color != "#00ADD8" {
		t.Fatalf("expected Go color #00ADD8, got %#v", goLang.Color)
	}

	jsLang := agg["JavaScript"]
	if jsLang == nil {
		t.Fatalf("expected JavaScript to be present")
	}
	if jsLang.Size != 10 {
		t.Fatalf("expected JavaScript size 10, got %d", jsLang.Size)
	}
	if jsLang.Color == nil || *jsLang.Color != "#f1e05a" {
		t.Fatalf("expected JavaScript color #f1e05a, got %#v", jsLang.Color)
	}

	if _, ok := agg["Python"]; ok {
		t.Fatalf("did not expect Python to be present")
	}
}

func TestBuildSortedAgg(t *testing.T) {
	agg := map[string]*models.LangAgg{
		"Go":         {Size: 30, Color: strPtr("#00ADD8")},
		"TypeScript": {Size: 20, Color: strPtr("#3178c6")},
	}

	stats := BuildSortedAgg(agg)

	if stats.Total != 50 {
		t.Fatalf("expected total 50, got %d", stats.Total)
	}
	if len(stats.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(stats.Items))
	}
	if stats.Items[0].Name != "Go" || stats.Items[1].Name != "TypeScript" {
		t.Fatalf("expected Go then TypeScript, got %s then %s", stats.Items[0].Name, stats.Items[1].Name)
	}
	if stats.Items[0].Percent != 60.0 {
		t.Fatalf("expected Go percent 60.0, got %v", stats.Items[0].Percent)
	}
	if stats.Items[1].Percent != 40.0 {
		t.Fatalf("expected TypeScript percent 40.0, got %v", stats.Items[1].Percent)
	}
}
