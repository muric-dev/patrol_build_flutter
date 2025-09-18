package commands

// / This class is used to define the commands that will be executed in the terminal.
type Command struct {
	Name string
	Args []string
}

// / Get pub dependencies in compact format
var FlutterPubDependencies = Command{
	Name: "flutter",
	Args: []string{"pub", "deps", "--style=compact"},
}

// / Get patrol verbose with extra information
var PatrolDoctor = Command{
	Name: "patrol",
	Args: []string{"doctor", "--verbose"},
}

// / This command is used to install the patrol_cli package globally using pub
var PatrolInstall = Command{
	Name: "dart",
	Args: []string{"pub", "global", "activate", "patrol_cli"},
}

var CreatePatrolFolder = Command{
	Name: "mkdir",
	Args: []string{"patrol"},
}

var CopyBuildsToFolder = Command{
	Name: "cp",
	Args: []string{},
}

var CompressIOSFiles = Command{
	Name: "zip",
	Args: []string{"-r", "patrol/ios_tests.zip"},
}

func (c Command) CopyWith(name *string, args []string) Command {
	copy := c

	if name != nil {
		copy.Name = *name
	}

	if args != nil {
		copy.Args = args
	}

	return copy
}
