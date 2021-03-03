package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zerodayz/rcli/helpers/ssh"
	"github.com/zerodayz/rcli/vars"
	"log"
	"os"
	"strings"
	"time"
)

var (
	hosts     []string
	username  string
	command   string
	filename  string
	hostsFile string
	hostsCli  []string
)

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.PersistentFlags().StringSliceVarP(&hostsCli, "hosts", "H", []string{}, "hosts"+"\nFor example 192.168.1.10:22,192.168.1.12:22")
	sshCmd.PersistentFlags().StringVarP(&username, "username", "U", ssh.GetEnvironment("USER"), "username")
	sshCmd.PersistentFlags().StringVar(&hostsFile, "hosts-file", "", "hosts file")
	sshCmd.PersistentFlags().BoolVarP(&vars.Debug, "debug", "d", false, "enable debug")

	sshCmd.AddCommand(runSshCmd)
	runSshCmd.Flags().StringVarP(&command, "command", "c", "", "command")
	runSshCmd.Flags().StringVarP(&filename, "filename", "f", "", "filename")
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh namespace",
	Long:  `ssh namespace`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}
	},
}

var runSshCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute commands over ssh",
	Long:  `execute commands over ssh`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		if hostsFile != "" {
			var err error
			hosts, err = ssh.ReadLines(hostsFile)
			if err != nil {
				log.Printf("ERROR: %s", err)
				os.Exit(1)
			}
		}
		if len(hostsCli) > 0 {
			for _, host := range hostsCli {
				if !strings.Contains(host, ":") {
					host += fmt.Sprintf(":%d", vars.SSHDefaultPort)
				}
				hosts = append(hosts, host)
			}
		}
		c := &ssh.Connection{Username: username,
			Hosts: hosts}
		config := ssh.InitializeSshAgent(&c.Username)

		if command == "" && filename != "" {
			ssh.Parallel(ssh.RunScriptCommand, filename, c.Hosts, config)
		} else if command != "" && filename == "" {
			ssh.Parallel(ssh.RunCommand, command, c.Hosts, config)
		}
		end := time.Now()
		log.Println(end.Sub(start))
	},
}
