package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func execCommand(args ...string) string {
	cmd := exec.Command(args[0], args[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("%s: %s", err, out.String()))
	}
	return out.String()
}

func execCommands(commands [][]string) {
	for _, cmd := range commands {
		execCommand(cmd...)
	}
}

func copyFile(src, dst string) {
	from, err := os.OpenFile(src, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		panic(err)
	}
}

func options(opts map[string]string) string {
	ret := []string{}
	for k, v := range opts {
		ret = append(ret, k+"="+v)
	}
	return strings.Join(ret, ",")
}
