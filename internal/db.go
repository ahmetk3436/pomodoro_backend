package internal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(connString string) (*DB, error) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

type myService struct {
	db DB
}

func NewMyService(connString string) (*myService, error) {
	db, err := NewDB(connString)
	if err != nil {
		return nil, err
	}
	return &myService{db: *db}, nil
}
