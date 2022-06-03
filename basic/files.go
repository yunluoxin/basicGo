package basic

import (
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}

func WriteToFile(content string, filename string) error {
	err := os.WriteFile(filename, []byte(content), 0600)
	return err
}