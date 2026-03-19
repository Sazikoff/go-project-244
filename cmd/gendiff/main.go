package main

import (
	// "code"
	// "fmt"
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	// result, _ := code.GenDiff("filepath1", "filepath2", "format")
	// fmt.Println(result)
}
