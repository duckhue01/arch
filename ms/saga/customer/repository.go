package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Repository struct {
	db *sql.DB
}

type Customer struct {
	id     int
	amount int
}

func NewDatabaseStore() *Repository {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "195106",
		DBName:               "customer",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// test if connect successfully
	if err := db.Ping(); err != nil {
		log.Fatal(err, "ping")
	}

	return &Repository{
		db: db,
	}
}

func (d *Repository) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := d.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (re *Repository) GetOrder(id int) (err error, res *Customer) {
	ctx := context.Background()
	c, err := re.connect(ctx)
	if err != nil {
		return err, nil
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "SELECT * FROM order WHERE `id` = ? LIMIT 1", id)
	if err != nil {
		return err, nil
	}
	defer rows.Close()

	rows.Scan(&res)

	return nil, res

}
