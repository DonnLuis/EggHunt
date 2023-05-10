package tests

import (
    "fmt"
    "testing"

    "github.com/spf13/cobra"
)

func TestMain(m *testing.M) {
    // Set up any test data or configurations here
    // ...

    // Run the tests
    m.Run()
}

func TestRootCommand(t *testing.T) {
    cmd := &cobra.Command{
        Use: "hunter",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Root command executed!")
        },
    }

    if err := cmd.Execute(); err != nil {
        t.Errorf("Error executing root command: %v", err)
    }
}

func TestSubcommand(t *testing.T) {
    cmd := &cobra.Command{
        Use:   "sub",
        Short: "A subcommand",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Subcommand executed!")
        },
    }

    rootCmd := &cobra.Command{Use: "hunter"}
    rootCmd.AddCommand(cmd)

    if err := rootCmd.Execute(); err != nil {
        t.Errorf("Error executing subcommand: %v", err)
    }
}
