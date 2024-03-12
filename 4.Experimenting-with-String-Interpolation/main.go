package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader // 讀取輸入，星號是抓取該記憶體位置
 
// 先自定義一個結構，開頭大寫代表可以匯出
type User struct {
	UserName string
	Age int
	FavouriteNumber float64 
}


func main() {
	reader = bufio.NewReader(os.Stdin) // 讀取字串輸入
	
	var user User

	// 
	user.UserName = readString("你叫什麼名字?") // 獲取名字
	user.Age = readInt("你幾歲?") // 獲取年齡
	user.FavouriteNumber = readFloat("輸入最喜歡的數字") 

	// %.2f 取道小數點第二位
	fmt.Printf("你的名字是 %s 以及你的年齡是 %d,最喜歡的數字 %.2f",user.UserName, user.Age,user.FavouriteNumber)
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

// readFloat:填寫最喜歡的數字
func readFloat(s string) float64{
	for {
		fmt.Println(s) // ??

		prompt() // 箭頭

		// 這邊一樣是做字串處理
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\r\n","",-1) // windows 的結尾是 \r\n

		num,err := strconv.ParseFloat(userInput,64)
		// 終止條件
		if err != nil {
			fmt.Println("請輸入數字")
		} else {
			return num // 返回所輸入的合法數字
		}
	}

}

