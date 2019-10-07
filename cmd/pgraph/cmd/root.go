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
	"log"
	"os"
	"sync"
	"time"

	"github.com/nboughton/go-roll"
	"github.com/spf13/cobra"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var wg sync.WaitGroup

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pgraph",
	Short: "Render a plot of one or more Dice strings as a png",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			dice, _   = cmd.Flags().GetStringArray("dice")
			labels, _ = cmd.Flags().GetStringArray("label")
			rolls, _  = cmd.Flags().GetInt("rolls")
			title, _  = cmd.Flags().GetString("title")
		)

		t := time.Now()

		// New plot
		pl, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		pl.Title.Text = title
		pl.X.Label.Text = "Rolled"
		pl.Y.Label.Text = "Probability (%)"

		pl.X.Tick.Marker = customTicks{}
		pl.Y.Tick.Marker = customTicks{}

		// Roll some dice and aggregate data
		var (
			argsLine []interface{}
			wait     = make(chan int, 2) // run two rolls concurrently at a time to speed thins up a little
		)

		for i, s := range dice {
			wg.Add(1)

			go func(i int, s string) {
				defer wg.Done()
				wait <- 1

				fmt.Println("rolling ", s)

				var (
					results roll.Results
					label   = dice[i]
				)

				if len(labels) > i {
					label = labels[i]
				}

				for i := 0; i < rolls; i++ {
					result, err := roll.FromString(s)
					if err != nil {
						log.Fatal(err)
					}

					results = append(results, result)
				}

				argsLine = append(argsLine, label, lineData(results, results.Min(), results.Max()))
				<-wait
			}(i, s)
		}

		wg.Wait()

		pl.Add(plotter.NewGrid())

		plotutil.AddLines(pl, argsLine...)
		pl.Legend.Top = true
		pl.Legend.Left = true

		// Save to png
		if err := pl.Save(20*vg.Centimeter, 15*vg.Centimeter, fmt.Sprintf("%s.png", title)); err != nil {
			log.Fatal(err)
		}
		fmt.Println("run time: ", time.Now().Sub(t).Round(time.Millisecond))
	},
}

func lineData(results roll.Results, min, max int) plotter.XYs {
	var (
		xLen = max - min + 1
		xy   = make(plotter.XYs, xLen)
	)

	for _, res := range results {
		n := res.Sum() - min
		xy[n].Y++
	}

	for i := range xy {
		xy[i].X = float64(i + min)
		xy[i].Y = xy[i].Y / float64(len(results)) * 100
	}

	return xy
}

type customTicks struct{}

func (customTicks) Ticks(min, max float64) []plot.Tick {
	var tks []plot.Tick

	for i := 0.; i < max; i++ {
		t := plot.Tick{Value: float64(i + 1)}

		switch {
		case max > 20 && max < 50:
			t.Label = label(i, 2)
		case max >= 50 && max < 100:
			t.Label = label(i, 5)
		case max >= 100:
			t.Label = label(i, int(max/4))
		default:
			t.Label = label(i, 1)
		}

		tks = append(tks, t)
	}

	return tks
}

func label(i float64, mod int) string {
	if int(i+1)%mod == 0 {
		return fmt.Sprintf("%d", int(i+1))
	}

	return ""
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().StringArrayP("dice", "d", []string{"2d6", "3d6", "4d6", "3d6Kh2", "4d6Kh2"}, "Dice strings to plot")
	RootCmd.Flags().StringArrayP("label", "l", []string{}, "Labels for plots, these are applied to their respective dice strings")
	RootCmd.Flags().IntP("rolls", "r", 100000, "Number of times to roll each dice set")
	RootCmd.Flags().StringP("title", "t", "graph", "Title of graph")
}
