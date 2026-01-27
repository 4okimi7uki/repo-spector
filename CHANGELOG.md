## v0.2.0

### Added

- Display an Overview section in the SVG when the number of detected languages is 2 or fewer, including the date, author, scanned repository count, and total byte size.
- Added unit tests for language aggregation logic and exclude-language normalization and matching.
- Added an Issue template.

### Changed

- Refactored SVG layout values into `SvgConfig` to centralize and simplify rendering configuration.
- Expanded the README Usage section (GitHub Actions / CLI usage and `GH_TOKEN` configuration details).
