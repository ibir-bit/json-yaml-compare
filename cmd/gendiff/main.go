package main

import (
	"code" // Импорт нашего корневого пакета
	"flag"
	"fmt"
	"os"
)

func main() {
	format := flag.String("format", "stylish", "set format of output")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: gendiff [-format <format>] <firstConfig> <secondConfig>")
		os.Exit(1)
	}

	diff, err := code.GenDiff(flag.Arg(0), flag.Arg(1), *format)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(diff)
}
