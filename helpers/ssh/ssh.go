package ssh

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/zerodayz/rcli/helpers/colors"
	"github.com/zerodayz/rcli/vars"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"sync"
)

type fn func(string, string, *ssh.ClientConfig) (string, string)

func Parallel(f fn, cmd string, hosts []string, config *ssh.ClientConfig) {
	var wg sync.WaitGroup
	for _, host := range hosts {
		wg.Add(1)
		go func(host string) {
			out, err := f(cmd, host, config)
			fmt.Printf(out)
			fmt.Printf(err)
			wg.Done()
		}(host)
	}
	wg.Wait()
}

func RunScriptCommand(scriptPath, host string, config *ssh.ClientConfig) (string, string) {
	var wg2 sync.WaitGroup

	if vars.Debug == true {
		log.Printf("DEBUG: connecting to: %s\n", host)
	}
	sshConnection, err := ssh.Dial("tcp", host, config)
	if err != nil {
		stdOutput := bytes.NewBuffer(nil)
		stdOutput.Write([]byte(colors.Cyan + " --- " + host + " ---" + colors.Reset +"\n"))
		stdOutput.Write([]byte(colors.Green + " Output:" + colors.Reset + "\n"))

		stdError := bytes.NewBuffer(nil)
		stdError.Write([]byte(colors.Red + " Error:" + colors.Reset + "\n"))
		stdError.Write([]byte(err.Error() + "\n"))
		return stdOutput.String(), stdError.String()
	}
	if vars.Debug == true {
		log.Printf("DEBUG: connected to: %s\n", host)
	}
	var buffer bytes.Buffer
	file, err := os.Open(scriptPath)
	if err != nil {
		log.Fatal(err)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: copying %s to: %s\n", scriptPath, host)
	}
	_, err = io.Copy(&buffer, file)
	if err != nil {
		log.Fatal(err)
	}

	scriptContent := &buffer
	if vars.Debug == true {
		log.Printf("DEBUG: content of the script %s:\n", scriptPath)
		fmt.Println(scriptContent)
	}
	sshSession, err := sshConnection.NewSession()
	if err != nil {
		panic(err)
	}
	defer sshSession.Close()
	if vars.Debug == true {
		log.Printf("DEBUG: executing [%s] on: %s\n", scriptPath, host)
	}
	sshSession.Stdin = scriptContent

	stdOutput := bytes.NewBuffer(nil)
	stdOutput.Write([]byte(colors.Cyan + " --- " + host + " ---" + colors.Reset +"\n"))
	stdOutput.Write([]byte(colors.Green + " Output:" + colors.Reset + "\n"))

	stdError := bytes.NewBuffer(nil)
	stdError.Write([]byte(colors.Red + " Error:" + colors.Reset + "\n"))

	wg2.Add(1)
	sessionStdOut, err := sshSession.StdoutPipe()
	go func() {
		defer wg2.Done()
		io.Copy(stdOutput, sessionStdOut)
	}()

	wg2.Add(1)
	sessionStderr, err := sshSession.StderrPipe()
	go func() {
		defer wg2.Done()
		io.Copy(stdError, sessionStderr)
	}()


	sshSession.Shell()
	sshSession.Wait()
	wg2.Wait()
	sshConnection.Close()

	return stdOutput.String(), stdError.String()
}

func RunCommand(cmd, host string, config *ssh.ClientConfig) (string, string) {
	var wg2 sync.WaitGroup

	if vars.Debug == true {
		log.Printf("DEBUG: connecting to: %s\n", host)
	}
	sshConnection, err := ssh.Dial("tcp", host, config)
	if err != nil {
		stdOutput := bytes.NewBuffer(nil)
		stdOutput.Write([]byte(colors.Cyan + " --- " + host + " ---" + colors.Reset +"\n"))
		stdOutput.Write([]byte(colors.Green + " Output:" + colors.Reset + "\n"))

		stdError := bytes.NewBuffer(nil)
		stdError.Write([]byte(colors.Red + " Error:" + colors.Reset + "\n"))
		stdError.Write([]byte(err.Error() + "\n"))
		return stdOutput.String(), stdError.String()
	}
	if vars.Debug == true {
		log.Printf("DEBUG: connected to: %s\n", host)
	}

	sshSession, err := sshConnection.NewSession()
	if err != nil {
		panic(err)
	}
	defer sshSession.Close()

	if vars.Debug == true {
		log.Printf("DEBUG: executing [%s] on: %s\n", cmd, host)
	}

	stdOutput := bytes.NewBuffer(nil)
	stdOutput.Write([]byte(colors.Cyan + " --- " + host + " ---" + colors.Reset +"\n"))
	stdOutput.Write([]byte(colors.Green + " Output:" + colors.Reset + "\n"))

	stdError := bytes.NewBuffer(nil)
	stdError.Write([]byte(colors.Red + " Error:" + colors.Reset + "\n"))

	wg2.Add(1)
	sessionStdOut, err := sshSession.StdoutPipe()
	go func() {
		defer wg2.Done()
		io.Copy(stdOutput, sessionStdOut)
	}()

	wg2.Add(1)
	sessionStderr, err := sshSession.StderrPipe()
	go func() {
		defer wg2.Done()
		io.Copy(stdError, sessionStderr)
	}()

	sshSession.Run(cmd)
	wg2.Wait()
	sshConnection.Close()

	return stdOutput.String(), stdError.String()
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hosts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hosts = append(hosts, scanner.Text())
	}
	return hosts, scanner.Err()
}