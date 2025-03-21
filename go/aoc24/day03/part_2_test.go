package day03_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day03"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	const sourceData = "{\"id\": 10, \"name\": \"Pie\"}"
	var data, _ = io.ReadAll(day03.NewDoReader(strings.NewReader(sourceData)))

	if i := strings.Compare(string(data), sourceData); i != 0 {
		t.Errorf(`Expected %s to match %s`, string(data), sourceData)
	}
}
func TestHelloName2(t *testing.T) {
	const sourceData = "asdfdo()yodon't()nodo()asdf"
	const targetData = "asdfyoasdf"
	var data, _ = io.ReadAll(day03.NewDoReader(strings.NewReader(sourceData)))

	if i := strings.Compare(string(data), targetData); i != 0 {
		t.Errorf(`Expected %s to match %s`, string(data), targetData)
	}
}
