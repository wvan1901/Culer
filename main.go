package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/wvan1901/Culer/internal"
)

func main() {
	sigChan := make(chan os.Signal, 1)   // creating a buffered channel and register it to catch sigint (ctrl+c).
	signal.Notify(sigChan, os.Interrupt) // calling signal.Notify prevents Go from exiting on the signal automatically.

	// Listen for signals in the background
	go func() {
		<-sigChan // First Ctrl+C: Catch it and do nothing, letting main() finish reading stdin
		<-sigChan // Second Ctrl+C: Force exit immediately
		os.Exit(130)
	}()

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

	if err := input.Err(); err != nil {
		return err
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

	return internal.ColorPrefix(prefixStr, f.PrefixBgColor, f.PrefixFgColor)
}

// Create a string replacer for all options
func createReplacer(f internal.Flag) *strings.Replacer {
	replaceStrings := []string{}
	infoColored := internal.ColorString(f.InfoBgColor, f.InfoFgColor, f.InfoReplaceStr)
	errColored := internal.ColorString(f.ErrorBgColor, f.ErrorFgColor, f.ErrorReplaceStr)
	debugColored := internal.ColorString(f.DebugBgColor, f.DebugFgColor, f.DebugReplaceStr)
	replaceStrings = append(replaceStrings, f.InfoReplaceStr, infoColored)
	replaceStrings = append(replaceStrings, f.ErrorReplaceStr, errColored)
	replaceStrings = append(replaceStrings, f.DebugReplaceStr, debugColored)
	for _, s := range f.ExtraStrings {
		extraColored := internal.ColorString(s.BackgroundColor, s.ForegroundColor, s.StringToColor)
		replaceStrings = append(replaceStrings, s.StringToColor, extraColored)
	}

	return strings.NewReplacer(replaceStrings...)
}
