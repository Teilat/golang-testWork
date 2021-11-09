package main

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	testArray := make(map[string]string)
	testArray["a"] = "2"
	testArray["b"] = "abc"
	testArray["d"] = "43f"

	array := ReadFromFile("test1.txt")

	eq := reflect.DeepEqual(testArray, array)
	if !eq {
		t.Error("TestRead want:", testArray, "\ngot:", array)
	}
}

func TestWrite(t *testing.T) {
	testArray := make(map[string]string)
	testArray["a"] = "2"
	testArray["b"] = "abc"
	testArray["d"] = "43f"

	WriteInFile(testArray, "test2.txt")

	file, _ := os.OpenFile("test2.txt", os.O_RDONLY, 0660)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	array := make(map[string]string)
	for scanner.Scan() {
		var str []string = strings.Split(scanner.Text(), "=")
		array[str[0]] = str[1]
	}
	eq := reflect.DeepEqual(testArray, array)
	if !eq {
		t.Error("TestWrite want:", testArray, "\ngot:", array)
	}
}
