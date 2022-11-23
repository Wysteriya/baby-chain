package jsoner

import (
	"bufio"
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

func WriteData(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0666); err != nil {
		return err
	}
	return nil
}
