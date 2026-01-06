package aggregate

import "strings"

func ToExcludeSet(exclude []string) map[string]struct{} {
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

func IsExcludeLang(lang string, exclude map[string]struct{}) bool {
	_, ok := exclude[strings.ToLower(lang)]
	return ok
}
