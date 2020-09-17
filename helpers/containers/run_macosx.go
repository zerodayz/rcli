// +build darwin

package containers

import (
	"bytes"
	"fmt"
	"github.com/zerodayz/rcli/vars"
	"log"
	"os"
	"os/exec"
	"strings"
)

// #import <mach-o/dyld.h>
import "C"

func NSGetExecutablePath() string {
	var buflen C.uint32_t = 1024
	buf := make([]C.char, buflen)

	ret := C._NSGetExecutablePath(&buf[0], &buflen)
	if ret == -1 {
		buf = make([]C.char, buflen)
		C._NSGetExecutablePath(&buf[0], &buflen)
	}
	return C.GoStringN(&buf[0], C.int(buflen))
}

func ChildRcli(command, image string) {
	var cmd *exec.Cmd

	commandArgs := strings.Split(command, " ")
	if vars.Debug == true {
		log.Printf("DEBUG: Executing command %v in container.\n", commandArgs)
	}
	if len(commandArgs) == 1 {
		cmd = exec.Command(commandArgs[0])
	} else {
		cmd = exec.Command(commandArgs[0], commandArgs[1:]...)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func RunRcli(command, image string) {
	var cmd *exec.Cmd

	Command := string(bytes.Trim([]byte(NSGetExecutablePath()), "\x00"))
	if vars.Debug == true {
		cmd = exec.Command(Command, append([]string{"container", "-d", "run", "fork", "-c"}, command)...)
	} else {
		cmd = exec.Command(Command, append([]string{"container", "run", "fork", "-c"}, command)...)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	}
}
