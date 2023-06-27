/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/XANi/loremipsum"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/spf13/cobra"
	"time"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Present a simple dashboard",
	Long:  `Show dashboard`,
	Run: func(cmd *cobra.Command, args []string) {
		terminalLayer, err := tcell.New(tcell.ColorMode(terminalapi.ColorMode256), tcell.ClearStyle(cell.ColorYellow, cell.ColorNavy))
		if err != nil {
			fmt.Println(err)
		}
		defer terminalLayer.Close()

		rollingText, err := text.New(text.RollContent(), text.WrapAtWords())
		if err != nil {
			panic(err)
		}

		value := make(chan string)
		logValue := make(chan string)
		go roll(rollingText, value)

		leftContainer := container.Left(
			container.Border(linestyle.Light),
			container.PlaceWidget(rollingText),
			container.PaddingLeft(1),
			container.PaddingRight(1),
			container.PaddingTop(1),
			container.PaddingBottom(1),
		)
		rollingLog, err := text.New(text.RollContent(), text.WrapAtWords())
		if err != nil {
			panic(err)
		}

		go roll(rollingLog, logValue)

		rightContainer := container.Right(
			container.SplitHorizontal(
				container.Top(
					container.Border(linestyle.Light),
					container.PlaceWidget(rollingLog),
					container.PaddingLeft(1),
					container.PaddingRight(1),
					container.PaddingTop(1),
					container.PaddingBottom(1),
				),
				container.Bottom(),
			),
		)

		containerLayer, _ := container.New(
			terminalLayer,
			container.SplitVertical(
				leftContainer,
				rightContainer,
				container.SplitPercent(60),
			),
		)

		ctx, cancel := context.WithCancel(context.Background())
		quitter := func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' {
				cancel()
			}
		}

		go gobshite(value)
		go logging(logValue)

		if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.KeyboardSubscriber(quitter)); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dashboardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dashboardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func roll(rollingText *text.Text, value <-chan string) {
	for {
		select {
		case v := <-value:
			err := rollingText.Write(v)
			if err != nil {
				panic(err)
			}
		default:
			// no value on the channel, continue without blocking
		}
	}
}

func gobshite(value chan string) {
	l := loremipsum.New()
	for i := 0; i < 100; i++ {
		value <- l.Paragraph() + "\n\n"

		time.Sleep(100 * time.Millisecond)
	}
}

func logging(value chan string) {
	l := loremipsum.New()
	for i := 0; i < 100; i++ {
		value <- l.Paragraph() + "\n\n"

		time.Sleep(400 * time.Millisecond)
	}
}
