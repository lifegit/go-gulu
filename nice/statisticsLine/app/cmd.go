/**
* @Author: TheLife
* @Date: 2021/7/22 下午4:18
 */
package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

var line StatisticsLine

var RestCmd = &cobra.Command{
	Use:     "statisticsLine",
	Short:   "statisticsLine a count the number of code lines",
	Example: `statisticsLine -p . -e -s .go`,
	Run: func(cmd *cobra.Command, args []string) {
		lineCount := line.Run()
		fmt.Printf("%d line", lineCount)
	},
}

func init() {
	RestCmd.Flags().StringVarP(&line.RootPath, "rootPath", "p", "./", "statistics directory")
	RestCmd.Flags().StringSliceVarP(&line.ExcludeDirs, "excludeDirs", "e", []string{}, "exclude directory")
	RestCmd.Flags().StringVarP(&line.SuffixName, "suffixName", "s", ".go", "extended name")
}
