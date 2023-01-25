package regex

import (
	PackageRegex "github.com/Eclalang/regex"
	"reflect"
	"regexp"
	"testing"
)

func TestFind(t *testing.T) {
	regex := "a"
	str := "anunaban"
	comp := regexp.MustCompile(regex)
	expected := comp.FindString(str)
	actual := PackageRegex.Find(regex, str)
	if actual != expected {
		t.Errorf("Find(\"%s\", \"%s\") returned %s, expected %s", regex, str, actual, expected)
	}
}

func TestFindAll(t *testing.T) {
	regex := "a"
	str := "anunaban"
	comp := regexp.MustCompile(regex)
	expected := comp.FindAllString(str, -1)
	actual := PackageRegex.FindAll(regex, str)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FindAll(\"%s\", \"%s\") returned %s, expected %s", regex, str, actual, expected)
	}
}

func TestFindAllIndex(t *testing.T) {
	regex := "a"
	str := "anunaban"
	comp := regexp.MustCompile(regex)
	expected := comp.FindAllStringIndex(str, -1)
	actual := PackageRegex.FindAllIndex(regex, str)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FindAllIndex(\"%s\", \"%s\") returned %v, expected %v", regex, str, actual, expected)
	}
}

func TestFindIndex(t *testing.T) {
	regex := "a"
	str := "anunaban"
	comp := regexp.MustCompile(regex)
	expected := comp.FindStringIndex(str)
	actual := PackageRegex.FindIndex(regex, str)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FindIndex(\"%s\", \"%s\") returned %v, expected %v", regex, str, actual, expected)
	}
}

func TestMatch(t *testing.T) {
	regex := "mundo"
	str := "mundointheworld"
	comp := regexp.MustCompile(regex)
	expected := comp.MatchString(str)
	actual := PackageRegex.Match(regex, str)
	if actual != expected {
		t.Errorf("Match(\"%s\", \"%s\") returned %t, expected %t", regex, str, actual, expected)
	}
}

func TestReplaceAll(t *testing.T) {
	regex := "a"
	str := "anunaban"
	new := "b"
	comp := regexp.MustCompile(regex)
	expected := comp.ReplaceAllString(str, new)
	actual := PackageRegex.ReplaceAll(regex, str, new)
	if actual != expected {
		t.Errorf("ReplaceAll(\"%s\", \"%s\", \"%s\") returned %s, expected %s", regex, str, new, actual, expected)
	}
}
