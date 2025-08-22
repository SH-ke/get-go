package main

import "fmt"

func reset(board *[8][8]rune) {
	// 初始化空棋盘
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = ' '
		}
	}

	// 设置白方后排
	board[0][0] = 'r' // 车
	board[0][1] = 'n' // 马
	board[0][2] = 'b' // 象
	board[0][3] = 'q' // 后
	board[0][4] = 'k' // 王
	board[0][5] = 'b' // 象
	board[0][6] = 'n' // 马
	board[0][7] = 'r' // 车

	// 设置白方兵
	for i := 0; i < 8; i++ {
		board[1][i] = 'p' // 兵
	}

	// 设置黑方兵
	for i := 0; i < 8; i++ {
		board[6][i] = 'P' // 兵（大写表示黑方）
	}

	// 设置黑方后排
	board[7][0] = 'R' // 车
	board[7][1] = 'N' // 马
	board[7][2] = 'B' // 象
	board[7][3] = 'Q' // 后
	board[7][4] = 'K' // 王
	board[7][5] = 'B' // 象
	board[7][6] = 'N' // 马
	board[7][7] = 'R' // 车
}

func printBoard(board *[8][8]rune) {
	// 打印列标记（a-h）
	fmt.Print("   ")
	for j := 0; j < 8; j++ {
		// %c 应该占4个格子，居中展示，用空白填充
		fmt.Printf("  %c ", 'a'+j)
	}
	fmt.Println()

	// 打印上下边框的函数应当复用，封装成一个函数
	boarder := func() {
		fmt.Print("   ")
		for j := 0; j < 8; j++ {
			fmt.Print(" ───")
		}
		fmt.Println()
	}

	// 打印上边框
	boarder()

	// 打印棋盘内容
	for i := 0; i < 8; i++ {
		// 打印行标记（8-1）
		fmt.Printf(" %d ", 8-i)

		for j := 0; j < 8; j++ {
			fmt.Print("│")
			if board[i][j] != ' ' {
				fmt.Printf(" %c ", board[i][j])
			} else {
				// 使用黑白方块填充空白格子
				if (i+j)%2 == 0 {
					fmt.Print("░░░") // 白色格子
				} else {
					fmt.Print("▓▓▓") // 黑色格子
				}
			}
		}
		fmt.Println("│")

		// 打印行分隔线
		if i < 7 {
			fmt.Print("   ")
			for j := 0; j < 8; j++ {
				fmt.Print("├───")
			}
			fmt.Println("┤")
		}
	}

	// 打印下边框
	boarder()
}

func main() {
	var board [8][8]rune
	reset(&board)
	printBoard(&board)
}
