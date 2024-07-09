package shell

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const shellToUse = "bash"

func ExecBashCommand(command string) error {
	cmd := exec.Command(shellToUse, "-c", command)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println(stdBuffer.String())
	return nil
}
