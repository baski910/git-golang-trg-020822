package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func getValues() (string, error) {
	time.Sleep(2 * time.Second)
	err := db.Exec("insert into products values(2,'mobile',15000)").Error
	if err != nil {
		return "", err
	}
	return "added", nil
}

func ShowValues(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("process completed")
	case <-ctx.Done():
		fmt.Println("process halted")
	}
}
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()
	var s string
	var err error
	db, _ = gorm.Open(mysql.Open("root:admin_12345@tcp(127.0.0.1:3306)/demodb_1"),
		&gorm.Config{})
	go func() {
		s, err = getValues()
		if err != nil {
			cancel()
		}
	}()
	ShowValues(ctx)
	fmt.Println("Done ...")

}
