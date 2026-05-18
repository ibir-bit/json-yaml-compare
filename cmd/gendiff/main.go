package main

import (
	"fmt"
	"log"
	"os"

	"code/pkg/gendiff" // Убедитесь, что "code" совпадает с именем в go.mod
	"code/pkg/parser"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return cli.Exit("Usage: gendiff <file1> <file2>", 1)
			}

			data1, err := parser.ReadFile(c.Args().Get(0))
			if err != nil {
				return err
			}
			data2, err := parser.ReadFile(c.Args().Get(1))
			if err != nil {
				return err
			}

			result := gendiff.GenDiffRecursive(data1, data2)
			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
