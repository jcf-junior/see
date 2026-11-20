package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"see/internals/models"
)

func PrintVersion() {
	fmt.Println("see version:", models.Version)
	
	latestVersion, err := CheckNewVersion()

	if err != nil {
		log.Fatal("Version check:\n", err)
		return
	}

	if (latestVersion == "v"+models.Version) {
		fmt.Println("You're running the latest version!")
	} else {
		fmt.Println("Your version is out of date. The latest version is: ", latestVersion)
	}
}

func CheckNewVersion() (string, error) {
	res, err := http.Get("https://api.github.com/repos/jcf-junior/see/releases/latest")
	if err != nil {
		return "", fmt.Errorf("Version check failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read response from GitHub API: %w", err)
	}

	var tagName struct {
		Tag string `json:"tag_name"`
	} 

	err = json.Unmarshal(body, &tagName)
	if err != nil {
		return "", fmt.Errorf("Failed to convert the JSON data received from the GitHub API: %w", err)
	}

	return tagName.Tag, nil
}

func PrintHelp() {
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


func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}