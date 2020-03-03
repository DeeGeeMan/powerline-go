package main

import (
	pwl "github.com/justjanne/powerline-go/powerline"
	"os"
)

func segmentShellVar(p *powerline) {
	shellVarName := *p.args.ShellVar
	varContent, varExists := os.LookupEnv(shellVarName)

	if varExists {
		if varContent != "" {
			p.appendSegment("shell-var", pwl.Segment{
				Content:    varContent,
				Foreground: p.theme.ShellVarFg,
				Background: p.theme.ShellVarBg,
			})
		}
	}
}
