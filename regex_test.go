package regex

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	pattern := "[a-z]*"
	str := "abxd"
	comp := regexp.MustCompile(pattern)
	expected := comp.FindString(str)
	actual := Find(pattern, str)
	if actual != expected {
		t.Errorf("Find(\"%s\", \"%s\") returned %s, expected %s", pattern, str, actual, expected)
	}
}

func TestFindAll(t *testing.T) {
	pattern := "a"
	str := "anunaban"
	comp := regexp.MustCompile(pattern)
	expected := comp.FindAllString(str, -1)
	actual := FindAll(pattern, str)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FindAll(\"%s\", \"%s\") returned %s, expected %s", pattern, str, actual, expected)
	}
}

func TestFindAllIndex(t *testing.T) {
	pattern := "a"
	str := "anunaban"
	comp := regexp.MustCompile(pattern)
	expectedTempo := comp.FindAllStringIndex(str, -1)
	var expected []int
	for _, match := range expectedTempo {
		expected = append(expected, match[0])
	}
	actual := FindAllIndex(pattern, str)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FindAllIndex(\"%s\", \"%s\") returned %v, expected %v", pattern, str, actual, expected)
	}
}

func TestFindIndex(t *testing.T) {
	pattern := "a"
	str := "anunaban"
	comp := regexp.MustCompile(pattern)
	expected := comp.FindStringIndex(str)
	actual := FindIndex(pattern, str)
	if !reflect.DeepEqual(actual, expected[0]) {
		t.Errorf("FindIndex(\"%s\", \"%s\") returned %v, expected %v", pattern, str, actual, expected)
	}
}

func TestMatch(t *testing.T) {
	pattern := "mundo"
	str := "mundointheworld"
	comp := regexp.MustCompile(pattern)
	expected := comp.MatchString(str)
	actual := Match(pattern, str)
	if actual != expected {
		t.Errorf("Match(\"%s\", \"%s\") returned %t, expected %t", pattern, str, actual, expected)
	}
}

func TestReplace(t *testing.T) {
	pattern := "a"
	str := "anunaban"
	newStr := "AA"
	nb := 2
	comp := regexp.MustCompile(pattern)
	expected := strings.Replace(str, comp.FindString(str), newStr, nb)
	actual := Replace(pattern, str, newStr, nb)
	if actual != expected {
		t.Errorf("Replace(\"%s\", \"%s\", \"%s\") returned %s, expected %s", pattern, str, newStr, actual, expected)
	}
}

func TestReplaceAll(t *testing.T) {
	pattern := "a"
	str := "anunaban"
	newStr := "AA"
	comp := regexp.MustCompile(pattern)
	expected := comp.ReplaceAllString(str, newStr)
	actual := ReplaceAll(pattern, str, newStr)
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("ReplaceAll(\"%s\", \"%s\", \"%s\") returned %s, expected %s", pattern, str, newStr, actual, expected)
	}
}
