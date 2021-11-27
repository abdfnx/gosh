# Shell

run powershell and bash with go.

## Run powershell and bash

```go
import "github.com/abdfnx/shell"

// run a command
shell.Run("git status")

// run a command with output
err, out, errout := shell.RunOut("whoami")

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

## Powershell

```go
import "github.com/abdfnx/shell"

// run a command
shell.PWSLCmd(`Write-Host "secman is 🔒"`)

// run a script
shell.PWSLCmd(`
  $x = git config user.name
  
  Write-Host $x
`)

// run a command with output
err, out, errout := shell.PWSLOut(`[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$SECMAN_PATH\bin", [System.EnvironmentVariableTarget]::User)`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

## Bash/Shell

```go
import "github.com/abdfnx/shell"

// run a command
shell.ShellCmd(`echo "secman is 🔒"`)

// run a script
shell.ShellCmd(`
  mood="👨‍💻"

  if [ $mood != "😪" ]; then
    echo "still coding"
  fi
`)

// run a command with output
err, out, errout := shell.ShellOut(`curl --silent "https://api.github.com/repos/scmn-dev/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```
