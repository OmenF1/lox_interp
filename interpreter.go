package main

import (
	"bufio"
	"fmt"
	"io"
	structs "lox/interpreter/structs"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: interpreter [script-name]")
		return
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}

}

func runFile(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file!", err)
		return
	}
	run(string(data))
	// //	if we don't wrap in string we get ascii value array.
	// fmt.Println("File contents:", string(data))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		run(input)
	}
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Println("[line ", line, "] error ", where, ": ", message)
	// hadError = true
}

func run(source string) {
	scanner := structs.Scanner{
		Source: source,
	}

	var tokens = scanner.Tokens
	// Scanner scanner = new Scanner(source);

	fmt.Println(tokens)
}
