package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("hello, world!")
	cmd := exec.Command("docker", "pull", "hello-world")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("succes")
}
