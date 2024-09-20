package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"reloaded"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Not enough Arguments!")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many Arguments!")
		return
	} else if !(strings.HasSuffix(args[0], ".txt")) || !(strings.HasSuffix(args[1], ".txt")) {
		fmt.Println("Wrong file name!")
		return
	}

	inputFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(args[1])
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		SplitSpace := reloaded.SplitWhiteSpaces(line)
		FlagsDone := reloaded.HandleTheFlags(SplitSpace)
		HandleVowel := reloaded.VowelSituation(FlagsDone)
		ConvTStr := strings.Join(HandleVowel, " ")
		SecSplit := strings.Split(ConvTStr, " ")
		SecConv := strings.Join(SecSplit, " ")
		HandlePunc := reloaded.HandlePunctuations(SecConv)
		finalResult := reloaded.HandleQuotes(HandlePunc)

		_, err = writer.WriteString(finalResult + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing output file:", err)
		return
	}
}
