package cmd

import (
	"github.com/spf13/cobra"

	"fmt"

	term "github.com/buger/goterm"
	"github.com/eiannone/keyboard"

	goju "github.com/hiraginoyuki/goju/puzzle"
)

var (
	puzzleCmd = &cobra.Command{
		Use:   "puzzle",
		Short: "commands related to puzzle",

		Run: func(cmd *cobra.Command, args []string) {
			puzzle := goju.Solved(4, 4)
			fmt.Printf("%v\n", puzzle)
		},
	}

	seed   int64
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "generate",
		Run: func(cmd *cobra.Command, args []string) {
			puzzle := goju.Gen(4, 4, seed)
			printPuzzle(&puzzle)
		},
	}
	playCmd = &cobra.Command{
		Use:   "play",
		Short: "play with puzzle interactively",
		Run: func(cmd *cobra.Command, args []string) {
			interactive(goju.Gen(4, 4, seed))
		},
	}
)

func init() {
	rootCmd.AddCommand(puzzleCmd)

	puzzleCmd.AddCommand(genCmd)
	puzzleCmd.AddCommand(playCmd)

	genCmd.Flags().Int64Var(&seed, "seed", 0, "seed")
	genCmd.MarkFlagRequired("seed")
	playCmd.Flags().Int64Var(&seed, "seed", 0, "seed")
	playCmd.MarkFlagRequired("seed")
}

func printPuzzle(p *goju.Puzzle) {
	for i := 0; i < p.Height(); i++ {
		first := true
		for _, v := range p.Pieces()[i*p.Width() : (i+1)*p.Width()] {
			if first {
				first = false
			} else {
				fmt.Printf(" ")
			}
			if v == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%2d", v)
			}
		}
		fmt.Printf("\n")
	}
}

func interactive(p goju.Puzzle) {
	term.Clear()
	for {
		term.MoveCursor(1, 1)
		term.Flush()

		printPuzzle(&p)
		term.MoveCursorDown(1)
		term.Flush()

		key, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		term.Printf("%c", key)

		if key == 'q' {
			return
		}
		if x, y, ok := keyToXY(key); ok {
			p.SlideFrom(x, y)
		}

		term.Flush()
	}
}

func keyToXY(key rune) (uint, uint, bool) {
	switch key {
		case '4':
			return 0, 0, true
		case '5':
			return 1, 0, true
		case '6':
			return 2, 0, true
		case '7':
			return 3, 0, true
		case 'r':
			return 0, 1, true
		case 't':
			return 1, 1, true
		case 'y':
			return 2, 1, true
		case 'u':
			return 3, 1, true
		case 'f':
			return 0, 2, true
		case 'g':
			return 1, 2, true
		case 'h':
			return 2, 2, true
		case 'j':
			return 3, 2, true
		case 'v':
			return 0, 3, true
		case 'b':
			return 1, 3, true
		case 'n':
			return 2, 3, true
		case 'm':
			return 3, 3, true
		default:
			return 0, 0, false
	}
}
