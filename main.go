package main

import (
	cmd "github.com/berttejeda/go389/cmd"
	model "github.com/berttejeda/go389/model"
	"gopkg.in/codegangsta/cli.v1"
	"os"
)
func main() {
	app := cli.NewApp()
	app.Name = model.ProgramName
	app.Version = "0.1.0"
	app.Author = "kernel164"
	app.Usage = "A simple LDAP server"
	app.Commands = []cli.Command{
		cmd.CmdServer,
		cmd.CmdHash,
	}
	app.Run(os.Args)
}
