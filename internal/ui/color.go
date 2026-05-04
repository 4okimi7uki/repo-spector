package ui

import "github.com/fatih/color"

var (
	Red        = color.New(color.FgRed).SprintfFunc()
	Green      = color.RGB(67, 219, 88).SprintfFunc()
	DeepTeal   = color.RGB(43, 119, 119).SprintfFunc()
	Yellow     = color.RGB(255, 219, 76).SprintfFunc()
	Mastered   = color.RGB(208, 175, 76).SprintfFunc()
	Lime       = color.RGB(37, 198, 168).SprintfFunc()
	LimeYellow = color.RGB(202, 234, 119).SprintfFunc()
	MossGray   = color.RGB(96, 94, 82).SprintfFunc()
)
var (
	Bold = color.New(color.Bold).SprintFunc()
)
