package main 

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32 
}
func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
	}
	//db handel
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	
	//err
	pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
	fmt.Println("[+] Connection Established To Server...")
	//
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[+] Albums found -> %v\n", albums)
	// ID test
	alb, err := albumByID(2)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("[+] Album found ->  %v\n", alb)


}