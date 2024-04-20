package helps

import (
	"net/url"
	"strings"
)

func IsURL(str string) bool {
	url, err := url.Parse(str)
	return err == nil && url.Scheme != "" && url.Host != ""
}

func FindInSlice(slice []string, target string) string {
	var result []string
	for _, str := range slice {
		if strings.Contains(str, target) {
			result = append(result, str) 
			return result[0]
		} 
	}
	return ""
}

func ExtractString(str string, target string) string {
		strSlice := strings.Split(str, " ") 
		res := FindInSlice(strSlice, target)
		return res
}