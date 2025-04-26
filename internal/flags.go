package internal

import (
	"flag"
)

const (
	PROGRAM_NAME = "Culer"
	INFO         = "INFO"
	DEBUG        = "DEBUG"
	ERROR        = "ERROR"
)

type Flag struct {
	EnableTimeStamp bool
	ProgramName     string
	PrefixBgColor   string
	InfoReplaceStr  string
	ErrorReplaceStr string
	DebugReplaceStr string
}

func InitFlags(args []string) Flag {
	fs := flag.NewFlagSet(PROGRAM_NAME, flag.ContinueOnError)

	enableTime := fs.Bool("time", false, "Add time stamp to prefix")
	programName := fs.String("program-name", PROGRAM_NAME, "Program name displayed in prefix")
	prefixBgColor := fs.String("prefix-bg-color", "light-green", "Prefix background color")
	infoStr := fs.String("info-str", INFO, "Value of the info string to color")
	errStr := fs.String("err-str", ERROR, "Value of the info string to color")
	debugStr := fs.String("debug-str", DEBUG, "Value of the debug string to color")

	fs.Parse(args)

	return Flag{
		EnableTimeStamp: *enableTime,
		ProgramName:     *programName,
		PrefixBgColor:   *prefixBgColor,
		InfoReplaceStr:  *infoStr,
		ErrorReplaceStr: *errStr,
		DebugReplaceStr: *debugStr,
	}
}

// TODO: Find solution to enable custom colors for Info, Debug, Error
// TODO: Find solution to have custom string & it to be colored
