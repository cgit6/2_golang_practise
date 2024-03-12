package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader // 讀取輸入

func main() {
	reader = bufio.NewReader(os.Stdin) // 讀取字串輸入
	userName := readString("你叫什麼名字?") // 獲取名字
	age := readInt("你幾歲?") // 獲取年齡

	fmt.Printf("你的名字是 %s 以及你的年齡是 %d",userName,age)
}

// 輸入函數
func prompt(){
	fmt.Print("->")
}

// 整理字串
func readString(s string) string {
	// 錯誤處理
	for {
		fmt.Println(s)
		prompt() 
	
		userInput,_ := reader.ReadString('\n') // 獲取的輸入文字
		userInput = strings.Replace(userInput,"\r\n","",-1) // windows 取代文字
		userInput = strings.Replace(userInput,"\n","",-1) // linux,mac 取代文字

		// 如果沒有輸入
		if userInput == ""{
			fmt.Println("請輸入一個整數")
		}else{
			return userInput // 名字
		}

	}
}

// 宣告一個問年齡的 function
func readInt(s string) int {
	
	for {
		fmt.Println(s)
		prompt() 
	
		userInput,_ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\r\n","",-1) // windows 取代文字
		userInput = strings.Replace(userInput,"\n","",-1) // linux,mac 取代文字
	
		num,err := strconv.Atoi(userInput)

		// 處理例外事件如果出現意外事件則不斷迴圈直到通過為止
		if err != nil{
			fmt.Println("請輸入一個整數")
		}else{
			return num // 年齡
		}
	}

}

