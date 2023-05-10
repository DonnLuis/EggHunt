package main
import (
	"Fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "hunter",
		Short: "the hunter",
		Long: "the hunter used to hunt for xyz",
		Run: func(cmd *cobra.Command, args []strings) {
			fmt.Println("this is your hunter")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
