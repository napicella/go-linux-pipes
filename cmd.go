package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

func runCommand() error {
	if isInputFromPipe() {
		print("data is from pipe")
		return toUppercase(os.Stdin, os.Stdout)
	} else {
		file, e := getFile()
		if e != nil {
			return e
		}
		defer file.Close()
		return toUppercase(file, os.Stdout)
	}
}

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode() & os.ModeCharDevice == 0
}

func getFile() (*os.File, error){
	if flags.filepath == "" {
		return nil, errors.New("please input a file")
	}
	if !fileExists(flags.filepath) {
		return nil, errors.New("the file provided does not exist")
	}
	file, e := os.Open(flags.filepath)
	if e != nil {
		return nil, errors.Wrapf(e,
			"unable to read the file %s", flags.filepath)
	}
	return file, nil
}

func toUppercase(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		_, e := fmt.Fprintln(
			w, strings.ToUpper(scanner.Text()))
		if e != nil {
			return e
		}
	}
	return nil
}

func fileExists(filepath string) bool {
	info, e := os.Stat(filepath)
	if os.IsNotExist(e) {
		return false
	}
	return !info.IsDir()
}
