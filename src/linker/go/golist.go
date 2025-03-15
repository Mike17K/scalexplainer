package linker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Mike17K/scalexplainer/src/utils"
)

func RunGoList(path string) error {
	finalPath, err := utils.GetPath(path)
	if err != nil {
		return err
	}
	fmt.Println("Root: ", finalPath)

	goListCmd := exec.Command("go", "list", "-json", filepath.Join(finalPath, "..."))
	goListCmdStdout, err := goListCmd.StdoutPipe()
	if err != nil {
		return err
	}
	goListCmd.Stderr = os.Stderr

	// , "|", "jq . > go-deps.json")
	jqCmd := exec.Command("jq", ".")
	jqCmd.Stdin = goListCmdStdout
	jqCmd.Stdout = os.Stdout
	jqCmd.Stderr = os.Stderr

	if err := jqCmd.Start(); err != nil {
		return err
	}

	if err := goListCmd.Start(); err != nil {
		return err
	}

	// wait for both commands to complete
	if err := goListCmd.Wait(); err != nil {
		return err
	}
	if err := jqCmd.Wait(); err != nil {
		return err
	}

	return nil
}
