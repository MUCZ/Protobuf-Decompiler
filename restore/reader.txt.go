package restore

import (
	"fmt"
	"os"
	"strings"
)

func txtRawDescReader(filepath string) []byte {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("File not found:", filepath)
		return nil
	}
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil
	}
	str := string(bytes)
	parts := strings.Split(str, ",")
	strBytes := make([]string, 0, len(parts))
	for i := 0; i < len(parts); i++ {
		s := parts[i]
		s = strings.TrimSpace(s)
		if s != "" {
			strBytes = append(strBytes, s)
		}
	}
	return strBytesToBytes(strBytes)
}
