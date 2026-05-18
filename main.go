package main

import (
	"fmt"
	"log"
	"os"

	"code/pkg/gendiff"
	"code/pkg/parser"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		// 1. Добавляем поддержку флага --format (по умолчанию "stylish")
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format [stylish, plain, json]",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return cli.Exit("Usage: gendiff <file1> <file2>", 1)
			}

			// Читаем файлы
			data1, err := parser.ReadFile(c.Args().Get(0))
			if err != nil {
				return err
			}
			data2, err := parser.ReadFile(c.Args().Get(1))
			if err != nil {
				return err
			}

			// 2. Получаем значение флага --format
			format := c.String("format")

			// 3. Передаем data1, data2 и format в GenDiff
			result, err := gendiff.GenDiff(data1, data2, format)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
