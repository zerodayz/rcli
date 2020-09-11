package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "rcli",
	Short: "rcli multipurpose CLI",
	Long: `Project documentation is available at http://github.com/zerodayz/rcli`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func ShowLogo() {
	fmt.Println(`           ___________________ .____    .___`)
	fmt.Println(`Welcome to \______   \_   ___ \|    |   |   |`)
	fmt.Println(`            |       _/    \  \/|    |   |   |`)
	fmt.Println(`            |    |   \     \___|    |___|   |`)
	fmt.Println(`            |____|_  /\______  /_______ \___|`)
	fmt.Println( `                  \/        \/        \/   `)
	fmt.Println( ``)
	fmt.Println( `This software comes with ABSOLUTELY NO WARRANTY.`)
	fmt.Println( `Use at your own risk.`)
	fmt.Println( ``)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}