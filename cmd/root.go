package cmd

import (
    "fmt"
    "os"
    "time"
	"sync"

    "pro7_finder/finder"

    "github.com/spf13/cobra"
)

var (
    dir      string
    fileName string
    method   string
)

var rootCmd = &cobra.Command{
    Use:   "finder",
    Short: "A fast file finder CLI tool",
    Long:  `A CLI tool to search for files in a directory using concurrent algorithms.`,
    Run: func(cmd *cobra.Command, args []string) {
        if dir == "" || fileName == "" {
            fmt.Println("Error: both --dir and --name flags are required")
            cmd.Help()
            os.Exit(1)
        }

        f := &finder.Finder{}
        start := time.Now()

        switch method {
        case "job":
            f.JobFinder(dir, fileName)
        case "sem":
            var wg sync.WaitGroup
            wg.Add(1)
            go f.SemFinder(dir, fileName, &wg)
            wg.Wait()
        default:
            f.BasicFinder(dir, fileName)
        }

        fmt.Printf("Search completed in %v\n", time.Since(start))
        fmt.Printf("Found %d result(s):\n", len(f.Res))
        for _, path := range f.Res {
            fmt.Printf("  → %s\n", path)
        }
    },
}

func init() {
    rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "Directory to search in")
    rootCmd.Flags().StringVarP(&fileName, "name", "n", "", "File name to search for")
    rootCmd.Flags().StringVarP(&method, "method", "m", "job", "Search method: job or sem")
    rootCmd.MarkFlagRequired("name")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}