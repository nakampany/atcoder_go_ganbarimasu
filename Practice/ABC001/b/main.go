package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	if m < 100 {
		fmt.Println("00")
	} else if 100 <= m && m <= 5000 {
		// %02d は，2桁の整数で出力するという意味．
		// mの値が3であれば，出力は「03」．
		// しかし，mの値が3.0であれば，出力は「3」．
		fmt.Printf("%02d\n", m/100)
	} else if 6000 <= m && m <= 30000 {
		fmt.Println(m/1000 + 50)
	} else if 35000 <= m && m <= 70000 {
		fmt.Println((m/1000-30)/5 + 80)
	} else if 70000 < m {
		fmt.Println(89)
	} else {
		fmt.Println("Error")
	}
}
