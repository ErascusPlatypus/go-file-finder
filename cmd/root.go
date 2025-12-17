package cmd

import (
	"fmt"
	"os"
	"pro7_finder/finder"
	"sort"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
    dir      string
    fileName string
    method   string
	regexExp string
	excludeDir [] string
)

var rootCmd = &cobra.Command{
    Use:   "file-finder",
    Short: "A fast file finder CLI tool",
    Long:  `A CLI tool to search for files in a directory using concurrent algorithms and compare execution speeds.`,
    Run: func(cmd *cobra.Command, args []string) {
        if dir == "" || fileName == "" {
            fmt.Println("Error: both --dir and --name flags are required")
            cmd.Help()
            os.Exit(1)
        }

        f := &finder.Finder{}
        start := time.Now()

		if len(excludeDir) > 0 {
			f.ToMap(excludeDir)
		}

		regexFlag := false 

		if regexExp != "" {
			regexFlag = true 
			f.SetRegex(regexExp)
		}

        switch method {
        case "job":
            f.JobFinder(dir, fileName, regexFlag)
        case "sem":
            var wg sync.WaitGroup
            wg.Add(1)
            go f.SemFinder(dir, fileName, regexFlag, &wg)
            wg.Wait()
        default:
            f.BasicFinder(dir, fileName)
        }

        fmt.Printf("Search completed in %v\n", time.Since(start))
		sort.Strings(f.Res)
        fmt.Printf("Found %d result(s):\n", len(f.Res))
        for _, path := range f.Res {
            fmt.Printf("-> %s\n", path)
        }
    },
}

func init() {
    rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "Directory to search in")
    rootCmd.Flags().StringVarP(&fileName, "name", "n", "", "File name to search for")
    rootCmd.Flags().StringVarP(&method, "method", "m", "job", "Search method: job or sem")
	rootCmd.Flags().StringVarP(&regexExp, "regex", "r", "", "Search by regex pattern")
	rootCmd.Flags().StringSliceVar(&excludeDir, "exclude", [] string {}, "Exclude search from")

    rootCmd.MarkFlagRequired("name")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}