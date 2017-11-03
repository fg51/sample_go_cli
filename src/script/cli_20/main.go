package main

import (
	"fmt"
	"os"
)
import (
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_20"
	app.Usage = "cli_20's sample"
	app.Version = "1.2.1"

	app.Before = func(c *cli.Context) error {
		fmt.Println("-- Before --")
		return nil
	}


	app.Action = func(c *cli.Context) error {
		fmt.Println("-- Action --")

		fmt.Printf("c.NArg()				: %+v\n", c.NArg())
		fmt.Printf("c.Args()				: %+v\n", c.Args())
		fmt.Printf("c.Args().Get(0) : %+v\n", c.Args().Get(0))
		fmt.Printf("c.Args()[0]     : %+v\n", c.Args()[0])
		fmt.Printf("c.FlagNames		  : %+v\n", c.FlagNames())

		// cli.ShowAppHelp(c)

		cli.ShowVersion(c)


		return nil
	}

	app.After = func(c *cli.Context) error {
		fmt.Println("-- After --")
		return nil
	}

	app.HideHelp = false
	app.Run(os.Args)
}
