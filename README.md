# Shell

run powershell and bash with go.

## Powershell

```go
import "github.com/secman-team/shell"

// run a command
shell.PWSLCmd(`Write-Host "secman is 🔒"`)

// run a script
shell.PWSLCmd(`
  $x = git config user.name
  
  Write-Host $x
`)

// run a command with out
err, out, errout := shell.PWSLOut(`[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$SECMAN_PATH\bin", [System.EnvironmentVariableTarget]::User)`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

## Bash/Shell

```go
import "github.com/secman-team/shell"

// run a command
shell.ShellCmd(`echo "secman is 🔒"`)

// run a script
shell.ShellCmd(`
  mood="👨‍💻"

  if [ $mood != "😪" ]; then
    echo "still coding"
  fi
`)

// run a command with out
err, out, errout := shell.ShellOut(`curl --silent "https://api.github.com/repos/secman-team/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```
