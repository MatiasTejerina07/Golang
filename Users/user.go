package users

import (
	"fmt"
	"time"

	"github.com/MatiasTejerina07/Golang/modelos"
)

func AltaUser() {
	u := new(modelos.User)
	u.AddUser(10, "Mat", time.Now(), true)

	fmt.Println(u)
}
