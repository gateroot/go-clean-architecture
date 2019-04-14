package repository

import (
	"awesomeProject1/model"
	"awesomeProject1/user"
	"awesomeProject1/user/repository"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var repo user.UserRepository = repository.NewUserRepository()
	user := model.User{
		Name: "Taro",
	}
	u, err := repo.Add(&user)
	if err != nil {
		fmt.Printf("add failed. : %s\n", err)
		return
	}
	fmt.Printf("add success. user id : %+v\n", u)

	u, err = repo.Get(u.Id)
	if err != nil {
		fmt.Printf("get failed. : %s\n", err)
		return
	}
	fmt.Printf("get success. user id : %d, name : %s\n", u.Id, u.Name)

	u.Name = "Shigeo"

	err = repo.Edit(u)
	if err != nil {
		fmt.Printf("edit failed. : %s\n", err)
		return
	}
	fmt.Printf("edit success.\n")

	u, err = repo.Get(u.Id)
	if err != nil {
		fmt.Printf("get failed. : %s\n", err)
		return
	}
	fmt.Printf("get success. user id : %d, name : %s\n", u.Id, u.Name)
}