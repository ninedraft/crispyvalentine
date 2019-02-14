package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	var help = flag.Bool("credits", false, "print credits info")
	flag.Parse()
	if *help {
		fmt.Printf("credits:\n\t+ %s", strings.Join(credits(), "\n\t+ "))
		return
	}
	var world = NewWorld(heart)
	for range time.Tick(50 * time.Millisecond) {
		tm.Clear()
		var text = string(world.state)
		tm.MoveCursor(1, 1)
		tm.Println(text)
		tm.Flush()
		world.Step()
	}
}

func printErrAndExit(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(100)
}

func credits() []string {
	return []string{
		"Normand Veilleux",
		"Dustin Slater (zombie@camelot.bradley.edu)",
	}
}
