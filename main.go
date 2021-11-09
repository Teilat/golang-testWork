package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	array := ReadFromFile("data.txt")

	if len(os.Args) <= 1 {
		View(array)
	} else {
		switch os.Args[1] {
		case "-add":
			var val string
			var key string
			fmt.Println("use -end to interrupt programm")

			for {

				fmt.Println("enter key")
				fmt.Scan(&key)
				if key == "-end" || val == "-end" {
					break
				}

				fmt.Println("enter value")
				fmt.Scan(&val)
				if key == "-end" || val == "-end" {
					break
				}

				array = add(array, key, val, "data.txt")
			}
		case "-remove":
			var key string
			fmt.Println("use -end to interrupt programm")

			for {

				fmt.Println("enter key")
				fmt.Scan(&key)
				if key == "-end" {
					break
				}

				array = Remove(array, key, "data.txt")

			}

		case "-help":
			fmt.Println("without params write all from file \n-add *key* *val* //add new value \n-remove *key* //remove value by key \n-edit *key* //edit value by key")
		default:
			fmt.Println("unknown command, use -help")
		}
	}

}

func add(array map[string]string, key, val, fileName string) map[string]string {

	array[key] = val //map сам обновит значение если ключ уже существует

	WriteInFile(array, fileName)
	return array
}

func Remove(array map[string]string, key, fileName string) map[string]string {
	var a string
	fmt.Println("Are you shure? \n [y] or [n]")
	fmt.Scan(&a)
	if a == "y" {
		delete(array, key)
	}

	WriteInFile(array, fileName)
	return array
}

func View(array map[string]string) {

	for key, value := range array {
		var str string = key + "=" + value
		fmt.Println(str)
	}
}

func WriteInFile(array map[string]string, fileName string) {

	os.Truncate(fileName, 0)
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0660)

	for key, value := range array {
		var str string = key + "=" + value + "\n"
		file.WriteString(str)
	}

	file.Close()
}

func ReadFromFile(fileName string) map[string]string {

	file, _ := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0660)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	array := make(map[string]string)
	for scanner.Scan() {
		var str []string = strings.Split(scanner.Text(), "=")
		array[str[0]] = str[1]
	}

	file.Close()
	return array
}
