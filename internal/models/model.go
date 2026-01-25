package models

type GraphQLResponse struct {
	Data   Data       `json:"data"`
	Errors []GQLError `json:"errors,omitempty"`
}

type GQLError struct {
	Message string `json:"message"`
}

type Data struct {
	Viewer Viewer `json:"viewer"`
}

type Viewer struct {
	Repositories Repositories `json:"repositories"`
}

type Repositories struct {
	PageInfo PageInfo         `json:"pageInfo"`
	Nodes    []RepositoryNode `json:"nodes"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	EndCursor   string `json:"endCursor"`
}

type RepositoryNode struct {
	NameWithOwner string    `json:"nameWithOwner"`
	URL           string    `json:"url"`
	IsPrivate     bool      `json:"isPrivate"`
	Languages     Languages `json:"languages"`
}

type Languages struct {
	TotalSize int            `json:"totalSize"`
	Edges     []LanguageEdge `json:"edges"`
}

type LanguageEdge struct {
	Size int          `json:"size"`
	Node LanguageNode `json:"node"`
}

type LanguageNode struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type LangAgg struct {
	Size  int
	Color *string
}

type LangStatWithTotal struct {
	Items []LangStat
	Total int
}

type LangStat struct {
	Name    string
	Size    int
	Color   *string
	Percent float64
}

type SvgConfig struct {
	Width  int
	Height int

	TitleX int
	TitleY int

	TopLangX int
	TopLangY int

	TopPercentX int
	TopPercentY int

	RowStartY int
	RowGap    int

	MaxBarWidth float64
}
