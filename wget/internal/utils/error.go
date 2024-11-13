package utils

import (
	"fmt"
	"os"
)

func PrintError() {
	fmt.Fprintln(os.Stderr, "wget: missing URL")
	fmt.Fprintln(os.Stderr, "Usage: wget [OPTION]... [URL]...\n")
	fmt.Fprintln(os.Stderr, "Try `wget --help' for more options.")
}

func OptionsError(option string) {
	fmt.Fprintf(os.Stderr, "wget: invalid option -- '%s'\n", option)
	fmt.Fprintln(os.Stderr, "Usage: wget [OPTION]... [URL]...\n")
	fmt.Fprintln(os.Stderr, "Try `wget --help' for more options.")
}
