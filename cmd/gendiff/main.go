package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name: "gendiff",

		Usage: "Compares two configuration files and shows a difference.",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "json",
				Usage:   "string  output format",
			},
		},

		Action: func(_ context.Context, cmd *cli.Command) error {
			if cmd.NArg() == 0 {
				return fmt.Errorf("path is required")
			}

			path := cmd.Args().Get(0)

			data, _ := code.GenDiff(path, "file2.json", cmd.String("format"))

			fmt.Println(data)

			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
