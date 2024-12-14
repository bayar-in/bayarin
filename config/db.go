package config

import (
	"os"

	"github.com/gocroot/helper/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoStringGeo string = os.Getenv("MONGOSTRINGGEO")

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dsdatabase",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)

var MongoInfoGeo = atdb.DBInfo{
	DBString: MongoStringGeo,
	DBName:   "Geo",
}

var MongoconnGeo, ErrorMongoconnGeo = atdb.MongoConnect(MongoInfoGeo)