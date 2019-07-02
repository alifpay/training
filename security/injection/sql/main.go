package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strings"
	"syscall"

	"github.com/jackc/pgx"
)

type userCriteria struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Status   string `json:"status"`
	UserRole string `json:"user_role"`
}

/*
SQL, NoSQL or other type of database input arguments must be validated for injection
*/
var pgPool *pgx.ConnPool

func main() {

	err := connDB("127.0.0.1", "postgres", "", "alifdb")
	if err != nil {
		log.Fatal(err)
	}
	eml := "''; DROP TABLE users;'';"
	id := 0
	nm := ""

	//very bad
	err = pgPool.QueryRow("SELECT id, name FROM users WHERE email="+eml).Scan(&id, &nm)
	if err != nil {
		log.Println(err)
	}

	//very good
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(eml) {
		log.Fatalln("invalid email")
	}
	err = pgPool.QueryRow("SELECT id, name FROM users WHERE email=$1", eml).Scan(&id, &nm)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(id, nm)

	//multi argument for query

	in := userCriteria{}

	var args pgx.QueryArgs

	sql := "SELECT id, name FROM users WHERE 1=1"
	in.FullName = strings.TrimSpace(in.FullName)
	if len(in.FullName) > 0 {
		args.Append(in.FullName)
		sql += " AND fnm like %" + args.Append(in.FullName) + "%"
	}
	in.Status = strings.TrimSpace(in.Status)
	if len(in.Status) > 0 {
		args.Append(in.Status)
		sql += " AND status = " + args.Append(in.Status)
	}
	rows, err := pgPool.Query(sql, args...)
	for rows.Next() {
		err = rows.Scan(&id, &nm)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(id, nm)
	}
	sigint := make(chan os.Signal)
	signal.Notify(sigint, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	<-sigint
	closeDB()
}

// Connect connect to postgres and adds new connection pool
func connDB(dbhost, dbuser, dbpass, dbname string) error {
	var err error

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     dbhost,
			User:     dbuser,
			Password: dbpass,
			Database: dbname,
		},
		MaxConnections: runtime.NumCPU() * 2,
	}
	pgPool, err = pgx.NewConnPool(connPoolConfig)
	if err != nil {
		return err
	}
	return nil
}

// Close db connection
func closeDB() {
	pgPool.Close()
}
