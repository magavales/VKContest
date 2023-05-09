package tables

import (
	"TelegramBot/pkg/model"
	"github.com/jackc/pgx"
	"golang.org/x/sys/windows"
	"log"
)

type DataAccess struct {
}

func (de DataAccess) AddUserCredentials(pool *pgx.ConnPool, data model.UserCredentials) {
	_, err := pool.Query("INSERT INTO vk_storage (chatid, service, login, password, secret_key) VALUES ($1, $2, $3, $4, $5)", data.ChatID, data.Service, data.Login, data.Password, "")
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}
}

func (de DataAccess) GetUserCredentials(pool *pgx.ConnPool, id int64, service string) (model.UserCredentials, error) {
	var data model.UserCredentials
	rows, err := pool.Query("SELECT * FROM vk_storage WHERE chatid = $1 and service = $2", id, service)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	if rows.Next() {
		value, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		data.ParseColumns(value)
		return data, nil
	}
	return data, windows.ERROR_OBJECT_NOT_FOUND
}

func (de DataAccess) DeleteUserCredentials(pool *pgx.ConnPool, id int64, service string) error {
	_, err := pool.Query("DELETE FROM vk_storage WHERE chatid = $1 and service = $2", id, service)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
		return err
	}
	return nil
}
