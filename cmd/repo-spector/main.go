package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/client"
	"github.com/4okimi7uki/repo-spector/internal/render"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	v := strings.TrimSpace(os.Getenv("GH_TOKEN"))

	excludeLang := []string{"HTML", "CSS", "Makefile", "MDX"}

	c := client.NewClient(v)

	agg, err := c.FetchAllRepo(excludeLang)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(agg)

	content := render.BuildSVG(agg)
	err = render.WriteSVG("./output/top6_lang.svg", content)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

	}

}
