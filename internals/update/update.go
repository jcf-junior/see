package update

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"see/internals/utils"
	"strings"
)

func Update(currentVersion string) {

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

	var input string

	for {
		fmt.Println("The installation will run 'sudo apt install'. If you don't want to update as sudo, please refer to the manual installation steps in GitHub.")
		fmt.Print("Proceed with installation? (Y/n) ")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "y" {
			break
		} else if strings.ToLower(input) == "n" {
			fmt.Println("Aborting update...")
			return
		}
		continue
	}

	url := fmt.Sprintf("https://github.com/jcf-junior/see/releases/latest/download/see_%s.deb", runtime.GOARCH)

	curl := exec.Command("curl", "-fSL", "--create-dirs", "-o", "/tmp/see/see.deb", url)
	setInstallPathPerms := exec.Command("chmod", "-R", "755", "/tmp/see/")
	install := exec.Command("sudo", "apt", "install", "/tmp/see/see.deb")
	removeUpdateFiles := exec.Command("rm", "/tmp/see/see.deb")

	curl.Stdin = os.Stdin
	curl.Stdout = os.Stdout
	curl.Stderr = os.Stderr

	install.Stdin = os.Stdin
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr

	setInstallPathPerms.Stdin = os.Stdin
	setInstallPathPerms.Stdout = os.Stdout
	setInstallPathPerms.Stderr = os.Stderr

	install.Stdin = os.Stdin
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr

	removeUpdateFiles.Stdin = os.Stdin
	removeUpdateFiles.Stdout = os.Stdout
	removeUpdateFiles.Stderr = os.Stderr

	fmt.Println("Downloading files..")
	err = curl.Run()
	if err != nil {
		log.Fatal("Failed to download update files: ", err)
	} else {
		fmt.Println("Files downloaded to /tmp/see/see.deb")
	}

	fmt.Println("Setting /tmp/see permissions to 755...")
	err = setInstallPathPerms.Run()
	if err != nil {
		log.Fatal("Failed to set permissions: ", err)
	}
	
	fmt.Println("Installing files..")
	err = install.Run()
	if err != nil {
		log.Fatal("Failed to install downloaded files: ", err)
	}

	fmt.Println("Removing temporary files..")
	err = removeUpdateFiles.Run()
	if err != nil {
		log.Fatal("Failed to remove downloaded files: ", err)
	} else {
		fmt.Println("Removed successfully")
	}
}