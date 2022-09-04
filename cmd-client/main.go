package main

import (
	"log"
	"os"
	"fmt"

	cli "github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Date string
	Comment string
	Amount  float64
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "language for the greeting",
				EnvVars: []string{"LEGACY_COMPAT_LANG", "APP_LANG", "LANG"},
			},
		},
		Action: func(c *cli.Context) error {
			name := "someone"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				log.Printf("Hola %s!", name)
			} else {
				log.Printf("Hello %s!", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func addRecord(r Record) string{
	db := initDB()
	db.Create(&r)
	return "Successfully added record"
}

func getAllRecords() []Record{
	db := initDB()
	var records []Record
	db.Find(&records)
	return records
}

func getRecordByDate(date string) []Record{
	db := initDB()
	var records []Record
	db.Where("date = ?", date).Find(&records)
	return records
}

func displayTable(records []Record){
	fmt.Println("Date\t\tComment\t\tAmount")
	for _, record := range records {
		fmt.Println(record.Date, record.Comment, record.Amount)
	}
}

