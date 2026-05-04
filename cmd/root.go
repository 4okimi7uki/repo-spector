package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/4okimi7uki/repo-spector/internal/client"
	"github.com/4okimi7uki/repo-spector/internal/gh"
	"github.com/4okimi7uki/repo-spector/internal/render"
	"github.com/4okimi7uki/repo-spector/internal/ui"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	showVersion bool
	excludeLang string
)

const (
	dirPath  = "./output"
	fileName = "top6_lang.svg"
)

var rootCmd = &cobra.Command{
	Use:   "repo-spector",
	Short: "Generate language stats SVG for your repositories",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			resolvedVersion := gh.ResolvedVersion()
			fmt.Printf("%s %s\n", resolvedVersion, ui.Mastered("(repo-spector)"))

			// check latest version
			PrintCheckLatestVersion()
			return nil
		}
		start := time.Now()

		_ = godotenv.Load()
		v := strings.TrimSpace(os.Getenv("GH_TOKEN"))

		err := ui.WithSpinner("Generating SVG...", func(update func(string)) error {
			resolvedExcludeLang := strings.Split(excludeLang, ",")

			c := client.NewClient(v)

			agg, repo, err := c.FetchAllRepo(resolvedExcludeLang)
			if err != nil {
				return err
			}

			if err = ui.PrintSummary(agg, resolvedExcludeLang); err != nil {
				return err
			}

			content, err := render.BuildSVG(agg, repo)
			if err != nil {
				return err
			}
			if err = render.WriteSVG(dirPath+"/"+fileName, content); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		elapsed := time.Since(start)
		fmt.Printf("Done in %.1fs 📈✨\n\n", elapsed.Seconds())

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
	rootCmd.PersistentFlags().StringVarP(&excludeLang, "exclude-lang", "x", "", "Exclude languages (e.g. -x 'html,shell')")
}

func PrintCheckLatestVersion() {
	resolvedVersion := gh.ResolvedVersion()
	if msg, err := gh.CheckLatestVersion("4okimi7uki", "repo-spector", resolvedVersion); err == nil && msg != "" {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", ui.LimeYellow(msg))
		_, _ = fmt.Fprintf(os.Stdout, "%s\n\n", "https://github.com/4okimi7uki/repo-spector/releases")
	}
}
