package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func IntStdin() int {
	var input string
	for {
		fmt.Scan(&input)
		num, err := strconv.Atoi(input)

		if err == nil && (num >= 0 && num <= 3) {
			return num
		}

		fmt.Println("0~3までの数字を入力してください。")
	}
}

func Display(you, computer string) {
	fmt.Println("---------------")
	fmt.Printf("あなた：%sを出しました\n", you)
	fmt.Printf("相手：%sを出しました\n", computer)
	fmt.Println("---------------")
}

func main() {
	choice := []string{"グー", "チョキ", "パー"}
	direction := []string{"上", "下", "右", "左"}

	for {
		winnerflg := false
		fmt.Println("0(グー)1(チョキ)2(パー)3(戦わない)")
		fmt.Println("じゃんけん...")
		you := IntStdin()
		computer := rand.Intn(3)

		if you == 3 {
			fmt.Println("ゲームを終了しました。")
			return
		}

		fmt.Printf("ホイ！\n")
		Display(choice[you], choice[computer])

		for you == computer {
			fmt.Println("あいこで")
			fmt.Println("0(グー)1(チョキ)2(パー)3(戦わない)")
			you = IntStdin()
			computer = rand.Intn(3)
			if you == 3 {
				fmt.Println("ゲームを終了しました。")
				return
			}
			Display(choice[you], choice[computer])
		}

		if you == 0 && computer == 1 || you == 1 && computer == 2 || you == 2 && computer == 0 {
			winnerflg = true
		}

		fmt.Println("あっち向いて〜")
		fmt.Println("0(上)1(下)2(左)3(右)")
		you = IntStdin()
		computer = rand.Intn(4)

		Display(direction[you], direction[computer])
		if winnerflg && you == computer {
			fmt.Printf("あなたの勝ちです!\n\n")
		} else if !winnerflg && you == computer {
			fmt.Printf("CPの勝ちです!\n\n")
		}
	}
}
