package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB

func InitDB() {
	host := Cfg.Section("mysql").Key("host").String()
	port := Cfg.Section("mysql").Key("port").String()
	user := Cfg.Section("mysql").Key("user").String()
	passwd := Cfg.Section("mysql").Key("password").String()
	database := Cfg.Section("mysql").Key("database").String()

	db = newMysql(host, port, user, passwd, database)
}

func newMysql(host, port, user, passwd, database string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, passwd, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		Logger().Panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(300)
	return db
}

func query(d *sql.DB, sqlstr string, params ...interface{}) ([]map[string]string, error) {
	stmt, ep := d.Prepare(sqlstr)
	if ep != nil {
		return nil, ep
	}
	defer stmt.Close()

	rows, eq := stmt.Query(params...)
	if eq != nil {
		return nil, eq
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	cols := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))
	for v := range cols {
		scans[v] = &cols[v]
	}

	var result []map[string]string
	for rows.Next() {
		_ = rows.Scan(scans...)
		row := make(map[string]string)
		for i, col := range cols {
			row[columns[i]] = string(col)
		}
		result = append(result, row)
	}
	return result, nil
}

func execute(d *sql.DB, sqlstr string, params ...interface{}) error {
	stmt, ep := d.Prepare(sqlstr)
	if ep != nil {
		return ep
	}
	defer stmt.Close()

	rows, ee := stmt.Exec(params...)
	if ee != nil {
		return ee
	}

	_, er := rows.RowsAffected()
	if er != nil {
		return er
	}
	return nil
}

func Query(sql string, params ...interface{}) ([]map[string]string, error) {
	return query(db, sql, params...)
}

func Execute(sql string, params ...interface{}) error {
	return execute(db, sql, params...)
}
