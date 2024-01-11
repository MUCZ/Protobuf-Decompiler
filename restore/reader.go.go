package restore

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func goRawDescReader(filepath string) []byte {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("File not found:", filepath)
		return nil
	}
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil
	}
	var str string
	var re = regexp.MustCompile(`var .*_rawDesc = \[\]byte\{([^}]*)\}`)
	match := re.FindStringSubmatch(string(bytes))
	if len(match) <= 1 {
		return nil
	}
	str = match[1] // first match is the whole string
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

func strBytesToBytes(bytes []string) []byte {
	value := make([]byte, 0, len(bytes))
	for _, part := range bytes {
		part = strings.TrimSpace(part)
		val, _ := strconv.ParseUint(part, 0, 8)
		value = append(value, byte(val))
	}
	return value
}
