# gosh

> Run powershell and bash commands easly in go.

## Install

```bash
go get -v github.com/abdfnx/gosh@v0.3.5
```

## Examples

### Run one command on both shell and powershell

```go
package main

import (
   "fmt"
   "log"
   
   "github.com/abdfnx/gosh"
)

// run a command
gosh.Run("git status")

// run a command with output
err, out, errout := gosh.RunOut("whoami")

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

### How `gosh.Run("COMMAND")` works ?

```go
// `Run` executes the same command for shell and powershell
func Run(cmd string) {
	err, out, errout := ShellOutput("")

	if runtime.GOOS == "windows" {
		err, out, errout = PowershellOutput(cmd)
	} else {
		err, out, errout = ShellOutput(cmd)
	}

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
	
	fmt.Print(out)
}
```

### Powershell

```go
package main

import (
   "fmt"
   "log"
   
   "github.com/abdfnx/gosh"
)

// run a command
gosh.PowershellCommand(`Write-Host "hello from powershell"`)

// run a script
gosh.PowershellCommand(`
  $git_username = git config user.name
  
  Write-Host $git_username
`)

// run a command with output
err, out, errout := gosh.PowerShelloutput(`[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$APP_PATH\bin", [System.EnvironmentVariableTarget]::User)`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

### Bash/Shell

```go
package main

import (
   "fmt"
   "log"
   
   "github.com/abdfnx/gosh"
)

// run a command
gosh.ShellCommand(`echo "shell or bash?"`)

// run a script
gosh.ShellCommand(`
  mood="👨‍💻"

  if [ $mood != "😪" ]; then
    echo "still coding"
  fi
`)

// run a command with output
err, out, errout := gosh.Shelloutput(`curl --silent "https://api.github.com/repos/abdfnx/resto/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```
