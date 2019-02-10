package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func Download(binaryPath string, image string, tag string) error {
	if _, err := os.Stat(binaryPath); !os.IsNotExist(err) {
		return nil
	}

	cmd := exec.Command(
		"docker",
		"run",
		"-v",
		fmt.Sprintf("%s:%s", path.Dir(binaryPath), "/share"),
		fmt.Sprintf("%s:%s", image, tag),
		"sh",
		"-c",
		fmt.Sprintf("cp $(which docker) /share/%s", path.Base(binaryPath)),
	)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
