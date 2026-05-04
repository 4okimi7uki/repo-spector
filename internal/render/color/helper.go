package color

func Safe(p *string, fallback string) string {
	if p == nil || *p == "" {
		return fallback
	}
	return *p
}
