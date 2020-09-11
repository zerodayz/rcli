package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zerodayz/rcli/helpers/containers"
	"github.com/zerodayz/rcli/vars"
	"log"
	"time"
)

var (
	containerCommand string
)

func init() {
	rootCmd.AddCommand(ContainerCmd)
	ContainerCmd.PersistentFlags().BoolVarP(&vars.Debug, "debug", "d", false, "enable debug")

	ContainerCmd.AddCommand(runContainerCmd)
	runContainerCmd.Flags().StringVarP(&containerCommand, "command", "c", "", "command")
	runContainerCmd.MarkFlagRequired("containerCommand")

	runContainerCmd.AddCommand(forkCmd)
	forkCmd.Flags().StringVarP(&containerCommand, "command", "c", "", "command")
	forkCmd.MarkFlagRequired("containerCommand")
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
		containers.Run(containerCommand)
	},
}

var forkCmd = &cobra.Command{
	Use:   "fork",
	Hidden: true,
	Short: "execute command in container",
	Long:  `execute command in container`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		containers.Child(containerCommand)
		end := time.Now()
		log.Println(end.Sub(start))
	},
}
