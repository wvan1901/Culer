package internal

import (
	"flag"
)

type Flag struct {
	EnableTimeStamp bool
	ProgramName     string
}

func InitFlags() Flag {
	enableTime := flag.Bool("time", false, "Decided to wrap line by a time stamp")
	programName := flag.String("program-name", "Culer", "Program name being wrapped")

	flag.Parse()

	return Flag{
		EnableTimeStamp: *enableTime,
		ProgramName:     *programName,
	}
}

// TODO: Add flag to select a background color for prefix
// TODO: Add flag that enable custom string for info, error, debug
// TODO: Find solution to enable custom colors (Flags?)
// TODO: Find solution to have custom string also be colored
