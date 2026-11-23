package update

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"see/internals/utils"
)

func UpdateSystemWide(currentVersion string) {

	latestVersion, err := utils.CheckNewVersion()
	if err != nil {
		log.Fatal("Failed to fetch latest version: ", err)
	}

	if latestVersion == currentVersion {
		fmt.Println("see is running the latest version!")
		return
	}

	fmt.Println("Latest Version: ", latestVersion)
	fmt.Println("Current Version: ", currentVersion)

	// TODO: Change .deb to include version
	url := fmt.Sprintf("https://github.com/jcf-junior/see/releases/latest/download/see_%s.deb", runtime.GOARCH)

	updatePath := "usr/bin/see"
	updateFilePath := fmt.Sprintf("%s/see_%s.update.deb", updatePath, runtime.GOARCH)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to get latest release")
	}
	defer res.Body.Close()
	
	updateFiles, err := os.Create(updateFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer updateFiles.Close()
	
	fmt.Println("Downloading files to", updateFilePath)
	_, err = io.Copy(updateFiles, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	install := exec.Command("apt", "install", updateFilePath)
	install.Stdin = os.Stdin
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr

	fmt.Println("Installing files..")
	err = install.Run()
	if err != nil {
		log.Fatal("Failed to install downloaded files: ", err)
	}

	fmt.Println("Removing temporary files...")
	err = os.Remove(updateFilePath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Update completed successfully")
	}
}
