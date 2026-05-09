# repo-spector

<div align="center">
  | <img src="./output/top6_lang.svg" alt="original_theme" height="280"/> | <img src="./output/classic_theme.svg" alt="classic_theme" /> |
  | :---: | :---: |
  | original theme | classic theme |

![Go Version](https://img.shields.io/badge/Go-1.25-blue?logo=go) ![CI](https://github.com/4okimi7uki/repo-spector/actions/workflows/lint.yml/badge.svg)

</div>

**repo-spector** (a.k.a. [**self-reposcope**](https://github.com/4okimi7uki/self-reposcope)) is a reimagined Go-based implementation of the original Rust-based self-reposcope, featuring a renewed design and a migration from the GitHub REST API to GitHub GraphQL.

## Features

- 📊 Aggregate top languages across repositories
- 🔍 Inspect tech stacks via GitHub GraphQL
- ⚡ Fast, lightweight Go-based CLI

## Requirements

- Go 1.25 or later (recommended)
- GitHub Personal Access Token (set as `GH_TOKEN`)
  - The accessible repositories are determined by the token's scopes

## Usage

```sh
./repo-spector
```

```sh
# Exclude specific languages
./repo-spector -x 'html,css,dockerfile'

# Change SVG style
./repo-spector --style classic
```

## GitHub Actions

To enable automatic updates, follow these two steps:

1. Set up a `GitHub Personal Access Token`

Go to _Settings > Secrets and variables > Actions > [Repository secrets]_,
then add a new secret with:

- **Name**: `GH_TOKEN`
- **Value**: Your GitHub Personal Access Token  
  (with `repo` and `workflow` scopes)

Optionally, to exclude specific languages via Actions, add another secret:

- **Name**: `EXCLUDED_LANGUAGES`
- **Value**: e.g. `html,css,dockerfile`

2. Add the following workflow file to `.github/workflows/repo-spector.yml`.

```yml:repo-spector.yml
name: repo-spector

on:
  schedule:
    - cron: "0 0 * * 1" # Every Monday (UTC)
  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest
    permissions:
      contents: write

    steps:
      - name: Checkout target repo
        uses: actions/checkout@v4

      - name: Download repo-spector binary from GitHub Release
        shell: bash
        run: |
          curl -L https://github.com/4okimi7uki/repo-spector/releases/latest/download/repo-spector -o repo-spector
          chmod +x ./repo-spector

      - name: Run repo-spector CLI
        shell: bash
        env:
          GH_TOKEN: ${{ secrets.GH_TOKEN }}
        run: |
          mkdir -p output
          ./repo-spector -x "${{ secrets.EXCLUDED_LANGUAGES }}"

      - name: Commit and Push updated SVGs
        shell: bash
        env:
          GH_TOKEN: ${{ secrets.GH_TOKEN }}
        run: |
          git config --global user.name 'github-actions[bot]'
          git config --global user.email 'github-actions[bot]@users.noreply.github.com'
          git add output/*.svg
          if git diff --cached --quiet; then
            echo "No changes to commit"
          else
            git commit -m "update: language stats svg"
            git push https://x-access-token:${GH_TOKEN}@github.com/${{ github.repository }} HEAD:main
          fi

```

## Flags

| Flag             | Short | Default      | Description                                    |
| ---------------- | ----- | ------------ | ---------------------------------------------- |
| `--version`      | `-v`  | `false`      | Print version information                      |
| `--exclude-lang` | `-x`  | `""`         | Exclude languages (e.g. `-x 'html,shell'`)     |
| `--style`        | —     | `"original"` | SVG output style (`"original"` or `"classic"`) |

<!--関連する語根 -spect を含む単語
また、「spector」という形ではありませんが、同じ語源を持つ一般的な単語には以下のようなものがあります。
inspect (インスペクト): 調査する、検査する (in- + spect = 中を見る)
expect (エクスペクト): 期待する、予期する (ex- + spect = 外を見る)
respect (リスペクト): 尊敬する、尊重する (re- + spect = 再び見る、顧みる)
suspect (サスペクト): 疑う、怪しいと思う (sus- + spect = 下から見る、見上げる)
perspective (パースペクティブ): 視点、見方、遠近法 (per- + spect = 通して見る)
aspect (アスペクト): 側面、様相 (a- + spect = の方を見る)
これらの単語は、いずれも「見る」という中心的な意味に関連しています。-->

---

<small>2026 Aoki Mizuki – Developed with 🍭 and a sense of fun.</small>
