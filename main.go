package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Getenv, os.Stdin, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, getEnv func(string) string, r io.Reader, w io.Writer, args []string) error {
	fmt.Println("Wicho: Starting")
	// Input that splits on new lines
	input := bufio.NewScanner(r)
	// Buffered output
	output := bufio.NewWriter(w)

	// TODO: Wicho: testing adding text
	lineNo := 0

	for input.Scan() {
		text := input.Text()
		lineNo++
		s := fmt.Sprintf("%03d %s\n", lineNo, text)

		// NOTE: Wicho: how should we handle func output?
		output.WriteString(s)
	}

	// Flushed remaining buffered output
	output.Flush()

	return nil
}
