package main

import (
	"fmt"
	"log"
	"os"

	"code/pkg/gendiff"

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

			path1 := c.Args().Get(0)
			path2 := c.Args().Get(1)
			format := c.String("format")

			// Передаем только пути
			result, err := gendiff.GenDiffRecursive(data1, data2)
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
