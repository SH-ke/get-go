package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
	// A place to store the first error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	// Skips the write if an error occurred previously
	_, sw.err = fmt.Fprintln(sw.w, s)
	// Writes a line and store any error
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors,handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface,the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite,yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don't communicate by sharing memory,share memory by communicating.")
	sw.writeln("Channels orchestrate;mutexes serialize.")
	return sw.err
	// Returns an error, if one occurred
}
func main() {
	err := proverbs("safe-writer.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
