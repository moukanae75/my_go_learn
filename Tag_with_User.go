package main

import (
	"fmt"
	"regexp"
)

func TagWithUserName(lines []string) []string {
	var new []string
	re := regexp.MustCompile(`User`)
	ex := regexp.MustCompile(`[a-zA-Z]+\d+`)
	for _, c := range lines {
		if re.MatchString(c) && ex.MatchString(c) {
			userName := ex.FindString(c)
			new = append(new, "[User]"+userName+" "+c)
		} else {
			new = append(new, c)
		}
	}
	return new
}

func main() {
	result := TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	})
	for _, line := range result {
		fmt.Println(line)
	}
}
