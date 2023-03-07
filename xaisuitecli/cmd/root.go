/*
Copyright © 2023 Shreyan Mitra <xaisuite@gmail.com>

*/
package cmd

import (
	"os"

        "fmt"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "XAISuite CLI",
	Short: "Train and explain machine learning models",
	Long: `This is the command-line interface for XAISuite, an unified platform for training and explaining machine learning models. Please check the github repository at github.com/11301858/xaisuitecli for details and use instructions`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
        fmt.Println("Welcome to XAISuite's CLI!")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.XAISuite.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


