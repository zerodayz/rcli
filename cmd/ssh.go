package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zerodayz/rcli/helpers/ssh"
	"github.com/zerodayz/rcli/vars"
	"log"
	"os"
	"time"
)

var (
	hosts 		[]string
	username 	string
	command 	string
	filename 	string
	hostsFile	string
	debug		bool
)

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.PersistentFlags().StringSliceVarP(&hosts, "hosts", "H", []string{""}, "hosts" + "\nFor example 192.168.1.10:22,192.168.1.12:22")
	sshCmd.PersistentFlags().StringVarP(&username, "username", "U", "", "username")
	sshCmd.PersistentFlags().StringVar(&hostsFile, "hosts-file", "", "hosts file")
	sshCmd.PersistentFlags().BoolVarP(&vars.Debug, "debug", "d", false, "enable debug")

	sshCmd.MarkFlagRequired("username")

	sshCmd.AddCommand(runSshCmd)
	runSshCmd.Flags().StringVarP(&command, "containerCommand", "c", "", "containerCommand")
	runSshCmd.MarkFlagRequired("containerCommand")

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
			hosts, err = ssh.ReadLines(hostsFile);
			if err != nil {
				log.Printf("ERROR: %s", err)
				os.Exit(1)
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

var runscriptSshCmd = &cobra.Command{
	Use:   "runscript",
	Short: "execute script over ssh",
	Long:  `execute script over ssh`,
	Run: func(cmd *cobra.Command, args []string) {
		ShowLogo()
		start := time.Now()
		if hostsFile != "" {
			var err error
			hosts, err = ssh.ReadLines(hostsFile);
			if err != nil {
				log.Printf("ERROR: %s", err)
				os.Exit(1)
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