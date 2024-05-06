package main

import (
	"regexp"
	"strings"
)

func Punctuation(text string) string {

	re1 := regexp.MustCompile(` +([,.!?;:])`)
	text = re1.ReplaceAllString(text, "$1")

	re2 := regexp.MustCompile(`([,.!?;:])([^,.'!?;:\s])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	return text
}

func Apostrophe(text string) string {
	lines := strings.Split(text, "\n")
	newLines := []string{}
	for _, line := range lines {

		re1 := regexp.MustCompile(`('\s+)`)
		line = re1.ReplaceAllString(line, " $1")
		re2 := regexp.MustCompile(`(\s+')`)
		line = re2.ReplaceAllString(line, "$1 ")
		re3 := regexp.MustCompile(`\A'`)
		line = re3.ReplaceAllString(line, " ' ")
		re4 := regexp.MustCompile(`'$`)
		line = re4.ReplaceAllString(line, " ' ")

		re5 := regexp.MustCompile(`'\s+`)
		count := 0
		line = re5.ReplaceAllStringFunc(line, func(match string) string {
			if count%2 == 0 {
				count++
				return "'"
			} else {
				count++
				return match
			}
		})
		count = 0

		re6 := regexp.MustCompile(`\s+'`)
		line = re6.ReplaceAllStringFunc(line, func(match string) string {
			if count%2 == 1 {
				count++
				return "'"
			} else {
				count++
				return match
			}
		})

		re7 := regexp.MustCompile(`[ ]+'`)
		line = re7.ReplaceAllString(line, " '")
		re8 := regexp.MustCompile(`'[ ]+`)
		line = re8.ReplaceAllString(line, "' ")
		re9 := regexp.MustCompile(`\A '`)
		line = re9.ReplaceAllString(line, "'")

		newLines = append(newLines, strings.TrimRight(line, " \t"))
	}
	text = strings.Join(newLines, "\n")

	return strings.TrimRight(text, " \t")
}

func Grammar(text string) string {
	re := regexp.MustCompile(`(?i)(\ba)( +)([aeiouh])`)
	for re.MatchString(text) {
		text = re.ReplaceAllString(text, "${1}n$2$3")
	}
	return text
}
