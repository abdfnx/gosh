# gosh

> A golang library for executing bash & powershell commands easly.

## Install

```bash
go get -v github.com/abdfnx/gosh
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
err, out, errout := gosh.RunOutput("echo 𝜋")

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

### Run Powershell Command(s)

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
err, out, errout := gosh.PowershellOutput(`[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$APP_PATH\bin", [System.EnvironmentVariableTarget]::User)`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

### Run Bash/Shell Command(s)

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
err, out, errout := gosh.ShellOutput(`curl --silent "https://get-latest.onrender.com/docker/compose"`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```
