package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
	// 匯入套件
)

// 現在還是有點問題
func main() {
	err := keyboard.Open() // 錯誤提示，如果有錯誤舊報錯如果沒有就nil
	if err != nil {
		log.Fatal(err) // 紀錄錯誤?
	}

	// 所以說它是一個匿名函數而且立刻執行
	defer func(){
		_ = keyboard.Close()
	}() // defer 它會確保這個函數會最後執行

	fmt.Println("MENU")
	fmt.Println("------------------------")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - 結束")
	fmt.Println("選擇後按 Enter")

	// 咖啡選單
	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"
	
	for {
		// 這個怪怪的
		char, _ ,err := keyboard.GetSingleKey() // 監聽目前按下去的鍵，返還三個值

		// 錯誤處理
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		// fmt.Println(i) 
		fmt.Printf("選擇: %s", coffees[i])

		// 如果按下 q 或 Q 結束程式
		if char == 'q'||char == 'Q' {
			break 
		}
	}
	fmt.Println("運行結束...")
}