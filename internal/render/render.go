package render

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/4okimi7uki/repo-spector/internal/models"
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

func RenderSVG(agg models.LangStatWithTotal, repo models.RepositoryCountAndAuthor) error {
	var svgData, err = original.BuildSVG(agg, repo)

	if err != nil {
		return fmt.Errorf("svg: failed to build svg: %w", err)
	}
	if err = writeSVG("", svgData); err != nil {
		return fmt.Errorf("svg: failed to write svg: %w", err)
	}

	return nil
}
