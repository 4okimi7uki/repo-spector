package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/4okimi7uki/repo-spector/internal/client"
	"github.com/4okimi7uki/repo-spector/internal/gh"
	"github.com/4okimi7uki/repo-spector/internal/models"
	"github.com/4okimi7uki/repo-spector/internal/render"
	"github.com/4okimi7uki/repo-spector/internal/ui"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	showVersion bool
	style       string
	excludeLang string
)

const (
	dirPath = "./output"
	// fileName = "top6_lang.svg"
	fileName = "classic_theme.svg"
)

var rootCmd = &cobra.Command{
	Use:   "repo-spector",
	Short: "Generate language stats SVG for your repositories",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			resolvedVersion := gh.ResolvedVersion()
			fmt.Printf("%s %s\n", resolvedVersion, ui.Turquoise("(repo-spector)"))

			// check latest version
			PrintCheckLatestVersion()
			return nil
		}
		start := time.Now()

		_ = godotenv.Load()
		v := strings.TrimSpace(os.Getenv("GH_TOKEN"))
		resolvedExcludeLang := strings.Split(excludeLang, ",")
		var agg models.LangStatWithTotal

		err := ui.WithSpinner("Fetching...", func(update func(string)) error {
			var repo models.RepositoryCountAndAuthor
			var err error
			c := client.NewClient(v)

			agg, repo, err = c.FetchAllRepo(resolvedExcludeLang)
			if err != nil {
				return err
			}

			update("Build SVG...")
			if err = render.RenderSVG(style, agg, repo, dirPath+"/"+fileName); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

		if err = ui.PrintSummary(agg, resolvedExcludeLang); err != nil {
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
	rootCmd.Flags().StringVar(&style, "style", "", "style")
}

func PrintCheckLatestVersion() {
	resolvedVersion := gh.ResolvedVersion()
	if msg, err := gh.CheckLatestVersion("4okimi7uki", "repo-spector", resolvedVersion); err == nil && msg != "" {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", ui.LimeYellow(msg))
		_, _ = fmt.Fprintf(os.Stdout, "%s\n\n", "https://github.com/4okimi7uki/repo-spector/releases")
	}
}
