package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	
	if err != nil {
		return err
	}
	
	defer file.Close()
	file.WriteString(message)

	return nil
}

func appendToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString("\n" + message)

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var message string

	for {
		line, _, err := reader.ReadLine()
		if (err == io.EOF) {
			break
		}

		// fmt.Println(line)
		message += string(line) + "\n"
	}

	return message, nil
}

func main() {
	// createNewFile("sample.log", "this is sample log")
	// result, _ := readFile("sample.log")
	// fmt.Println(result)

	appendToFile("sample.log", "this is added message")
}