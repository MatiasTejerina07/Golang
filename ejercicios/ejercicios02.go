package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func InputNumber() string {
	scan := bufio.NewScanner(os.Stdin)
	var err error
	var number int
	var text string

	for {
		fmt.Println("Ingrese el número:")
		if scan.Scan() {
			number, err = strconv.Atoi(scan.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}

	}
	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", number, i, number*i)
	}
	return text
}
