package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDb() error {
	var err error
	dsn := "root:root@tcp(localhost:3306)/golang_db"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"string"`
	Age  int            `db:"age"`
}

func testQueryData() {
	for i := 0; i < 101; i++ {
		fmt.Printf("query %d times\n", i)
		sqlstr := "select id, name, age from user where id=?"
		row := DB.QueryRow(sqlstr, 2)
		/*if row != nil {
			continue
		}*/
		// row := DB.QueryRow(sqlstr,2) 后，一定要把它scan掉
		// 因为前面设置了MaxOpenConns==100
		// 如果直接continue，第101次就会崩溃
		var user User
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}

		fmt.Printf("id:%d name:%v age:%d\n", user.Id, user.Name, user.Age)
	}

}

func testQueryMultilRow() {
	sqlstr := "select id, name, age from user where id > ?" // 多行查询
	rows, err := DB.Query(sqlstr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	//注意：rows对象一定要Close掉
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%v\n", user)
	}
}

func testInsertData() {
	sqlstr := "insert into user(name, age) values(?, ?)"
	result, err := DB.Exec(sqlstr, "tom", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func testUpdateData() {
	sqlstr := "update user set name=? where id=?"
	result, err := DB.Exec(sqlstr, "jim", 3)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed, err:%v\n", err)
	}
	fmt.Printf("update db successfully, affected rows:%d\n", affected)
}

func testDeleteData() {
	sqlstr := "delete from user where id=?"
	result, err := DB.Exec(sqlstr, 3)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed, err:%v\n", err)
	}
	fmt.Printf("delete successfully, affected rows:%d\n", affected)
}

/*
	一般sql处理流程
	1. 客户端拼接好sql语句
	2. 客户端发送sql语句到mysql服务器
	3. mysql服务器解析sql语句并执行，把执行结果发送给客户端

	预处理流程
	1. 把sql分为两部分，`命令部分` 和 `数据部分`
	2. 首先把 `命令部分` 发送给mysql服务器，mysql进行sql预处理
	3. 然后把 `数据部分` 发送给mysql服务器，mysql进行占位符替换
	4. mysql服务器执行sql语句并返回结果给客户端

	预处理优势：
		同一条sql语句反复执行，性能会很高
		避免sql注入问题
*/

func testPrepareQueryData() {
	sqlstr := "select id, name, age from user where id > ?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	rows, err := stmt.Query(0)
	//rows对象一定要Close掉
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

func testPrepareInsertData() {
	sqlstr := "insert into user(name, age) values(?, ?)"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Exec("jim", 30)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

/* Mysql 事务
应用场景：
1. 同时更新，多个表
2. 同时更新多行数据

事务的ACID：
1. 原子性：要么都成功，要么都失败
2. 一致性：数据是一致的，不会错乱
3. 隔离性：多个事务的修改，它们之间是隔离的
4. 持久性：不会因为程序的异常错误而导致数据丢失
*/

func testTrans() {
	conn, err := DB.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}

	sqlstr := "update user set age = 22 where id = ?"
	_, err = conn.Exec(sqlstr, 1)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}

	sqlstr = "update user set age = 102 where id = ?"
	_, err = conn.Exec(sqlstr, 2)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}
	err = conn.Commit()
	if err != nil {
		fmt.Printf("commit failed, err:%v\n", err)
		conn.Rollback()
		return
	}
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}

	//testQueryData()
	//testQueryMultilRow()
	//testInsertData()
	//testUpdateData()
	//testDeleteData()
	//testPrepareQueryData()
	//testPrepareInsertData()
	testTrans()
}
