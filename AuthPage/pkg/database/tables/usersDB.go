package tables

import (
	"VK/pkg/model"
	"github.com/jackc/pgx"
	"golang.org/x/sys/windows"
	"log"
)

type UsersDB struct {
}

func (*UsersDB) GetUser(pool *pgx.Conn, username string) (model.User, error) {
	var (
		user model.User
	)
	rows, err := pool.Query("SELECT * FROM vk_users WHERE username = $1", username)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	if rows.Next() {
		value, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		user.ParseData(value)
		return user, nil
	}
	return user, windows.ERROR_OBJECT_NOT_FOUND
}
