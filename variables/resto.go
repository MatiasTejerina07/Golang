package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float32
var Date time.Time

func RestVariables() {
	Name = "Matias"
	Status = true
	Salary = 1000.09
	Date = time.Now()
	fmt.Println(Name, Status, Salary, Date)
}

func ConverText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
