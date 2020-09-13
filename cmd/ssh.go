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
	hosts		[]string
	username 	string
	command 	string
	filename 	string
	hostsFile	string
	hostsCli 	[]string
	src			string
	dst			string
	debug		bool
)

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.PersistentFlags().StringSliceVarP(&hostsCli, "hosts", "H", []string{""}, "hosts" + "\nFor example 192.168.1.10:22,192.168.1.12:22")
	sshCmd.PersistentFlags().StringVarP(&username, "username", "U", "", "username")
	sshCmd.PersistentFlags().StringVar(&hostsFile, "hosts-file", "", "hosts file")
	sshCmd.PersistentFlags().BoolVarP(&vars.Debug, "debug", "d", false, "enable debug")

	sshCmd.MarkFlagRequired("username")

	sshCmd.AddCommand(runSshCmd)
	runSshCmd.Flags().StringVarP(&command, "containerCommand", "c", "", "containerCommand")
	runSshCmd.MarkFlagRequired("containerCommand")

	sshCmd.AddCommand(cpSshCmd)
	cpSshCmd.Flags().StringVarP(&src, "source", "", "", "source file or directory")
	cpSshCmd.Flags().StringVarP(&dst, "destination", "", "", "destination")
	cpSshCmd.MarkFlagRequired("source")
	cpSshCmd.MarkFlagRequired("destination")

	sshCmd.AddCommand(runscriptSshCmd)
	runscriptSshCmd.Flags().StringVarP(&filename, "filename", "f", "", "filename")
	runscriptSshCmd.MarkFlagRequired("filename")
}


var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh namespace",
	Long:  `ssh namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		end := time.Now()
		log.Println(end.Sub(start))
	},
}

var runSshCmd = &cobra.Command{
	Use:   "run",
	Short: "execute commands over ssh",
	Long:  `execute commands over ssh`,
	Run: func(cmd *cobra.Command, args []string) {
		ShowLogo()
		start := time.Now()
		if hostsFile != "" {
			var err error
			hosts, err = ssh.ReadLines(hostsFile)
			if err != nil {
				log.Printf("ERROR: %s", err)
				os.Exit(1)
			}
		}
		if len(hostsCli) > 1 {
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
		ssh.Parallel(ssh.RunCommand, command, c.Hosts, config)

		end := time.Now()
		log.Println(end.Sub(start))
	},
}

var cpSshCmd = &cobra.Command{
	Use:   "cp",
	Short: "copy file or directory over ssh",
	Long:  `copy file or directory over ssh`,
	Run: func(cmd *cobra.Command, args []string) {
		ShowLogo()
		start := time.Now()
		// TODO
		end := time.Now()
		log.Println(end.Sub(start))
	},
}

var runscriptSshCmd = &cobra.Command{
	Use:   "runscript",
	Short: "execute script over ssh",
	Long:  `execute script over ssh`,
	Run: func(cmd *cobra.Command, args []string) {
		ShowLogo()
		start := time.Now()
		if hostsFile != "" {
			var err error
			hosts, err = ssh.ReadLines(hostsFile)
			if err != nil {
				log.Printf("ERROR: %s", err)
				os.Exit(1)
			}
		}
		if len(hostsCli) > 1 {
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
		ssh.Parallel(ssh.RunScriptCommand, filename, c.Hosts, config)

		end := time.Now()
		log.Println(end.Sub(start))
	},
}