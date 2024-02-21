package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(num uint32) int32 {
	var bin []int32

	for {
		mod := int32(num % 2)
		bin = append(bin, mod)
		num = num / 2

		if num == 0 {
			break
		}
	}

	var result int32

	for i := 0; i < len(bin); i++ {
		if bin[i] == 1 {
			result++
		}
	}

	return result
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	numTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	num := uint32(numTemp)

	result := countBits(num)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
