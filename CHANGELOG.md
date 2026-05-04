## v0.3.0

### Added

- Added `--theme` option to switch between `original` (default) and `classic` themes for SVG output.
- Added animation to SVG output.
- Added version checking: notifies when a newer release is available via GitHub API.

### Changed

- Reorganized rendering logic into `internal/render/original/` and `internal/render/classic/` packages.
- Moved spinner from `cmd/` to `internal/ui` package.
- Added color helper functions to `internal/ui` package.
- Migrated from pre-commit to lefthook for lint hooks.

### Dependencies

- Bumped `golangci/golangci-lint-action` from 7 to 9.
- Bumped `softprops/action-gh-release` from 2 to 3.
- Bumped `actions/checkout` from 4 to 6.
- Bumped `actions/setup-go` from 5 to 6.

## v0.2.0

### Added

- Display an Overview section in the SVG when the number of detected languages is 2 or fewer, including the date, author, scanned repository count, and total byte size.
- Added unit tests for language aggregation logic and exclude-language normalization and matching.
- Added an Issue template.

### Changed

- Refactored SVG layout values into `SvgConfig` to centralize and simplify rendering configuration.
- Expanded the README Usage section (GitHub Actions / CLI usage and `GH_TOKEN` configuration details).
