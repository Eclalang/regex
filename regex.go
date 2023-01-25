package regex

import "regexp"

// Find returns the first match of the regex in the string
func Find(regex, str string) string {
	re := regexp.MustCompile(regex)
	return re.FindString(str)
}

// FindAll returns all matches of the regex in the string
func FindAll(regex, str string) []string {
	re := regexp.MustCompile(regex)
	return re.FindAllString(str, -1)
}

// FindAllIndex returns the indexes of all matches of the regex in the string
func FindAllIndex(regex, str string) [][]int {
	re := regexp.MustCompile(regex)
	return re.FindAllStringIndex(str, -1)
}

// FindIndex returns the first and last index of the first match of the regex in the string
func FindIndex(regex, str string) []int {
	re := regexp.MustCompile(regex)
	return re.FindStringIndex(str)
}

// Match returns true if the regex matches the string
func Match(regex, str string) bool {
	re := regexp.MustCompile(regex)
	return re.MatchString(str)
}

// ReplaceAll replaces all matches of the regex in the string with the new string
func ReplaceAll(regex, str, new string) string {
	re := regexp.MustCompile(regex)
	return re.ReplaceAllString(str, new)
}
