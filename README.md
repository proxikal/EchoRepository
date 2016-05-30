# Gotools
Just a personal module I use for most of my DiscordGo bots.
  
### Main Purpose to Gotools
Basically i created this module to localized everything into one system.  
  
More documentation soon.

### gto.Replace()

```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	t := gto.Replace(text, "Hello", "Goodbye", -1)
	gto.Print(t)
}
```
  

### gto.Contains()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	if gto.Contains(text, "User") {
		gto.Print("The string contains the word: User")
	}
}
```
  
  
### gto.Split()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Welcome User!"
	s := gto.Split(text, " ")
	beginning := s[0] // Welcome
	after := s[1] // User!
}
```
  

### gto.ToLower()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	t := gto.ToLower(text)
	gto.Print(t) // prints hello user!
}
```
  

### gto.ToUpper()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	t := gto.toUpper(text)
	gto.Print(t) // prints HELLO USER!
}
```
  

### gto.TrimPrefix()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	t := gto.TrimPrefix(text, "Hello ")
	gto.Print(t) // prints User!
}
```
  

### gto.TrimSuffix()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	t := gto.TrimSuffix(text, "!")
	gto.Print(t) // displays Hello User
}
```
  
  
### gto.HasPrefix()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	if gto.HasPrefix(text, "Hello") {
		gto.Print(t) // string has the prefix Hello
	} else {
		gto.Print("The string doesn't have the prefix Hello")
	}
}
```
  

### gto.HasSuffix()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	if gto.HasSuffix(text, "!") {
		gto.Print(t) // string has the suffix !
	} else {
		gto.Print("The string doesn't have the suffix !")
	}
}
```
  

### gto.String()
```go
package main
import (
	"gotools"
)

func main() {
	mint := 32
	io, err := gto.String(mint)
	if err == nil {
		gto.Print("Converted int to string: "+io)
	}
}
```
  

### gto.Integer()
```go
package main
import (
	"gotools"
)

func main() {
	text := "32"
	io, err := gto.Integer(mint)
	if err == nil {
		integer := io // displays 32 as integer
	}
}
```
  

### gto.Print()
```go
package main
import (
	"gotools"
)

func main() {
	text := "Hello User!"
	gto.Print(t) // displays Hello User!
}
```
  

## gto.OpenURL()
```go
package main
import (
	"gotools"
)

func main() {
	// opens your default browser.
	gto.OpenURL("http://www.youtube.com")
}
```
  

### gto.DownloadFile()
```go
package main
import (
	"gotools"
)

func main() {
	if gto.DownloadFile("path\to\file.txt", "http://link.com/to/file.txt") {
		// the download was complete without errors.
	} else {
		// the download had an error
		gto.Print(err)
	}
}
```
  

### gto.getJson()
```go
package main
import (
	"gotools"
)

func main() {
	var jso map[string]interface{}
	getJson("http://url.com/to.json", &jso)
	if _, ok := jso["Username"]; ok {
		// the json entry "Username" exists.
		user := jso["Username"].(string)
		gto.Print(user)
	}
}
```
  

### gto.Unmarshal()
```go
package main
import (
	"gotools"
)

func main() {
	var json map[string]interface{}
	file, err := gto.ReadFile("path.json")
	if err == nil {
		gto.Unmarshal(file, &json)
		user := json["Username"].(string)
		gto.Print(user)
	}
}
```
  

### gto.Marshal()
INFO: `gto.Marshal()` is the same as `json.MarshalIndent()`  
  
```go
package main
import (
	"gotools"
)

func main() {
	var json map[string]interface{}
	file, err := gto.ReadFile("path.json")
	if err == nil {
		gto.Unmarshal(file, &json)
		json["Username"] = "Proxy"
		jf, err := gto.Marshal(json)
		if err == nil {
			gto.WriteFile("path.json", jf, 0777)
		}
	}
}
```
  

### gto.ReadFile()
```go
package main
import (
	"gotools"
)

func main() {
	file, err := gto.ReadFile("path.to.file.txt")
	if err == nil {
		// do work here
	}
}
```
  

### gto.WriteFile()
```go
package main
import (
	"gotools"
)

func main() {
	var jf []byte
	gto.WriteFile("path.to.file", jf, 0777)
}
```
  

### gto.Random()
```go
package main
import (
	"gotools"
)

func main() {
	t := gto.Random(1, 52)
	gto.Print(t) // displays random integer from 1 to 52
}
```
  
  
### gto.readLines()
```go
package main
import (
	"gotools"
)

func main() {
	lines, err := gto.readLines("path.to.txt")
	if err == nil {
		gto.Print(lines[0]) // prints first line of a file.
	}
}
```
  

### gto.countLines()
```go
package main
import (
	"gotools"
)

func main() {
	count := gto.countLines("path.to.file.txt")
	io, err := gto.String(count)
	if err == nil {
		gto.Print("There are " + io + " lines in this file.")
	}
}
```
  

### gto.CopyFileContents()
Soon!
  

## gto.CopyFile()
Soon!
  

## gto.SendMessage()
Using `gto.SendMessage(s *discordgo.Session, ChannelID, "Message Here")`
  

