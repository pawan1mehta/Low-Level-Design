package main

import (
	"sync"
)

type Database struct {
	name string
}

var (
	dbInstance *Database
	once       sync.Once
)

func GetDatabase() *Database {
	once.Do(func() {
		dbInstance = &Database{
			name: "MySQL Database",
		}
	})

	return dbInstance
}

func main() {
	_ = GetDatabase()

}
