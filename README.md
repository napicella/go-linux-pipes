Companion code for the article Using Linux Pipes with Go

### Build
```bash
go build -o ./bin/uppercase
```

### Example
```bash
> cat words | sort | uniq | uppercase
APPLE
BYE
HELLO
ZEBRA
```

### Command line options
```bash
> ./bin/uppercase --help
Simple demo of the usage of linux pipes
Transform the input (pipe or file) to uppercase letters

Usage:
  uppercase [flags]

Flags:
  -f, --file string   path to the file
  -h, --help          help for uppercase
  -v, --verbose       log verbose output
  
  
> echo "hello from the shell" | ./bin/uppercase
  HELLO FROM THE SHELL
  
> ./bin/uppercase -f /tmp/test 
  HELLO FROM FILE    
```