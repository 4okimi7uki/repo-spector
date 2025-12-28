package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/models"
)

type Client struct {
	httpClient *http.Client
	endpoint   string
	token      string
}

const githubGraphQLEndpoint = "https://api.github.com/graphql"

func NewClient(token string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		endpoint:   githubGraphQLEndpoint,
		token:      token,
	}
}

func (c *Client) Do(query string, vars map[string]any, v any) error {
	body := map[string]any{
		"query":     query,
		"variables": vars,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GitHub GraphQL error: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) fetchRepos(first int, afterCursor *string) (*models.GraphQLResponse, error) {
	const query = `
	query($first: Int!, $after: String) {
  viewer {
    repositories(
      first: $first
      after: $after
      ownerAffiliations: OWNER
      orderBy: {field: UPDATED_AT, direction: DESC}
    ) {
      pageInfo { hasNextPage endCursor }
      nodes {
        nameWithOwner
        url
        isPrivate
        languages(first: 20, orderBy: {field: SIZE, direction: DESC}) {
          totalSize
          edges {
            size
            node { name color }
          }
        }
      }
    }
  }
}`

	vars := map[string]any{
		"first": first,
		"after": afterCursor,
	}

	var result models.GraphQLResponse

	if err := c.Do(query, vars, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FetchAllRepo(excludeLang []string) (models.LangStatWithTotal, error) {
	agg := map[string]*models.LangAgg{}
	var after *string = nil
	excludeSet := toExcludeSet(excludeLang)

	for {
		resp, err := c.fetchRepos(50, after)
		if err != nil {
			return models.LangStatWithTotal{}, err
		}
		c.AggregateLanguages(resp, agg, excludeSet)

		pi := resp.Data.Viewer.Repositories.PageInfo
		if !pi.HasNextPage || pi.EndCursor == "" {
			break
		}

		after = &pi.EndCursor
	}

	return BuildSortedAgg(agg), nil
}

func (c *Client) AggregateLanguages(resp *models.GraphQLResponse, agg map[string]*models.LangAgg, excludeLang map[string]struct{}) {
	nodes := resp.Data.Viewer.Repositories.Nodes
	for _, repo := range nodes {
		for _, e := range repo.Languages.Edges {
			name := e.Node.Name
			if name == "" {
				continue
			}
			if isExcludeLang(name, excludeLang) {
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

func toExcludeSet(exclude []string) map[string]struct{} {
	set := make(map[string]struct{}, len(exclude))
	for _, ex := range exclude {
		x := strings.ToLower(strings.TrimSpace(ex))
		if x == "" {
			continue
		}
		set[x] = struct{}{}
	}
	return set
}

func isExcludeLang(lang string, exclude map[string]struct{}) bool {
	_, ok := exclude[strings.ToLower(lang)]
	return ok
}
