package stat

import (
	"bot/models"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	login    = "postgres"
	password = "1"
	dbname   = "postgres"
)

func DbUpdate(arga string) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	today := time.Now()
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully connected!")

	userSql := "INSERT INTO users(nic,ferstReq,qreq)values ($1,$2,$3)"
	_, err = db.Exec(userSql, arga, today, 0)
	if err != nil {
		userSql := "update Users set qreq = qreq + 1  where nic = $1"
		_, err := db.Exec(userSql, arga)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func Stat(arga string) (argb int, argc string) {

	DbUpdate(arga)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	User := models.Users{}
	rows := db.QueryRow("select qreq,ferstReq from users where nic = $1", arga)

	err = rows.Scan(&User.Qreq, &User.FerstReq)
	if err != nil {
		fmt.Println(err)
	}
	return User.Qreq, User.FerstReq

}
