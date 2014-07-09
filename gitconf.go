package main

import (
  "github.com/codegangsta/cli"
  "os"
)

func main() {
  app := cli.NewApp()
  app.Name = "gitconf"
  app.Version = Version
  app.Usage = ""
  app.Author = "fumitoito"
  app.Email = "weathercook@gmail.com"
  app.Commands = Commands

  app.Run(os.Args)
}