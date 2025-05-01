package internal

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

const (
	PROGRAM_NAME     = "Culer"
	INFO             = "INFO"
	DEBUG            = "DEBUG"
	ERROR            = "ERROR"
	DEFAULT_BG_COLOR = "black"
	INFO_FG_COLOR    = "light-blue"
	DEBUG_FG_COLOR   = "light-yellow"
	ERROR_FG_COLOR   = "light-red"
)

type Flag struct {
	EnableTimeStamp bool
	ProgramName     string
	PrefixBgColor   string
	PrefixFgColor   string
	InfoReplaceStr  string
	ErrorReplaceStr string
	DebugReplaceStr string
	InfoBgColor     string
	InfoFgColor     string
	ErrorBgColor    string
	ErrorFgColor    string
	DebugBgColor    string
	DebugFgColor    string
	ExtraStrings    stringsToColor
}

func InitFlags(args []string) Flag {
	fs := flag.NewFlagSet(PROGRAM_NAME, flag.ContinueOnError)

	enableTime := fs.Bool("time", false, "Add time stamp to prefix")
	programName := fs.String("program-name", PROGRAM_NAME, "Program name, this gets displayed in prefix")
	prefixBgColor := fs.String("prefix-bg-color", "black", "Prefix background color")
	prefixFgColor := fs.String("prefix-fg-color", "light-green", "Prefix foreground color")
	infoStr := fs.String("info-str", INFO, "Value of the info string to color")
	errStr := fs.String("err-str", ERROR, "Value of the error string to color")
	debugStr := fs.String("debug-str", DEBUG, "Value of the debug string to color")
	infoBgColor := fs.String("info-bg-color", DEFAULT_BG_COLOR, "background color of the info string")
	infoFgColor := fs.String("info-fg-color", INFO_FG_COLOR, "foreground color of the info string")
	errBgColor := fs.String("err-bg-color", DEFAULT_BG_COLOR, "background color of the error string")
	errFgColor := fs.String("err-fg-color", ERROR_FG_COLOR, "foreground color of the error string")
	debugBgColor := fs.String("debug-bg-color", DEFAULT_BG_COLOR, "background color of the debug string")
	debugFgColor := fs.String("debug-fg-color", DEBUG_FG_COLOR, "foreground color of the debug string")
	extraStrOpt := &stringsToColor{}
	fs.Var(extraStrOpt, "extra-str", "Advanced option to color any additional custom string, must follow this format -> [stringValue]&[back ground color]:[foreground color] -> Ex: replace-me&red:black -> This will replace all strings 'replace-me' with a red background and a black foreground")

	fs.Parse(args)

	return Flag{
		EnableTimeStamp: *enableTime,
		ProgramName:     *programName,
		PrefixBgColor:   *prefixBgColor,
		PrefixFgColor:   *prefixFgColor,
		InfoReplaceStr:  *infoStr,
		ErrorReplaceStr: *errStr,
		DebugReplaceStr: *debugStr,
		InfoBgColor:     *infoBgColor,
		InfoFgColor:     *infoFgColor,
		ErrorBgColor:    *errBgColor,
		ErrorFgColor:    *errFgColor,
		DebugBgColor:    *debugBgColor,
		DebugFgColor:    *debugFgColor,
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
