package utils

import (
	"strings"
)

var extToLang = map[string]string{
	".go":   "Go",
	".py":   "Python",
	".js":   "Javascript",
	".ts":   "Typescript",
	".java": "Java",
	".cpp":  "C++",
	".c":    "C",
	".rs":   "Rust",
	".rb":   "Ruby",
	".sh":   "Shell",
	".html": "HTML",
	".css":  "CSS",//add your desired languages further
}

func GetLangFromExt(ext string) string {
	ext = strings.ToLower(ext)
	return extToLang[ext]
}
