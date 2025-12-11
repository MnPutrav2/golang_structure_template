package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func Build() {
	os.MkdirAll("build", 0755)

	cmd := exec.Command("go", "build", "-o", "./build/server", "./cmd")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Build error:", err)
		return
	}

	fmt.Println("Build success!")
}
