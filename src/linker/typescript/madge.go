package linker

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Mike17K/scalexplainer/src/utils"
)

func checkAndInstallMadge() error {
	fmt.Println("Checking if madge is installed...")
	cmd := exec.Command("npx", "--yes", "madge", "--version")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Installing madge...")
		cmd = exec.Command("npm", "install", "-g", "madge")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	fmt.Println("Madge installed successfully")
	return nil
}

func RunMadge(path string) error {
	finalPath, err := utils.GetPath(path)
	if err != nil {
		return err
	}
	fmt.Println("Root: ", finalPath)

	err = checkAndInstallMadge()
	if err != nil {
		fmt.Println("Error installing madge")
		return err
	}
	cmd := exec.Command("npx", "madge", "--json", finalPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
