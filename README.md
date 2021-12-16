# Shell

run powershell and bash easly with go.

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
shell.PowershellCommand(`Write-Host "hello from powershell"`)

// run a script
shell.PowershellCommand(`
  $git_username = git config user.name
  
  Write-Host $git_username
`)

// run a command with output
err, out, errout := shell.PowerShelloutput(`[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$APP_PATH\bin", [System.EnvironmentVariableTarget]::User)`)

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
shell.ShellCommand(`echo "shell or bash?"`)

// run a script
shell.ShellCommand(`
  mood="👨‍💻"

  if [ $mood != "😪" ]; then
    echo "still coding"
  fi
`)

// run a command with output
err, out, errout := shell.Shelloutput(`curl --silent "https://api.github.com/repos/abdfnx/resto/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```
