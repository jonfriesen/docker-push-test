package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("running pull")
	cmd := exec.Command("docker", "pull", "hello-world")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("running run")
	cmd2 := exec.Command("docker", "run", "hello-world")
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	if err := cmd2.Run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("succes")
}
