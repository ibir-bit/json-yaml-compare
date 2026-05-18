package main

import (
	"fmt"
	"log"
	"os"

	"code/pkg/gendiff"
	// Парсер больше не нужен напрямую в main.go, так как GenDiff парсит сам

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
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

			// Получаем пути к файлам из аргументов
			path1 := c.Args().Get(0)
			path2 := c.Args().Get(1)

			// Получаем значение флага --format
			format := c.String("format")

			// ПЕРЕДАЕМ ПУТИ (path1, path2), а не распарсенные данные
			result, err := gendiff.GenDiff(path1, path2, format)
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
