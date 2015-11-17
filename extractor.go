package extractor

import (
	"regexp"
	"strings"
)

const blkSize int = 3

//remove html tag
func strip(src string) string {
	src = strings.ToLower(src)
	re, _ := regexp.Compile(`<!doctype.*?>`)
	src = re.ReplaceAllString(src, "")
	re, _ = regexp.Compile(`<!--.*?-->`)
	src = re.ReplaceAllString(src, "")
	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")
	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")
	re, _ = regexp.Compile(`<.*?>`)
	src = re.ReplaceAllString(src, "")
	re, _ = regexp.Compile(`&.{1,5};|&#.{1,5};`)
	src = re.ReplaceAllString(src, "")
	src = strings.Replace(src, "\r\n", "\n", -1)
	src = strings.Replace(src, "\r", "\n", -1)
	return src
}

//parse block
func parse(src string) ([]string, []int) {
	var (
		lines   []string
		blksLen []int
	)
	array := strings.Split(src, "\n")
	for _, a := range array {
		a = strings.Replace(a, " ", "", -1)
		lines = append(lines, a)
	}
	blen := 0
	for i := 0; i < blkSize; i++ {
		blen += len(lines[i])
	}
	blksLen = append(blksLen, blen)
	for i := 1; i < len(lines)-blkSize; i++ {
		blen = blksLen[i-1] + len(lines[i-1+blkSize]) - len(lines[i-1])
		blksLen = append(blksLen, blen)
	}
	return lines, blksLen
}

// Extractor main content from html
func Extractor(html string) string {
	lines, blksLen := parse(strip(html))
	i, max := 0, 0
	blkNum := len(blksLen)
	plainText := ""
	for i < blkNum {
		for i < blkNum && blksLen[i] == 0 {
			i++
		}
		if i > blkNum {
			break
		}
		curTextLen, portion := 0, ""
		for i < blkNum && blksLen[i] > 0 {
			if lines[i] != "" {
				portion += lines[i] + "\n"
				curTextLen += len(lines[i])
			}
			i++
		}
		if curTextLen > max {
			plainText = portion
			max = curTextLen
		}
	}
	return plainText
}
