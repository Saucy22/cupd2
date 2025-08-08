// Saucy22 on GitHub 2025
package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Variables declared at runtime, im new to golang so im not sure if it is best practice to declare variables outside of main... oh well!
var version string = "2.0 Golang Beta"
var isUbuntu bool = false
var isFedora bool = false
var isSuse bool = false
var isArch bool = false

// A little abstraction. I could have added this to the main function but i wanted to make it more readable. You're welcome.
func windowsCheck() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}

// You have no idea how long it took me to come up with this solution for executing system commands. Its not perfect but it works for this program
// The issue with this compared to python's os.system implementation is how convoluted it is to get continuous output from a command. It got so confusing that
// I gave up on doing that and instead made this.
func executeCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", output)
}

func main() {
	if windowsCheck() == true {
		time.Sleep(time.Second)
		panic("Windows Detected, Abort.")
	}
	executeCommand("clear")
	fmt.Println("Welcome to Calebs Updater Version: " + version)
	time.Sleep(time.Second)
	//Entering the main loop
	for {
		fmt.Println("Which distro are you using right now?")
		//Notes: AFAIK Dnf5 is now default on fedora.
		fmt.Println("Debian/Ubuntu/Mint (1) | Fedora (Dnf5) (2) | Opensuse (3) | Arch/Endeavour OS (yay is a dependency) (4)")
		var choice string
		fmt.Scan(&choice)
		if choice == "1" {
			isUbuntu = true
			break
		} else if choice == "2" {
			isFedora = true
			break
		} else if choice == "3" {
			isSuse = true
			break
		} else if choice == "4" {
			isArch = true
			break
		} else {
			fmt.Println("Invalid choice, Please try again.")
			time.Sleep(time.Second * 3)
			executeCommand("clear")
		}
	}
	fmt.Print("Working... this may take a sec.")
	//System Updater
	if isUbuntu == true {
		executeCommand("sudo apt update && sudo apt upgrade -y")
	} else if isFedora == true {
		executeCommand("sudo dnf5 update -y")
	} else if isSuse == true {
		executeCommand("sudo zypper update && sudo zypper upgrade -y")
	} else if isArch == true {
		executeCommand("yay -y")
	}

	//Flatpak updater block, Yes is default because most people have flatpak (pre)installed on their system
	fmt.Println("Is flatpak installed on your system? [Y/n]")
	var choice2 string
	fmt.Scan(&choice2)
	if strings.ToLower(choice2) != "n" {
		executeCommand("flatpak update -y")
	}

	//Snap updater block, No is default because most people either dont use snaps or dont have it installed by default
	fmt.Println("Is snap/snapd installed on your system? [y/N]")
	var choice3 string
	fmt.Scan(&choice3)
	if strings.ToLower(choice3) != "y" {
		executeCommand("snap refresh -y")
	} //Line 100!!!
	fmt.Println("Updates (should be) finished! Thank you for using Caleb's Updater!")
	//The end... until i add more crap to this program
}
