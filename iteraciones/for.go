package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func JumpsFromTo() {
	for a := 0; a <= 100; a += 20 {
		fmt.Println(a)
	}
}

func BackJump() {
	for i := 100; i > 10; i -= 5 {
		fmt.Println(i)
	}
}

func AllExcept() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			continue
		}
		fmt.Print(i)
	}
}
