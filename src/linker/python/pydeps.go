package linker

import (
	"fmt"
	"os"
	"os/exec"
)

func checkAndInstallPydeps() error {
	fmt.Println("Checking if pydeps is installed...")
	cmd := exec.Command("pip", "show", "pydeps")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Installing pydeps...")
		cmd = exec.Command("pip", "install", "pydeps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	fmt.Println("pydeps installed successfully")
	return nil
}

func RunPydeps(path string) error {
	err := checkAndInstallPydeps()
	if err != nil {
		fmt.Println("Error installing pydeps")
		return err
	}
	cmd := exec.Command("python", "-m", "pydeps", "--json", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
