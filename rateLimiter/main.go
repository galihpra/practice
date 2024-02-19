package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func rateLimiter(inputFile, outputFile string) {
	var timeLimit = 1 * time.Second
	var maxRequests = 5

	input, _ := os.Open(inputFile)
	defer input.Close()

	output, _ := os.Create(outputFile)
	defer output.Close()

	read := bufio.NewScanner(input)

	write := bufio.NewWriter(output)
	defer write.Flush()

	requestCount := make(map[string]int)

	lastRequestTime := make(map[string]time.Time)

	for read.Scan() {
		request := read.Text()
		split := strings.Split(request, ",")

		clientID := split[0]

		lastTime, ok := lastRequestTime[clientID]
		if ok && time.Since(lastTime) >= timeLimit {
			delete(requestCount, clientID)
		}

		requestCount[clientID]++

		if requestCount[clientID] > maxRequests {
			fmt.Fprintln(write, "BLOCKED:")
		} else {
			fmt.Fprintln(write, "PASSED:")
		}
	}
}

func main() {
	inputFile := "fileInput.log"
	outputFile := "fileOutput.log"

	rateLimiter(inputFile, outputFile)
}
