package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/awh/egctl/pkg/energenie"
	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio"
)

var offCodes = map[string][4]rpio.State{
	"one":   [4]rpio.State{rpio.Low, rpio.High, rpio.High, rpio.High},
	"two":   [4]rpio.State{rpio.Low, rpio.High, rpio.High, rpio.Low},
	"three": [4]rpio.State{rpio.Low, rpio.High, rpio.Low, rpio.High},
	"four":  [4]rpio.State{rpio.Low, rpio.High, rpio.Low, rpio.Low},
	"all":   [4]rpio.State{rpio.Low, rpio.Low, rpio.High, rpio.High},
}

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("socket argument required")
		}

		if code, ok := offCodes[args[0]]; ok {
			if err := energenie.Execute(code); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatalf("bad socket argument")
		}

	},
}

func init() {
	RootCmd.AddCommand(offCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// offCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// offCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
