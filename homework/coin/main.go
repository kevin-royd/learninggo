package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Peter", "Giana", "Adriano",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	for username, coin := range distribution {
		fmt.Printf("user:%s 分的金币:%d\n", username, coin)
	}
	fmt.Printf("剩余金币数:%d", left)
}

//分金币方法
func calcCoin(user string) int {
	var sum = 0
	for _, i := range user {
		switch i {
		case 'a', 'A':
			sum = sum + 1
		case 'e', 'E':
			sum = sum + 1
		case 'i', 'I':
			sum = sum + 2
		case 'o', 'O':
			sum = sum + 3
		case 'u', 'U':
			sum = sum + 5
		}
	}
	return sum
}

func dispatchCoin() int {
	var left = coins
	//	遍历用户slice得到单个用户
	for _, user := range users {
		//定义剩余的金币数
		allCoin := calcCoin(user)
		left = left - allCoin
		//	将分得的金币存入map中，该map在常量中已经初始化
		value, ok := distribution[user]
		if !ok {
			//	如果不存在则不能分到金币 allCoin也为0
			distribution[user] = allCoin
		} else {
			distribution[user] = value + allCoin
		}
	}
	return left
}
