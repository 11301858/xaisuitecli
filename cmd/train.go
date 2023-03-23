/*
Copyright © 2023 Shreyan Mitra <xaisuite@gmail.com>

*/
package cmd


import (
	"fmt"

	"github.com/spf13/cobra"

        "strings"

        "os"
     
        "log"

        "os/exec"

        "errors"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Trains a machine learning model",
	Long: `Uses XAISuite to train a model with a given dataset. The --model and --data subcommands are required`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializaing XAISuite training...")
		model, _ := cmd.Flags().GetString("model")
		data, _ := cmd.Flags().GetString("data")
		target, _ := cmd.Flags().GetString("target")
		explainers, _ := cmd.Flags().GetString("explainers")
                compare, _ := cmd.Flags().GetBool("compare")
                fmt.Println("Checking provided values...")

		if !(checkModel(model) && checkData(data)){
                        fmt.Println("Initial Checks Failed.")
			os.Exit(1)
		}
                if explainers == "" && compare {
                    fmt.Println("Cannot specify -c flag for explanation comparison if -e flag for explainers is not set.")
                    cmd.Help()
                    os.Exit(1)
                }
                fmt.Println("Checking explainers...")
		var explainerList = []string{}
		if explainers != "" {
			explainerList = strings.Split(explainers, " ")
			for _, element := range explainerList {
				if !(checkExplainer(element)){
					fmt.Println("Not a valid explainer " + string(element))
					os.Exit(1)
				}
			}
			
		}
                fmt.Println("Installing XAISuite...")
		install := exec.Command("zsh", "-c", "pip install XAISuite --upgrade")
                install.Stdin = os.Stdin
                install.Stdout = os.Stdout
                install.Stderr = os.Stderr
                err_install := install.Run()
                if err_install != nil {
                    log.Fatalf("Installing XAISuite failed with %s\n", err_install)
                }
                fmt.Println("Initializing commands to run...")
		explainerList_string := "[\"" + strings.Join(explainerList , "\" , \"") + "\"]"
		var run_python = exec.Command ("zsh", "-c", "python -c 'from xaisuite import*;train_and_explainModel(\"" + model + "\", load_data_CSV(\"" + data + "\", \"" + target + "\")," + explainerList_string + ")'")
		if explainers == "" {
			run_python = exec.Command ("zsh", "-c", "python -c 'from xaisuite import*;train_and_explainModel(\"" + model + "\", load_data_CSV(\"" + data + "\", \"" + target + "\"))'")
		}
		//fmt.Println(run_python)
                run_python.Stdin = os.Stdin
	        run_python.Stdout = os.Stdout
	        run_python.Stderr = os.Stderr
                err := run_python.Run()
	        //out, err := run_python.Output()
	        if err != nil {
		  log.Fatalf("Running XAISuite Trainer failed with %s\n", err)
	        }
                //fmt.Println(string(out))

                if compare {
                    var compare_str = "" //"["
                    fmt.Println("Collecting data files for comparison...")
                    for _, fileheader := range explainerList {
                        fmt.Println("\tIdentified " + "\"" + fileheader + " Importance Scores " + model + " " + target + ".csv\"")

                        compare_str += fileheader + " Importance Scores " + model + " " + target + ".csv-"
                    }
                    //compare_str += "]"
                    rootCmd.SetArgs(append([]string{"compare"}, strings.Split(compare_str, "-")...))
                    rootCmd.Execute()
                }

                
	    },
}

func checkModel(model string) bool{
         
        acceptedModels := strings.Split("SVC NuSVC LinearSVC SVR NuSVR LinearSVR AdaBoostClassifier AdaBoostRegressor BaggingClassifier BaggingRegressor ExtraTreesClassifier ExtraTreesRegressor GradientBoostingClassifier GradientBoostingRegressor RandomForestClassifier RandomForestRegressor StackingClassifier StackingRegressor VotingClassifier VotingRegressor HistGradientBoostingClassifier HistGradientBoostingRegressor GaussianProcessClassifier GaussianProcessRegressor IsotonicRegression KernelRidge LogisticRegression LogisticRegressionCV PassiveAgressiveClassifier Perceptron RidgeClassifier RidgeClassifierCV SGDClassifier SGDOneClassSVM LinearRegression Ridge RidgeCV SGDRegressor ElasticNet ElasticNetCV Lars LarsCV Lasso LassoCV LassoLars LassoLarsCV LassoLarsIC OrthogonalMatchingPursuit OrthogonalMatchingPursuitCV ARDRegression BayesianRidge MultiTaskElasticNet MultiTaskElasticNetCV MultiTaskLasso MultiTaskLassoCV HuberRegressor QuantileRegressor RANSACRegressor TheilSenRegressor PoissonRegressor TweedieRegressor GammaRegressor PassiveAggressiveRegressor BayesianGaussianMixture GaussianMixture OneVsOneClassifier OneVsRestClassifier OutputCodeClassifier ClassifierChain RegressorChain MultiOutputRegressor MultiOutputClassifier BernoulliNB CategoricalNB ComplementNB GaussianNB MultinomialNB KNeighborsClassifier KNeighborsRegressor BernoulliRBM MLPClassifier MLPRegressor DecisionTreeClassifier DecisionTreeRegressor ExtraTreeClassifier ExtraTreeRegressor", " ")



        for _, str := range acceptedModels {
            if strings.Compare(model, str) == 0{
                fmt.Println("Model " + model + " is valid.")
                return true
            }
        }
        fmt.Println("Model " + model + " is not valid or not provided.")
        return false

}

func checkData(data string) bool{
        extension := data[len(data)-4:]
        if strings.Compare(extension, ".csv") != 0{
            fmt.Println("Data is not valid or not provided: Not a CSV file. Make sure your file ends with .csv")
            return false
        }
        if _, err := os.Stat(data); err == nil {
            fmt.Println("Data is valid.")
            return true

        } else if errors.Is(err, os.ErrNotExist) {
            fmt.Println("Data is not valid or not provided: Not a correct file path.")
            return false
        } else {
            fmt.Println("Data is not valid or not provided: Data check failed. Please try again.")
            return false

}
}

func checkExplainer(explainer string) bool{
       acceptedExplainers := strings.Split("shap lime pdp ale mace sensitivity", " ")
       for _,str := range acceptedExplainers {
           if strings.Compare(explainer, str) == 0{
               return true
           }
       }
       return false 
}


func init() {
	rootCmd.AddCommand(trainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
        trainCmd.Flags().BoolP("compare", "c", false, "Compare generated explanations")
        trainCmd.Flags().StringP("model", "m", "", "Model to use")
        trainCmd.Flags().StringP("data", "d", "", "Filepath for data to use")
        trainCmd.Flags().StringP("target", "t", "", "Name of target variable")
        trainCmd.Flags().StringP("explainers", "e", "", "Name of explainers to use. To use multiple explainers, include all names in double quotes, separated by space.")


        trainCmd.MarkFlagRequired("model")
        trainCmd.MarkFlagRequired("data")
        trainCmd.MarkFlagRequired("target")
       }
