package jsoner

import (
	"bufio"
	"os"
)

func ReadData(filename string) ([]byte, error) {
    file, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

    stats, err := file.Stat()
	if err != nil {
		return []byte{}, err
	}

	size := stats.Size()
	bytes := make([]byte, size)
	if _, err := bufio.NewReader(file).Read(bytes); err != nil {
	    return []byte{}, err
	}

	return bytes, nil
}

func WriteData(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0666); err != nil {
		return err
	}
	return nil
}
