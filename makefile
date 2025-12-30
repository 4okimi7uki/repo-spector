APP_NAME := repo-spector
BUILD_DIR := dist
ROOT_PKG := github.com/4okimi7uki/repo-spector/cmd
VERSION ?= v0.0.0-dev

.PHONY: default build clean

default: build

build:
	@echo "🚀 Building for your current OS (version: $(VERSION))"
	go build -ldflags "-X $(ROOT_PKG).version=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME) .

clean:
	@rm -rf $(BUILD_DIR)
	@echo "🧹 Cleaned build directory"
