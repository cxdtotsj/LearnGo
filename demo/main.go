package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	t := flag.String("f", "", "f")
	flag.Parse()

	if *t != "" {
		switch *t {
		case "a":
			os.Stderr.WriteString("aaaaa")
			return
		default:
			os.Stdout.WriteString("bbbbb")
			os.Exit(1)
		}
	}

	cmd := exec.Command("./test", "-f", "b")

	outBuf := &bytes.Buffer{}
	cmd.Stdout = outBuf
	errBuf := &bytes.Buffer{}
	cmd.Stderr = errBuf

	err := cmd.Run()
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("out:", outBuf.String(), "stderr:", errBuf.String())
}
