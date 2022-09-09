package cmd_client

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Date    string
	Comment string
	Amount  float64
}

func Client() {
	app := &cli.App{
		Usage: "A simple expense tracker",

		Commands: []*cli.Command{commands()},
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
	db.AutoMigrate(&Record{})
	return db
}

func addRecord(r Record) string {
	db := initDB()
	if err:=db.Create(&r).Error;err!=nil{
		return "Error adding record:"+err.Error()
	}
	return "Successfully added record"
}

func getAllRecords() []Record {
	db := initDB()
	var records []Record
	if err := db.Find(&records).Error; err != nil {
		panic("failed to get records")
	}
	return records
}

func getRecordByDate(date string) []Record {
	db := initDB()
	var records []Record
	if err := db.Where("date = ?", date).Find(&records).Error; err != nil {
		panic("failed to find record")
	}
	return records
}

func displayTable(records []Record) {
	fmt.Println("Date\t\tComment\t\tAmount")
	for _, record := range records {
		//print as a table
		fmt.Printf("%s\t%s\t\t%f\n", record.Date, record.Comment, record.Amount)
	}
}
//commands

func commands() *cli.Command{
	return &cli.Command{
		Name:  "add",
		Usage: "Add a record of your spending",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "date",
				Aliases:  []string{"d"},
				Usage:    "Date in format YYYY-MM-DD",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "comment",
				Aliases:  []string{"c"},
				Usage:    "Comment (What did you spend money on?)",
				Required: true,
			},
			&cli.Float64Flag{
				Name:     "amount",
				Aliases:  []string{"a"},
				Usage:    "Amount (The amount of money you spent)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			r := Record{
				Date:    c.String("date"),
				Comment: c.String("comment"),
				Amount:  c.Float64("amount"),
			}
			fmt.Println(addRecord(r))
			return nil
		},

	}
}

var getCommand = &cli.Command{
	Name:  "get",
	Usage: "Get records of your spending",
	Subcommands: []*cli.Command{
		{
			Name:  "all",
			Usage: "Get all records",
			Action: func(c *cli.Context) error {
				records := getAllRecords()
				displayTable(records)
				return nil
			},
		},
		{
			Name:  "date",
			Usage: "Get records by date",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "date",
					Aliases:  []string{"d"},
					Usage:    "Date in format YYYY-MM-DD",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				records := getRecordByDate(c.String("date"))
				displayTable(records)
				return nil
			}
		},
	}
}

func AddCommand(*cli.Command) {
	commands().Subcommands = append(commands().Subcommands, getCommand)
}
