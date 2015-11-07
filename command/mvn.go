package command

import (
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdMvn(c *cli.Context) {
  args := append([]string{"run", "--rm", "mvn"}, c.Args()...)
  helpers.RunCommand(dockerComposeCommand(c, args...))
}
