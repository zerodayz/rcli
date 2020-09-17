package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zerodayz/rcli/helpers/containers"
	"github.com/zerodayz/rcli/vars"
	"log"
	"os"
	"time"
)

var (
	containerCommand string
	image            string
)

func init() {
	rootCmd.AddCommand(ContainerCmd)
	ContainerCmd.PersistentFlags().BoolVarP(&vars.Debug, "debug", "d", false, "enable debug")
	ContainerCmd.PersistentFlags().StringVarP(&vars.ContainerRuntime, "runtime", "r", "rcli", "container runtime")

	ContainerCmd.AddCommand(runContainerCmd)
	runContainerCmd.Flags().StringVarP(&containerCommand, "command", "c", "", "command")
	runContainerCmd.Flags().StringVarP(&image, "image", "i", "", "rootfs directory")
	runContainerCmd.MarkFlagRequired("containerCommand")
	runContainerCmd.MarkFlagRequired("image")

	runContainerCmd.AddCommand(forkCmd)
	forkCmd.Flags().StringVarP(&containerCommand, "command", "c", "", "command")
	forkCmd.Flags().StringVarP(&image, "image", "i", "", "rootfs directory")
	forkCmd.MarkFlagRequired("containerCommand")
	forkCmd.MarkFlagRequired("image")
}

var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "create container",
	Long:  `create container`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		end := time.Now()
		log.Println(end.Sub(start))
	},
}

var runContainerCmd = &cobra.Command{
	Use:   "run",
	Short: "execute command in container",
	Long:  `execute command in container`,
	Run: func(cmd *cobra.Command, args []string) {
		ShowLogo()
		if vars.ContainerRuntime == "rcli" {
			if vars.Debug == true {
				log.Printf("DEBUG: executing container runtime: %s.\n", vars.ContainerRuntime)
			}
			containers.RunRcli(containerCommand, image)
		} else {
			if vars.Debug == true {
				log.Printf("DEBUG: container runtime %s not supported.\n", vars.ContainerRuntime)
				os.Exit(1)
			}
		}
	},
}

var forkCmd = &cobra.Command{
	Use:    "fork",
	Hidden: true,
	Short:  "execute command in container",
	Long:   `execute command in container`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		containers.ChildRcli(containerCommand, image)
		end := time.Now()
		log.Println(end.Sub(start))
	},
}
