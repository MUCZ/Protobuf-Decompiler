package restore

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func pyDescriptorReader(filepath string) []byte {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil
	}
	var re = regexp.MustCompile(`AddSerializedFile\(b\'([^)]*)\'\)`)
	match := re.FindStringSubmatch(string(bytes))
	if len(match) <= 1 {
		return nil
	}
	return stringToHex(match[1])
}

func stringToHex(s string) []byte {
	var result []byte
	i := 0
	for i < len(s) {
		if value, _, tail, err := strconv.UnquoteChar(s[i:], '"'); err == nil {
			result = append(result, byte(value))
			i += len(string(value))
			s = s[:i] + tail
		} else {
			fmt.Println(err)
			return nil
		}
	}
	return result
}
