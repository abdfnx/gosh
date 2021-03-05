package shell

import (
	"fmt"
	"log"
	"runtime"
	"bytes"
	"os/exec"
)

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}


func SHCore(cmd string, winCmd string) {
	err, out, errout := Shellout("")
	
	if runtime.GOOS == "windows" {
		err, out, errout = Shellout(winCmd)
	} else {
		err, out, errout = Shellout(cmd)
	}
	
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Println(errout)
	}
	
	fmt.Println(out)
}
