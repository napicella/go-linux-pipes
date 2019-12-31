Companion code for the article Using Linux Pipes with Go

### Build
```bash
go build -o ./bin/go-pipe
```

### Run
```bash
> ./bin/go-pipe --help
Simple demo of the usage of linux pipes
Transform the input (pipe of file) to uppercase letters

Usage:
  uppercase [flags]

Flags:
  -f, --file string   path to the file
  -h, --help          help for uppercase
  -v, --verbose       log verbose output
  
  
> echo "hello from the shell" | ./bin/go-pipe
  HELLO FROM THE SHELL
  
> ./bin/go-pipe -f /tmp/test 
  HELLO FROM FILE    
```