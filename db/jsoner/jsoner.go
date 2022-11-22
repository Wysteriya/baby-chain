package jsoner

import (
	"bufio"
	"log"
	"os"
)

func ReadData(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	return bytes, err
}

func WriteData(filename string, data []byte) {
	err := os.WriteFile(filename, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
