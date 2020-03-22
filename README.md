
# CS50 in GO LANG

A GO implementation of **Harvard's cs50** library. Feel free to look around, fix bugs or make improvements.

## Installation
Make sure you have [GO](https://golang.org/) installed first. Than simply run the command in your command-line.

`go get github.com/muhammadidrees/cs50-golang`

## Usage
``` go
package main

import (
	"github.com/muhammadidrees/cs50-golang"
)

func main(){
  c := cs50.GetChar("Prompt: ")
  s := cs50.GetString("Prompt: ")
  i := cs50.GetInt("Prompt: ")
  f := cs50.GetFloat("Prompt: ")
} 
```
## Todo
- [x] Implement CS50 library in golang
- [ ] Write Tests
