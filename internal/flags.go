package internal

import (
	"flag"
)

const (
	PROGRAM_NAME = "Culer"
)

type Flag struct {
	EnableTimeStamp bool
	ProgramName     string
}

func InitFlags(args []string) Flag {
	fs := flag.NewFlagSet(PROGRAM_NAME, flag.ContinueOnError)

	enableTime := fs.Bool("time", false, "Wrap line by a time stamp")
	programName := fs.String("program-name", PROGRAM_NAME, "Program name being wrapped")

	fs.Parse(args)

	return Flag{
		EnableTimeStamp: *enableTime,
		ProgramName:     *programName,
	}
}

// TODO: Add flag to select a background color for prefix
// TODO: Add flag that enable custom string for info, error, debug
// TODO: Find solution to enable custom colors (Flags?)
// TODO: Find solution to have custom string also be colored
