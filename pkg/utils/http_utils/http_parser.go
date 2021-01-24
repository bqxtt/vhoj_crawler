package http_utils

import (
	"fmt"
	"regexp"
)

func ParseHtmlReg(reg string, html string) (string, error) {
	return ParseHtmlRegIndex(reg, html, 1)
}

func ParseHtmlRegIndex(reg string, html string, index int) (string, error) {
	compile := regexp.MustCompile(reg)
	subMatch := compile.FindStringSubmatch(html)
	if subMatch != nil && len(subMatch) > 0 {
		//fmt.Println(subMatch)
		return HtmlTagFilter(subMatch[index]), nil
	}
	return "", fmt.Errorf("no match")
}

func HtmlTagFilter(html string) string {
	compile := regexp.MustCompile(`<[\s\S]*?>`)
	html = compile.ReplaceAllString(html, "")
	return html
}
