package render

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/4okimi7uki/repo-spector/internal/models"
	"github.com/4okimi7uki/repo-spector/internal/render/classic"
	"github.com/4okimi7uki/repo-spector/internal/render/original"
)

func writeSVG(path string, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() { _ = f.Close() }()
	_, err = io.WriteString(f, content)
	return err
}

func RenderSVG(style string, agg models.LangStatWithTotal, repo models.RepositoryCountAndAuthor, dirPath string) error {
	var (
		svgData  string
		err      error
		fileName string
	)
	style = strings.TrimSpace(style)

	switch style {
	case "classic":
		fileName = "classic_theme.svg"
		svgData, err = classic.BuildSVG(agg, repo)
		if err != nil {
			return fmt.Errorf("svg: failed to build svg: %w", err)
		}
	default:
		fileName = "top6_lang.svg"
		svgData, err = original.BuildSVG(agg, repo)
		if err != nil {
			return fmt.Errorf("svg: failed to build svg: %w", err)
		}
	}
	filePath := dirPath + "/" + fileName
	if err = writeSVG(filePath, svgData); err != nil {
		return fmt.Errorf("svg: failed to write svg: %w", err)
	}

	return nil
}
