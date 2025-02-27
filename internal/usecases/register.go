package usecases

import (
	"math/rand"
	"fmt"
	"github.com/hassanpdn/task-manager-go/internal/entities"
)

func Register(email, password string) entities.User {
	id := email
	user := entities.User{
		ID:      rand.Int(),
		Email:   email,
		Passkey: password,
	}
	fmt.Println(id, email, password)
	return user
}
