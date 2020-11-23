package db

import "fmt"

type Database struct {
	filePath string
}

func New(filePath string) *Database {
	return &Database{
		filePath: filePath,
	}
}

func (d *Database) FilePath() string {
	return d.filePath
}

func (d *Database) AddRecord(record interface{}) {
	fmt.Printf("Added Record: %v", record)
}

func (d *Database) Reset() bool {
	d.clear()
	return true
}

func (d *Database) clear() {
	// You shouldn't access this from the Database outside of this package.
}
