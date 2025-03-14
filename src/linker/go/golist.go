package linker

import (
	"os"
	"os/exec"
	"path/filepath"
)

func RunGoList(path string) error {
	cmd := exec.Command("go", "list", "-json", filepath.Join(path, "..."), "|", "jq . > go-deps.json")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
