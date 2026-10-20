/*
Copyright © 2022 Kevin Jayne <kevin.jayne@icloud.com>
*/
package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var (
	xMin,
	xMax,
	yMin,
	yMax int
	loop float64

	clickClickPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mm",
	Short: "A CLI Mouse Move and Click Automation Tool",
	Long:  `Automate Mouse Movements and Clicks on your Desktop`,
	Run: func(cmd *cobra.Command, args []string) {
		getScreenSize()
		randRClick()
		if loop > 0 {
			for {
				sleepSeconds(loop)
				randRClick()
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(clickCmd)
	rootCmd.PersistentFlags().Float64Var(&loop, "loop", 0, "Toggle Automation Looping and set the delay in seconds between each iteration (fractions allowed, e.g. 0.3 = 300ms)")
	rootCmd.PersistentFlags().IntVar(&xMin, "xmin", 0, "Minimum Horizontal Screen Address, defaults to 0")
	rootCmd.PersistentFlags().IntVar(&xMax, "xmax", -1, "Maximum Horizontal Screen Address, defaults current screen width")
	rootCmd.PersistentFlags().IntVar(&yMin, "ymin", 0, "Minimum Vertical Screen Address, defaults to 0")
	rootCmd.PersistentFlags().IntVar(&yMax, "ymax", -1, "Maximum Vertical Screen Address, defaults current screen height")
}

func sleepSeconds(seconds float64) {
	time.Sleep(time.Duration(seconds * float64(time.Second)))
}

func randRClick() {
	robotgo.Move(randCoords())
	robotgo.Click("right", true)
	sleepSeconds(1)
	// get current mouse position
	x, y := robotgo.Location()
	robotgo.Move(x-10, y)
	robotgo.Click()
}

func randCoords() (int, int) {
	x := rand.Intn(xMax-xMin) + xMin
	y := rand.Intn(yMax-yMin) + yMin
	return x, y
}

func getScreenSize() {
	w, h := robotgo.GetScreenSize()
	if xMax == -1 {
		xMax = w
	}
	if yMax == -1 {
		yMax = h
	}
}
