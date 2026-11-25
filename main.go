package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"see/internals/models"
	"see/internals/update"
	"see/internals/utils"
	"strings"
)

var version = "dev"
var formattedCurrentVersion = "v" + version

func main() {
	flag.Usage = utils.PrintHelp
	showVersion := flag.Bool("v", false, "display installed version")
	useCat := flag.Bool("c", false, "use cat, even if bat is installed")
	updateVersion := flag.Bool("u", false, "update see to the latest version")

	flag.Parse()

	if *updateVersion {
		update.UpdateSystemWide(formattedCurrentVersion)
	}

	if *showVersion {
		utils.PrintVersion(formattedCurrentVersion)
		return
	}

	var target models.Target

	if len(os.Args) == 1 {
		target.Path = "."
	} else {
		for _, arg := range os.Args[1:] {
			if !strings.HasPrefix(arg, "-") {
				target.Path = arg
			}
		}
	}

	info, err := os.Stat(target.Path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: path does not exist or is invalid")
		os.Exit(1)
	}

	if info.IsDir() {
		target.Cmd = "ls"
	} else if info.Mode().IsRegular() {
		target.Cmd = "bat"

		if *useCat || !utils.CommandExists("bat") {
			target.Cmd = "cat"
		}
	}

	cmd := exec.Command(target.Cmd, target.Path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
