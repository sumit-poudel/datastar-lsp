package util

import (
	"embed"
	"io/fs"
	"path"
	"strings"
	"unicode"
)

func ExtractMDBases(fileList []fs.DirEntry) []string {
	var names []string
	for _, fileMD := range fileList {
		if fileMD.IsDir() {
			continue
		}
		if name, ok := strings.CutSuffix(fileMD.Name(), ".md"); ok {
			names = append(names, name)
		}
	}
	return names
}

func GetAttributeDoc(attr string, embedFs embed.FS) string {
	filePath := path.Join("datastar", "attributes", attr+".md")
	data, err := embedFs.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(data)
}

func GetActionDoc(attr, action string, embedFS embed.FS) string {
	if !strings.HasSuffix(action, ".md") {
		action += ".md"
	}
	filePath := path.Join("datastar", attr, action)
	data, err := embedFS.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(data)
}

func GetLS(path string, embedFS embed.FS) []fs.DirEntry {
	attributes, err := embedFS.ReadDir(path)
	if err != nil {
		return nil
	}
	return attributes
}

func WordAt(line string, char int) string {
	isWordChar := func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == ':'
	}
	runes := []rune(line)
	if char < 0 {
		char = 0
	}
	if char > len(runes) {
		char = len(runes)
	}
	start := char
	for start > 0 && isWordChar(runes[start-1]) {
		start--
	}
	end := char
	for end < len(runes) && isWordChar(runes[end]) {
		end++
	}
	return string(runes[start:end])
}
