package aliases

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

type BinaryManager struct {
	binaryDir string
}

func (manager *BinaryManager) getFilename(image string, tag string) string {
	filename := fmt.Sprintf("%s:%s", image, tag)
	filename = strings.Replace(filename, "/", "-", -1)
	filename = strings.Replace(filename, ":", "-", -1)
	filename = strings.Replace(filename, ".", "-", -1)
	filename = strings.Replace(filename, "_", "-", -1)

	return filename
}

func (manager *BinaryManager) exists(image string, tag string) bool {
	filename := manager.getFilename(image, tag)
	filepath := path.Join(manager.binaryDir, filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}

	return true
}

func (manager *BinaryManager) download(image string, tag string) error {
	filename := manager.getFilename(image, tag)

	cmd := exec.Command(
		"docker",
		"run",
		"-v",
		fmt.Sprintf("%s:%s", manager.binaryDir, "/data"),
		fmt.Sprintf("%s:%s", image, tag),
		"sh",
		"-c",
		fmt.Sprintf("cp $(which docker) /data/%s", filename),
	)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (manager *BinaryManager) Get(image string, tag string) (*string, error) {
	if !manager.exists(image, tag) {
		if err := manager.download(image, tag); err != nil {
			return nil, fmt.Errorf("runtime error: %s", err)
		}
	}

	filepath := path.Join(manager.binaryDir, manager.getFilename(image, tag))

	return &filepath, nil
}
