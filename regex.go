package regex

import (
	"regexp"
	"strings"
)

// Find returns the first match of the pattern in the string
func Find(pattern, str string) string {
	re := regexp.MustCompile(pattern)
	return re.FindString(str)
}

// FindAll returns all matches of the pattern in the string
func FindAll(pattern, str string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(str, -1)
}

// FindAllIndex returns the indexes of all matches of the pattern in the string
func FindAllIndex(pattern, str string) []int {
	re := regexp.MustCompile(pattern)
	var indexes []int
	for _, match := range re.FindAllStringIndex(str, -1) {
		indexes = append(indexes, match[0])
	}
	return indexes
}

// FindIndex returns the first and last index of the first match of the pattern in the string
func FindIndex(pattern, str string) int {
	re := regexp.MustCompile(pattern)
	return re.FindStringIndex(str)[0]
}

// Match returns true if the pattern matches the string
func Match(pattern, str string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}

// Replace replaces the nb first matches of the pattern in the string with the new string
func Replace(pattern, str, new string, nb int) string {
	re := regexp.MustCompile(pattern)
	return strings.Replace(str, re.FindString(str), new, nb)
}

// ReplaceAll replaces all matches of the pattern in the string with the new string
func ReplaceAll(pattern, str, new string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(str, new)
}
