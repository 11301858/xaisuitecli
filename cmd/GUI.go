/*
Copyright © 2023 Shreyan Mitra <xaisuite@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

        "os"

        "os/exec"

        "log"
)

// GUICmd represents the GUI command
var GUICmd = &cobra.Command{
	Use:   "gui",
	Short: "Opens XAISuite's GUI",
	Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Installing XAISuite...")
		install := exec.Command("zsh", "-c", "pip install XAISuite==1.0.8")
                install.Stdin = os.Stdin
                install.Stdout = os.Stdout
                install.Stderr = os.Stderr
                err_install := install.Run()
                if err_install != nil {
                    log.Fatalf("Installing XAISuite failed with %s\n", err_install)
                }

		fmt.Println("Opening GUI...")
                gui := exec.Command("zsh", "-c", "python -c 'import xaisuitegui.runner'")
                gui.Stdin = os.Stdin
                gui.Stdout = os.Stdout
                gui.Stderr = os.Stderr
                err:= gui.Run()
                if err != nil {
                    log.Fatalf("Opening XAISuite GUI failed with %s\n", err)
                }
	},
}

func init() {
	rootCmd.AddCommand(GUICmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GUICmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GUICmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
