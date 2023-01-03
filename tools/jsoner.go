package tools

import (
	"bufio"
	"fmt"
	"os"
)

func ReadData(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	stats, err := file.Stat()
	if err != nil {
		panic(err)
	}

	size := stats.Size()
	bytes := make([]byte, size)
	if _, err := bufio.NewReader(file).Read(bytes); err != nil {
		panic(err)
	}

	if err := file.Close(); err != nil {
		fmt.Println(err)
	}

	return bytes
}

func WriteData(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0666); err != nil {
		return err
	}
	return nil
}
