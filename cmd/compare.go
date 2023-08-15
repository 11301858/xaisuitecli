/*
Copyright © 2023 Shreyan Mitra <xaisuite@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

        "os"

        "log"

        "os/exec"

        "strings"
)

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compares two explanatory files generated by XAISuite",
	Long: `Pathnames are CSV files that generally have the form '<Explainer> ImportanceScores - <Model> <Target>.csv'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Comparing...")
                fmt.Println("Checking data files...")
                for _, element := range args {
                    if !(checkData(element)) {
                        fmt.Println("Not a valid data " + string(element))
                        os.Exit(1)
                    }
                }
                fmt.Println("Initializing Commands...")
                files := strings.Join(args[:], "\", \"")
                compare_python := exec.Command("zsh", "-c", "python -c import pkg_resources; pkg_resources.require('XAISuite==1.0.8'); from xaisuite import*;compare_explanations([\"" + files + "\"])")
                fmt.Println(compare_python)
                compare_python.Stdin = os.Stdin
                compare_python.Stdout = os.Stdout
	        compare_python.Stderr = os.Stderr
                err := compare_python.Run()
	        //out, err := run_python.Output()
	        if err != nil {
		  log.Fatalf("Running XAISuite Analyzer failed with %s\n", err)
	        }
                //fmt.Println(string(out))

                                
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
