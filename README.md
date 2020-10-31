# tidy-file
File utils library for Golang projects.

## Usage

### Command
#### command.Executor.Execute()
Execute io commands based on the destination and source path.
```go
package main

import (
    "github.com/alancesar/tidy-file/command"
    "os"
)

func main() {
    sourcePath := "source/path"
    destinationPath := "destination/path"
    
    command.
        NewExecutor(sourcePath, destinationPath).
        Execute(command.MkDirCommand, os.Rename)
}
```

### MIME Type
#### mime.Is()
Check MIME types from files.
```go
package main

import (
    "fmt"
    "github.com/alancesar/tidy-file/mime"
)

func main() {
    path := "source/path/audio.mp3"
    if mime.Is(path, mime.AudioType) {
        fmt.Printf("%s is an audio file", path)
    }
}
```

### Path
#### path.LookFor()
Look for a specific file type in a directory and its sub directories.
```go
package main

import (
    "fmt"
    "github.com/alancesar/tidy-file/mime"
    "github.com/alancesar/tidy-file/path"
)

func main() {
    root := "source/root"
    images := path.LookFor(root, mime.ImageType)

    for _, image := range images {
        fmt.Printf("%s is an image file", image)
    }
}
```

#### path.Builder.FromPattern
creates a path from an interface{} based on a pattern using text/template engine.
```go
package main

import (
    "fmt"
    "github.com/alancesar/tidy-file/path"
)

func main() {
    myStruct := struct {
		Track  int
		Title  string
		Artist string
		Album  string
		Year   int
	}{
		Track:  7,
		Title:  "Basket Case",
		Artist: "Green Day",
		Album:  "Dookie",
		Year:   1994,
	}

    output, _ := path.NewBuilder().
    		FromPattern("{{.Artist}}/[{{.Year}}] {{.Album}}/{{printf \"%02d\" .Track}} - {{.Title}}", myStruct)

    // "Green Day/[1994] Dookie/07 - Basket Case"
    fmt.Printf(output) 
}
```