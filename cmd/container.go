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
	image 			 string
)

func init() {
	rootCmd.AddCommand(ContainerCmd)
	ContainerCmd.PersistentFlags().BoolVarP(&vars.Debug, "debug", "d", false, "enable debug")

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
		containers.Run(containerCommand, image)
	},
}

var forkCmd = &cobra.Command{
	Use:   "fork",
	Hidden: true,
	Short: "execute command in container",
	Long:  `execute command in container`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		containers.Child(containerCommand, image)
		end := time.Now()
		log.Println(end.Sub(start))
	},
}
