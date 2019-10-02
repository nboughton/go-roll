// Copyright © 2019 Nick Boughton <nicholasboughton@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/nboughton/go-roll"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	tw      = tabwriter.NewWriter(os.Stdout, 1, 2, 1, ' ', 0)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dprob",
	Short: "Calculate the probability of rolling a desired set of numbers from a dice string",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			dice, _   = cmd.Flags().GetStringArray("dice")
			labels, _ = cmd.Flags().GetStringArray("label")
			want, _   = cmd.Flags().GetIntSlice("want")
			rolls, _  = cmd.Flags().GetInt("rolls")
		)

		for i, s := range dice {
			var results roll.Results
			for j := 0; j < rolls; j++ {
				result, err := roll.FromString(s)
				if err != nil {
					fmt.Println(err)
					return
				}

				results = append(results, result)
			}

			p := 0
			for _, r := range results {
				for _, w := range want {
					if r.Sum() == w {
						p++
					}
				}
			}
			l := s
			if len(labels) > i {
				l = labels[i]
			}
			fmt.Fprintf(tw, "%s\t==\t%v\t%.3f%%\n", l, want, float64(p)/float64(rolls)*100)
		}
		tw.Flush()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringArrayP("dice", "d", []string{"1d10", "2d10Kl1"}, "Dice strings to test")
	rootCmd.Flags().StringArrayP("label", "l", []string{}, "Labels for plots, these are applied to their respective dice strings")
	rootCmd.Flags().IntP("rolls", "r", 100000, "Number of times to roll each dice set")
	rootCmd.Flags().IntSliceP("want", "w", []int{9, 10}, "Numbers to test for")
}
