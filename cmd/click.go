/*
Copyright © 2022 Kevin Jayne <kevin.jayne@icloud.com>
*/
package cmd

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var clickCmd = &cobra.Command{
	Use:   "click",
	Short: "Click the mouse at current mouse position",
	Long:  `Click the mouse at current mouse position`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Clicking the mouse at current mouse position")
		if loop > 0 {
			for {
				sleepMilliseconds(loop)
				robotgo.Click()
			}
		}
		robotgo.Click()
	},
}

