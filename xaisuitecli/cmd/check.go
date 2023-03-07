/*
Copyright © 2023 Shreyan Mitra <xaisuite@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

        "os"

        "strings"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check whether a model, data, or explainer is valid before using it for training",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking...")
                model, _ := cmd.Flags().GetString("model") 
                data, _ := cmd.Flags().GetString("data")
                explainers, _ := cmd.Flags().GetString("explainers")
                if model == "" && data == "" && explainers == ""{
                    fmt.Println("Must pass a flag to check command; no valid flags provided")
                    cmd.Help()
                    os.Exit(1)
                }   
                if model != "" {
                    checkModel(model)
                } else if data != "" {
                    checkData(data)
                } else{
			explainerList := strings.Split(explainers, " ")
			for _, element := range explainerList {
				if !(checkExplainer(element)){
					fmt.Println("Not a valid explainer " + string(element))
					os.Exit(1)
                                }
                        }

                }
                 
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

        checkCmd.Flags().StringP("model", "m", "", "Model to check")
        checkCmd.Flags().StringP("data", "d", "", "Filepath for data to check")
        checkCmd.Flags().StringP("explainers", "e", "", "Name of explainers to check. To use multiple explainers, include all names in double quotes, separated by space.")

        checkCmd.MarkFlagsMutuallyExclusive("model", "data", "explainers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
