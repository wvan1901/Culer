package internal

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

const (
	PROGRAM_NAME   = "Culer"
	INFO           = "INFO"
	DEBUG          = "DEBUG"
	ERROR          = "ERROR"
	INFO_BG_COLOR  = "blue"
	DEBUG_BG_COLOR = "yellow"
	ERROR_BG_COLOR = "red"
)

type Flag struct {
	EnableTimeStamp bool
	ProgramName     string
	PrefixBgColor   string
	InfoReplaceStr  string
	ErrorReplaceStr string
	DebugReplaceStr string
	InfoBgColor     string
	ErrorBgColor    string
	DebugBgColor    string
	ExtraStrings    stringsToColor
}

func InitFlags(args []string) Flag {
	fs := flag.NewFlagSet(PROGRAM_NAME, flag.ContinueOnError)

	enableTime := fs.Bool("time", false, "Add time stamp to prefix")
	programName := fs.String("program-name", PROGRAM_NAME, "Program name, this gets displayed in prefix")
	prefixBgColor := fs.String("prefix-bg-color", "light-green", "Prefix background color")
	infoStr := fs.String("info-str", INFO, "Value of the info string to color")
	errStr := fs.String("err-str", ERROR, "Value of the error string to color")
	debugStr := fs.String("debug-str", DEBUG, "Value of the debug string to color")
	infoBgStr := fs.String("info-color", INFO_BG_COLOR, "background color of the info string")
	errBgStr := fs.String("err-color", ERROR_BG_COLOR, "background color of the error string")
	debugBgStr := fs.String("debug-color", DEBUG_BG_COLOR, "background color of the debug string")
	extraStrOpt := &stringsToColor{}
	fs.Var(extraStrOpt, "extra-str", "Advanced option to color any additional custom string, must follow this format -> [stringValue]&[back ground color]:[foreground color] -> Ex: replace-me&red:black -> This will replace all strings 'replace-me' with a red background and a black foreground")

	fs.Parse(args)

	return Flag{
		EnableTimeStamp: *enableTime,
		ProgramName:     *programName,
		PrefixBgColor:   *prefixBgColor,
		InfoReplaceStr:  *infoStr,
		ErrorReplaceStr: *errStr,
		DebugReplaceStr: *debugStr,
		InfoBgColor:     *infoBgStr,
		ErrorBgColor:    *errBgStr,
		DebugBgColor:    *debugBgStr,
		ExtraStrings:    *extraStrOpt,
	}
}

type colorStringOption struct {
	StringToColor   string
	BackgroundColor string
	ForegroundColor string
}

func (c colorStringOption) IsValid() error {
	if c.StringToColor == "" {
		return errors.New("empty value: StringToColor")
	}
	if c.BackgroundColor == "" {
		return errors.New("empty value: BackgroundColor")
	}
	if c.ForegroundColor == "" {
		return errors.New("empty value: ForegroundColor")
	}
	return nil
}

type stringsToColor []colorStringOption

func (s *stringsToColor) String() string {
	return fmt.Sprintf("%s", *s)
}
func (s *stringsToColor) Set(value string) error {
	if len(value) < 4 {
		return errors.New("value is too short")
	}

	parts := strings.Split(value, "&")
	if len(parts) != 2 {
		return errors.New("exactly one '&' is needed")
	}

	colorParts := strings.Split(parts[1], ":")
	if len(colorParts) != 2 {
		return errors.New("exactly ':' is needed")
	}

	newOpt := colorStringOption{
		StringToColor:   parts[0],
		BackgroundColor: colorParts[0],
		ForegroundColor: colorParts[1],
	}

	if err := newOpt.IsValid(); err != nil {
		return err
	}

	*s = append(*s, newOpt)

	return nil
}
