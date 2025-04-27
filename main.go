package main

import (
	"bufio"
	"fmt"
	"github.com/wvan1901/Culer/internal"
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
	// String color replacer
	re := createReplacer(flags)

	// Listens for inputs and ends once we reach EOF
	for input.Scan() {
		lineText := input.Text()
		newLine := prefix(flags) + " " + re.Replace(lineText)
		s := fmt.Sprintf("%s\n", newLine)

		// NOTE: Ignoring outputs
		output.WriteString(s)
		// Flushed buffered output
		output.Flush()
	}

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

// Create a string replacer for all options
func createReplacer(f internal.Flag) *strings.Replacer {
	replaceStrings := []string{}
	infoColored := internal.ColorString(f.InfoBgColor, "black", f.InfoReplaceStr)
	errColored := internal.ColorString(f.ErrorBgColor, "black", f.ErrorReplaceStr)
	debugColored := internal.ColorString(f.DebugBgColor, "black", f.DebugReplaceStr)
	replaceStrings = append(replaceStrings, f.InfoReplaceStr, infoColored)
	replaceStrings = append(replaceStrings, f.ErrorReplaceStr, errColored)
	replaceStrings = append(replaceStrings, f.DebugReplaceStr, debugColored)
	for _, s := range f.ExtraStrings {
		extraColored := internal.ColorString(s.BackgroundColor, s.ForegroundColor, s.StringToColor)
		replaceStrings = append(replaceStrings, s.StringToColor, extraColored)
	}

	return strings.NewReplacer(replaceStrings...)
}
