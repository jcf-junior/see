package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var version = "dev"

func main() {

	flag.Usage = printHelp
	// useCat := flag.Bool("c", false, "use cat even if bat is installed")
	showVersion := flag.Bool("v", false, "display installed version")
	flag.Parse()

	if *showVersion {
		printVersion()
		return
	}

	if len(os.Args) == 1 {
		printHelp()
		return
	}

	var target struct {
		path string
		cmd  string
	}

	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "-") {
			target.path = arg
		}
	}

	info, err := os.Stat(target.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: path does not exist or is invalid")
		os.Exit(1)
	}

	if info.IsDir() {
		target.cmd = "ls"
	} else if info.Mode().IsRegular() {
		target.cmd = "cat"
	}

	cmd := exec.Command(target.cmd, target.path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func printHelp() {
	const (
		bold      = "\033[1m"
		underline = "\033[4m"
		reset     = "\033[0m"
	)
	fmt.Println("See the contents of a file or folder. See uses ls and cat under the hood, so it's already supported.")
	fmt.Println(bold + underline + "Usage:" + reset + " " + "see <flags> <path/to/target>")
	fmt.Println("\n" + bold + underline + "Flags:" + reset)
	fmt.Println(bold + "-c" + reset + "\t" + "Use cat even if bat is installed")
	fmt.Println(bold + "-h" + reset + "\t" + "Print this help text")
	fmt.Println(bold + "-v" + reset + "\t" + "Print the current version")
}

func printVersion() {
	fmt.Println("see version", version)
}
