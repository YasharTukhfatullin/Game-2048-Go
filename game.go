package main 
import "fmt"

func main()  {
	var board [4][4]int
	PrintBoard(board)
}

func PrintBoard(board) {
	fmt.Println(board)
}
func printBoard(board [4][4]int) {
	for rowIndex, row := range board {
		for colIndex, col := range row {
			if col == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(col)
			}
			if colIndex < 3 {
				fmt.Println("|")
			}
		}
		fmt.Println()
		if rowIndex < 3 {
			fmt.Println("---------")
		}
	}
	fmt.Println()
}
