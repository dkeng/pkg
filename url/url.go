package url

import "regexp"

// GetParam 获取URL参数内容
func GetParam(url, name string) string {
	regex := regexp.MustCompile(`[^\?&]?` + name + `=[^&]+`)
	url = regex.FindString(url)
	if len(url) > 2 {
		return url[2:]
	}
	return ""
}
