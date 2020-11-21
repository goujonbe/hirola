package cmd

import (
	"database/sql"
	"flag"
	"log"
	"strings"
)

func NewLoadCommand() *LoadCommand {
	lc := &LoadCommand{
		fs: flag.NewFlagSet("load", flag.ContinueOnError),
	}

	lc.fs.StringVar(&lc.csvPath, "csv-path", "", "absolute csv path")
	lc.fs.StringVar(&lc.connectionString, "to", "", "db url")
	lc.fs.StringVar(&lc.table, "table", "", "destination table name")

	return lc
}

type LoadCommand struct {
	fs *flag.FlagSet

	csvPath          string
	connectionString string
	table            string
}

func (l *LoadCommand) Init(args []string) error {
	return l.fs.Parse(args)
}

func (l *LoadCommand) Run() error {
	db, err := sql.Open("postgres", l.connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: build query differently to prevent SQL injection
	builder := strings.Builder{}
	builder.WriteString("COPY ")
	builder.WriteString(l.table)
	builder.WriteString(" FROM '")
	builder.WriteString(l.csvPath)
	builder.WriteString("' csv header;")

	rows, err := db.Query(builder.String())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return nil
}

func (l *LoadCommand) Name() string {
	return l.fs.Name()
}
