package models

import (
	"database/sql"
	"strings"
	"log"
	_ "github.com/go-sql-driver/mysql"
)
func InitDatabase() (*sql.DB, error) {
    //将数据转换成数据库url作为返回值
    url := strings.Join([]string{"root", ":", "123123", "@tcp(", "127.0.0.1", ":", "3306",")/", "shop"}, "")
    db, err := sql.Open("mysql", url)
    if err != nil {
        log.Printf("open database error:%v", err)
        return nil, err
    }
    return db, nil
}
//执行增、改、删任务
func Execute(db *sql.DB, sql string, params ...interface{}) error {
    stmt, _ := db.Prepare(sql) //预编译
    defer stmt.Close()
    _, err := stmt.Exec(params...)
    if err != nil {
        log.Printf("execute sql error:%v\n", err)
        return err
    }
    log.Println("execute sql success")
    return nil
}

//查询数据库数据
func QueryData(db *sql.DB, sql string, params ...interface{}) (*sql.Rows, error) {
    stmt, _ := db.Prepare(sql)
    defer stmt.Close()
    rows, err := stmt.Query(params...)
    if err != nil {
        log.Printf("query data error:%v", err)
        return nil, err
    }
    log.Println("query data success")
    return rows, nil
}
