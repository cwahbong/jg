package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cwahbong/jg/backend"
	"labix.org/v2/mgo"
	"sort"
)

func findMongoName(names []string, name string) bool {
	var idx = sort.SearchStrings(names, name)
	return (idx != len(names)) && (names[idx] == name)
}

type initdbArgs struct {
	reset  bool
	test   bool
	dbname string
}

func jgInitdb(argstrs []string) error {
	var args initdbArgs
	flagSet := flag.NewFlagSet("jg-initdb", flag.ExitOnError)
	flagSet.BoolVar(&args.reset, "r", false, "Remove old data from database.")
	flagSet.BoolVar(&args.test, "t", false, "Add testing data.")
	flagSet.StringVar(&args.dbname, "d", "jg", "Name of database.")
	flagSet.Parse(argstrs)

	session := backend.DefaultDial()
	db := session.DB(args.dbname)

	if args.reset {
		dbNames, _ := session.DatabaseNames()
		if findMongoName(dbNames, args.dbname) {
			db.DropDatabase()
		} else {
			return errors.New(fmt.Sprintf(`Database "%s" not found.`, args.dbname))
		}
	}

	defaultCollectionInfo := mgo.CollectionInfo{}
	db.C("tag").Create(&defaultCollectionInfo)
	db.C("problem").Create(&defaultCollectionInfo)
	// TODO make some index (such like tagged)

	if args.test {
		// TODO add test data
	}
	return nil
}
