package jsoner

import (
	"bufio"
	"log"
	"os"
)

//main read function
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

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}

//main write function
func WriteData(filename string, data []byte) {
	err := os.WriteFile(filename, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// func main() {

// 	//writeFile()
// 	//readFile()
// 	bytedata := []byte(`{"header1":"data about header1","body1":"data about body1"},{"header2":"data about header2","body2":"data about body2"}`)
// 	WriteData("data.bin", bytedata)
// 	readdata, err := ReadData("data.bin")
// 	if err != nil {
// 		fmt.Println("error reading bin file")
// 	}
// 	str := string(readdata[:])
// 	fmt.Println(str)

// }
