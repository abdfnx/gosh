package shell

import (
	"fmt"
	"log"
	"runtime"
	"bytes"
	"os/exec"
)

func ShellOut(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func ShellCmd(command string) {
	err, out, errout = ShellOut(command)

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Println(errout)
	}
	
	fmt.Println(out)
}

func SHCore(cmd string, winCmd string) {
	err, out, errout := ShellOut("")
	
	if runtime.GOOS == "windows" {
		err, out, errout = ShellOut(winCmd)
	} else {
		err, out, errout = ShellOut(cmd)
	}
	
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Println(errout)
	}
	
	fmt.Println(out)
}
