package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	bufferSize         = 32 * 1024 // 32kb
	separatorCharacter = '\n'      // newline character
)

func main() {
	if len(os.Args) != 2 {
		executableName := os.Args[0]
		fmt.Println("usage:", executableName, "[FILE]")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	linesCount, err := CountLinesInFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(linesCount, filePath)
}

func CountLinesInFile(source io.Reader) (int, error) {
	buffer := make([]byte, bufferSize)
	linesCount := 0
	separator := []byte{separatorCharacter}

	for {
		charsCount, err := source.Read(buffer)
		linesCount += bytes.Count(buffer[:charsCount], separator)

		if err == io.EOF {
			return linesCount, nil
		} else if err != nil {
			return linesCount, err
		}
	}
}
