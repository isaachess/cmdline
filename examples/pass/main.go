package main

import (
	"cmdline"
)

func main() {
	var pws = []*pw{
		{
			site: "gmail.com",
			pw:   "passwordgmail",
		},
		{
			site: "apple.com",
			pw:   "passwordapple",
		},
		{
			site: "amazon.com",
			pw:   "passwordamazon",
		},
	}
	cmd := cmdline.NewCmd("pass", nil, nil, nil)
	cmdCp := cmdline.NewCmd(cmdCpName, &cmdCpRunner{pws: pws}, cmdCpArgs, cmdCpFlags)
	cmdCreate := cmdline.NewCmd(cmdCreateName, &cmdCreateRunner{pws: pws}, cmdCreateArgs, cmdCreateFlags)
	cmdDelete := cmdline.NewCmd(cmdDeleteName, &cmdDeleteRunner{pws: pws}, cmdDeleteArgs, cmdDeleteFlags)

	cmdListVerbose := cmdline.NewCmd(cmdListVerboseName, &cmdListVerboseRunner{pws: pws}, cmdListVerboseArgs, cmdListVerboseFlags)
	cmdListSimple := cmdline.NewCmd(cmdListSimpleName, &cmdListSimpleRunner{pws: pws}, cmdListSimpleArgs, cmdListSimpleFlags)
	cmdList := cmdline.NewCmd("list", nil, nil, nil)
	cmdList.RegisterSub(cmdListVerbose)
	cmdList.RegisterSub(cmdListSimple)

	cmd.RegisterSub(cmdCp)
	cmd.RegisterSub(cmdCreate)
	cmd.RegisterSub(cmdList)
	cmd.RegisterSub(cmdDelete)
	cmdline.Execute(cmd)
}
