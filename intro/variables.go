package intro

import "fmt"

func Variables() {

	fmt.Println("-----------")
	fmt.Println("Go Variables :")
	fmt.Println("-----------")

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
