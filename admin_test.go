package postgresql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func TestDemo(t *testing.T) {
	// 创建连接
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "192.168.31.19", 5432, "postgres", "aaaaaa", "postgres"))
	//db, err := sql.Open("postgres", fmt.Sprintf("%s:%s@tcp(%s)/%s", "postgres", "aaaaaa", "192.168.31.19:5432", "postgres"))
	if err != nil {
		fmt.Printf("sql.Open, err: %+v\n", err)
		return
	}
	defer db.Close()

	if pingErr := db.Ping(); pingErr != nil {
		fmt.Printf("ping.err: %+v\n", pingErr)
		return
	}

	rows := db.QueryRow("SELECT id, user_name FROM demo;")
	var id *sql.NullInt16
	var userName *sql.NullString
	err = rows.Scan(&id, &userName)
	if err != nil {
		fmt.Printf("Scan, err: %+v\n", err)
		return
	}
	if id.Int16 == 0 || userName.String == "" {
		fmt.Printf("Scan.data.null\n")
		return
	}

	fmt.Println("id:", id.Int16)
	fmt.Println("user_name:", userName.String)
}
