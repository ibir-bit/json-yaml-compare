package main

import (
	"flag"
	"fmt"
	"os"

	"code" // Импортируем модуль "code" из корня
)

func main() {
	format := flag.String("format", "stylish", "output format [stylish, plain, json]")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: gendiff [--format <format>] <file1> <file2>")
		os.Exit(1)
	}

	file1 := flag.Arg(0)
	file2 := flag.Arg(1)

	// Вызываем функцию через code.GenDiff
	result, err := code.GenDiff(file1, file2, *format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
