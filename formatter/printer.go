package formatter

import (
	"fmt"
	"strings"
)

func readableSize(bytes int) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := unit, 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func PrintFileDeets(fileName, fileType string, lineCount int, byteCount int) {
	sizeStr := readableSize(byteCount)
	fmt.Printf("File: %s\n", fileName)
	fmt.Printf("   ├── Type : %s\n", strings.ToUpper(fileType))
	fmt.Printf("   ├── Lines: %d\n", lineCount)
	fmt.Printf("   └── Size : %s\n", sizeStr)
}
