package handler

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"

	cloudsqlproxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/go-sql-driver/mysql"
)

type schema struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
}

func (s *schema) ToMap() map[string]string {
	return map[string]string{
		"message":   s.Message,
		"name":      s.Name,
		"id":        strconv.Itoa(s.ID),
		"timestamp": s.Timestamp.Format(time.RFC822),
	}
}

// AddRecords POSTパラメータを使ってDBに新しいデータを追加する
func AddRecords(c echo.Context) error {
	newdat := new(schema)
	err := c.Bind(newdat)
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = addRecords(ctx, *newdat)
	return err
}

// Env 環境変数から読み込む値
type Env struct {
	DatabaseAddress string `required:"true" split_words:"true"`
}

func connDatabase() (*sql.DB, error) {
	var goenv Env
	if err := envconfig.Process("", &goenv); err != nil {
		return nil, err
	}

	db, err := cloudsqlproxy.DialCfg(&mysql.Config{
		Addr:                 goenv.DatabaseAddress, // インスタンス接続名
		DBName:               "mydb",                // データベース名
		User:                 "user",                // ユーザ名
		Passwd:               "user_password",       // ユーザパスワード
		Net:                  "cloudsql",            // Cloud SQL Proxy で接続する場合は cloudsql 固定です
		ParseTime:            true,                  // DATE/DATETIME 型を time.Time へパースする
		TLSConfig:            "",                    // TLSConfig は空文字を設定しなければなりません
		AllowNativePasswords: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetRecords DBからデータを取得してくる
func GetRecords(ctx context.Context) ([]map[string]string, error) {
	db, err := connDatabase()
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SELECT * FROM chatlog ORDER BY id DESC LIMIT 20")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dat := []map[string]string{}
	for rows.Next() {
		var d schema
		if err := rows.Scan(&(d.ID), &(d.Timestamp), &(d.Name), &(d.Message)); err != nil {
			return nil, err
		}
		dat = append(dat, d.ToMap())
	}

	return dat, nil
}

func addRecords(ctx context.Context, d schema) error {
	db, err := connDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, "INSERT INTO chatlog VALUES (?,?,?,?);", 0, nil, d.Name, d.Message)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
