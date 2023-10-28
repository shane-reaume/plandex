package cmd

import (
	"fmt"
	"os"

	"plandex/lib"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(plansCmd)
}

// plansCmd represents the list command
var plansCmd = &cobra.Command{
	Use:   "plans",
	Short: "List all available plans",
	Run:   plans,
}

func plans(cmd *cobra.Command, args []string) {
	plandexDir, _, err := lib.FindOrCreatePlandex()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	plans, err := os.ReadDir(plandexDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println("Available plans:")
	for _, p := range plans {
		if p.IsDir() {
			fmt.Println("-", p.Name())
		}
	}
}