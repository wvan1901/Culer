package main

import (
	"bufio"
	"culer/internal"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	if err := run(os.Stdin, os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(r io.Reader, w io.Writer, args []string) error {
	flags := internal.InitFlags(args)
	// Input that splits on new lines
	input := bufio.NewScanner(r)
	// Buffered output
	output := bufio.NewWriter(w)

	// Listens for inputs and ends once we reach EOF
	for input.Scan() {
		lineText := input.Text()
		newLine := prefix(flags) + " " + addColor(lineText)
		s := fmt.Sprintf("%s\n", newLine)

		// NOTE: Ignoring output from func
		output.WriteString(s)
	}

	// Flushed remaining buffered output
	output.Flush()

	return nil
}

// Adds prefix to string to differentiate between line logs
func prefix(f internal.Flag) string {
	prefixStr := "[" + f.ProgramName

	if f.EnableTimeStamp {
		curTime := time.Now()
		prefixStr += fmt.Sprintf(" - %s", curTime.Format("15:04:05.000"))
	}

	prefixStr += "]"

	return internal.ColorPrefix(prefixStr, f.PrefixBgColor)
}

// Looks for substrings: INFO, ERROR, DEBUG. Then wraps color around the string
func addColor(s string) string {
	// TODO: Add flag that enable custom string for info, error, debug
	// TODO: Add flag that enable custom colors

	// Color Info
	infoSubStr := "INFO"
	infoIndex := strings.Index(s, infoSubStr)
	if infoIndex != -1 {
		s = s[:infoIndex] + internal.ColorString("blue", "black", infoSubStr) + s[infoIndex+len(infoSubStr):]
	}

	// Color Error
	errSubStr := "ERROR"
	errIndex := strings.Index(s, errSubStr)
	if errIndex != -1 {
		s = s[:errIndex] + internal.ColorString("red", "black", errSubStr) + s[errIndex+len(errSubStr):]
	}

	// Color Debug
	debugSubStr := "DEBUG"
	debugIndex := strings.Index(s, debugSubStr)
	if debugIndex != -1 {
		s = s[:debugIndex] + internal.ColorString("yellow", "black", debugSubStr) + s[debugIndex+len(debugSubStr):]
	}

	return s
}
