package tables

import (
	"TelegramBot/pkg/model"
	"github.com/jackc/pgx"
	"golang.org/x/sys/windows"
	"log"
)

type DataEntity struct {
}

func (de DataEntity) AddData(pool *pgx.Conn, data model.Data) {
	_, err := pool.Query("INSERT INTO vk_storage (chatid, service, login, password, secret_key) VALUES ($1, $2, $3, $4, $5)", data.ChatID, data.Service, data.Login, data.Password, "")
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}
}

func (de DataEntity) GetData(pool *pgx.Conn, id int64, service string) (model.Data, error) {
	var data model.Data
	rows, err := pool.Query("SELECT * FROM vk_storage WHERE chatid = $1 and service = $2", id, service)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	if rows.Next() {
		value, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		data.ParseData(value)
		return data, nil
	}
	return data, windows.ERROR_OBJECT_NOT_FOUND
}

func (de DataEntity) DelData(pool *pgx.Conn, id int64, service string) {
	_, err := pool.Query("DELETE FROM vk_storage WHERE chatid = $1 and service = $2", id, service)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}
}
