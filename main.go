package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"time"
)

func main() {
	app := &cli.App{
		Name:  "miniattack",
		Usage: "Execute a specified command at regular intervals",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "interval-ms",
				Value: 1000,
				Usage: "Interval between command executions in milliseconds",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return cli.NewExitError("Please provide a command to execute.", 1)
			}

			interval := time.Duration(c.Int("interval-ms")) * time.Millisecond
			ticker := time.NewTicker(interval)
			defer ticker.Stop()

			command := c.Args().Get(0)
			args := c.Args()[1:]

			for {
				select {
				case <-ticker.C:
					if err := runCommand(command, args); err != nil {
						fmt.Printf("Error executing command: %v\n", err)
					}
				}
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func runCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
