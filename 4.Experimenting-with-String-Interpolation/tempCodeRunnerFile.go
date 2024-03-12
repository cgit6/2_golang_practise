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