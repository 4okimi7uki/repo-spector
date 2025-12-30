package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/client"
	"github.com/4okimi7uki/repo-spector/internal/render"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var version = "v0.0.0-dev"
var showVersion bool

const (
	dirPath  = "./output"
	fileName = "top6_lang.svg"
)

var rootCmd = &cobra.Command{
	Use:   "repo-spector",
	Short: "Generate language stats SVG for your repositories",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Printf("repo-spector %s\n", version)
			return nil
		}

		_ = godotenv.Load()
		v := strings.TrimSpace(os.Getenv("GH_TOKEN"))

		excludeLang := []string{"HTML", "CSS", "Makefile", "MDX", "TypeScript"}

		c := client.NewClient(v)

		agg, err := c.FetchAllRepo(excludeLang)
		if err != nil {
			return err
		}
		fmt.Println(agg)
		bar := strings.Repeat("─", 20+len(excludeLang)*8)
		_, _ = fmt.Fprintln(os.Stdout, bar)
		_, _ = fmt.Fprintf(os.Stdout, " Exclude Languages: %s\n", excludeLang)
		_, _ = fmt.Fprintln(os.Stdout, bar)

		content := render.BuildSVG(agg)
		if err = render.WriteSVG(dirPath+"."+fileName, content); err != nil {
			return err
		}

		return nil
	},
}

func Excute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Print version information")
}
