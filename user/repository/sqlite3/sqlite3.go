package sqlite3

import (
	"awesomeProject1/model"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


type userRepository struct {
	sqlClient *sql.DB
}

func NewUserRepository(db *sql.DB) userRepository {
	return userRepository{db}
}

func (ur userRepository) Get(userID int) (*model.User, error) {
	row := ur.sqlClient.QueryRow(`SELECT * FROM USERS WHERE ID=?`, userID)
	var id int
	var name string
	err := row.Scan(&id, &name)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Id:   int(id),
		Name: name,
	}, nil
}

func (ur userRepository) Add(user *model.User) (*model.User, error) {
	result, err := ur.sqlClient.Exec(`INSERT INTO USERS (NAME) VALUES (?)`, user.Name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.Id = int(id)

	return user, nil
}

func (ur userRepository) Edit(user *model.User) error {
	_, err := ur.sqlClient.Exec(`UPDATE USERS SET NAME=? WHERE ID=?`, user.Name, user.Id)
	return err
}

func (ur userRepository) Delete(userID int) error {
	_, err := ur.sqlClient.Exec(`DELETE FROM USERS WHERE ID=?`, userID)
	return err
}
