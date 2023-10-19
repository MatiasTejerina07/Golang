package arreglos_slice

import "fmt"

var tabla []int = []int{2, 5, 4, 7}
var arreglo [10]int = [10]int{6, 92, 27, 5}

func ShowSlice() {
	fmt.Println(tabla)

	porcion := arreglo[3:]
	porcion2 := arreglo[:3]
	porcion3 := arreglo[5:7]
	fmt.Println(porcion, porcion2, porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n largo %d, capacidad %d", len(nums), cap(nums))
}
