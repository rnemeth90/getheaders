package getheaders

import "regexp"

func IsValidURL(url string) bool {
	var urlRegex = regexp.MustCompile(`^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`)
	return urlRegex.MatchString(url)
}
