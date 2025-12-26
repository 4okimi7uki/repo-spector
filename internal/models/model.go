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

// [[{4okimi7uki/adventOfCode2025 https://github.com/4okimi7uki/adventOfCode2025 true {1427 [{1427 {Go #00ADD8}}]}} {4okimi7uki/repo-spector https://github.com/4okimi7uki/repo-spector false {0 []}} {4okimi7uki/obsidian-vault https://github.com/4okimi7uki/obsidian-vault true {0 []}} {4okimi7uki/gh-pr-formatter https://github.com/4okimi7uki/gh-pr-formatter false {30871 [{29772 {Go #00ADD8}} {1099 {Makefile #427819}}]}} {4okimi7uki/4okimi7uki https://github.com/4okimi7uki/4okimi7uki false {0 []}} {4okimi7uki/self-reposcope https://github.com/4okimi7uki/self-reposcope false {12083 [{11540 {Rust #dea584}} {543 {CSS #663399}}]}} {4okimi7uki/website https://github.com/4okimi7uki/website false {16716 [{9651 {TypeScript #3178c6}} {6506 {SCSS #c6538c}} {559 {JavaScript #f1e05a}}]}} {4okimi7uki/business-card https://github.com/4okimi7uki/business-card true {15056 [{13599 {TypeScript #3178c6}} {852 {CSS #663399}} {605 {JavaScript #f1e05a}}]}} {4okimi7uki/mSnip https://github.com/4okimi7uki/mSnip false {4551 [{1406 {TypeScript #3178c6}} {1405 {CSS #663399}} {1386 {Rust #dea584}} {354 {HTML #e34c26}}]}} {4okimi7uki/ui-components https://github.com/4okimi7uki/ui-components false {37218 [{20422 {TypeScript #3178c6}} {13063 {MDX #fcb32c}} {2878 {CSS #663399}} {855 {JavaScript #f1e05a}}]}} {4okimi7uki/logo_game https://github.com/4okimi7uki/logo_game false {43569 [{30110 {TypeScript #3178c6}} {11850 {CSS #663399}} {962 {JavaScript #f1e05a}} {647 {HTML #e34c26}}]}} {4okimi7uki/logo-game-nextjs https://github.com/4okimi7uki/logo-game-nextjs false {11885 [{9061 {TypeScript #3178c6}} {2165 {CSS #663399}} {659 {JavaScript #f1e05a}}]}} {4okimi7uki/agi-lovely https://github.com/4okimi7uki/agi-lovely false {38869 [{32998 {TypeScript #3178c6}} {4191 {Python #3572A5}} {1042 {Shell #89e051}} {555 {CSS #663399}} {83 {JavaScript #f1e05a}}]}} {4okimi7uki/self-reposcope-action https://github.com/4okimi7uki/self-reposcope-action false {1048 [{1048 {JavaScript #f1e05a}}]}} {4okimi7uki/shutdown-watcher https://github.com/4okimi7uki/shutdown-watcher false {792 [{792 {Rust #dea584}}]}} {4okimi7uki/kongari-toast https://github.com/4okimi7uki/kongari-toast false {16085 [{11801 {TypeScript #3178c6}} {3292 {CSS #663399}} {992 {JavaScript #f1e05a}}]}} {4okimi7uki/chatGPT-lovely https://github.com/4okimi7uki/chatGPT-lovely true {0 []}}
// {
// 	4okimi7uki/compare_compress_size
// 	https://github.com/4okimi7uki/compare_compress_size
// 	false
// 	{33936 [
// 			{15265 {JavaScript #f1e05a}}
// 			{6453 {Kotlin #A97BFF}}
// 			{3213 {CSS #663399}}
// 			{2721 {HTML #e34c26}}
// 			{2330 {Dockerfile #384d54}}
// 			{1540 {Python #3572A5}}
// 			{966 {Shell #89e051}}
// 			{753 {Rust #dea584}}
// 			{695 {PHP #4F5D95}}
// 		]
// 	}
// }

// {4okimi7uki/React-Portfolio https://github.com/4okimi7uki/React-Portfolio true {0 []}} {4okimi7uki/myMaterials https://github.com/4okimi7uki/myMaterials true {0 []}}]]
