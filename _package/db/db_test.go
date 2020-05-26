package db

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestDb(t *testing.T) {
	database, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Printf("type: %T, value: %+v\n", err, err)
	}

	t.Run("instert", func(t *testing.T) {
		r, err := database.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		id, err := r.LastInsertId()
		assert.True(t, int(id) > 0)
	})

	t.Run("select", func(t *testing.T) {
		type Person struct {
			UserId   int    `db:"user_id"`
			Username string `db:"username"`
			Sex      string `db:"sex"`
			Email    string `db:"email"`
		}
		var person []Person
		err := database.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
		assert.Nil(t, err)
	})

	t.Run("update ", func(t *testing.T) {
		res, err := database.Exec("update person set username=? where user_id=?", "stu0003", 1)
		assert.Nil(t, err)

		row, err := res.RowsAffected()
		assert.Nil(t, err)
		fmt.Println("update succ:", row)
	})

	t.Run("delete  ", func(t *testing.T) {
		res, err := database.Exec("delete from person where user_id=?", 1)
		assert.Nil(t, err)

		row, err := res.RowsAffected()
		assert.Nil(t, err)
		fmt.Println("update succ:", row)
	})

	t.Run("transaction", func(t *testing.T) {
		conn, err := database.Begin()
		if err != nil {
			fmt.Println("begin failed :", err)
			return
		}
		r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
		if err != nil {
			fmt.Println("exec failed, ", err)
			conn.Rollback()
			return
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Println("exec failed, ", err)
			conn.Rollback()
			return
		}
		fmt.Println("insert succ:", id)

		r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
		if err != nil {
			fmt.Println("exec failed, ", err)
			conn.Rollback()
			return
		}
		id, err = r.LastInsertId()
		if err != nil {
			fmt.Println("exec failed, ", err)
			conn.Rollback()
			return
		}
		fmt.Println("insert succ:", id)

		conn.Commit()
	})
}
