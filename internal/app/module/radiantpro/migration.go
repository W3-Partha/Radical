package radiantpro

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"
	"strings"

	radicalLogger "github.com/W3-Partha/Radical/logger"
	"github.com/W3-Partha/Radical/utils"
)

var SQL utils.DocValue
var SQLMode utils.DocValue
var SQLModePath utils.DocValue

var (
	SQLModeUp   = "up"
	SQLModeDown = "down"
)

func (c *Container) Migration(args []string) {
	c.initUserOption()
	db, err := sql.Open(c.UserOption.Driver, c.UserOption.Dsn)
	if err != nil {
		radicalLogger.Log.Fatalf("Could not connect to '%s' database using '%s': %s", c.UserOption.Driver, c.UserOption.Dsn, err)
		return
	}
	defer db.Close()
	switch SQLMode.String() {
	case SQLModeUp:
		doByMode(db, "up.sql")
	case SQLModeDown:
		doByMode(db, "down.sql")
	default:
		doBySqlFile(db)
	}
}

func doBySqlFile(db *sql.DB) {
	fileName := SQL.String()
	if !utils.IsExist(fileName) {
		radicalLogger.Log.Fatalf("sql mode path not exist, path %s", SQL.String())
	}
	doDb(db, fileName)
}

func doByMode(db *sql.DB, suffix string) {
	pathName := SQLModePath.String()
	if !utils.IsExist(pathName) {
		radicalLogger.Log.Fatalf("sql mode path not exist, path %s", SQLModePath.String())
	}

	rd, err := ioutil.ReadDir(pathName)
	if err != nil {
		radicalLogger.Log.Fatalf("read dir err, path %s, err %s", pathName, err)
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			if !strings.HasSuffix(fi.Name(), suffix) {
				continue
			}
			doDb(db, filepath.Join(pathName, fi.Name()))
		}
	}
}

func doDb(db *sql.DB, filePath string) {
	absFile, _ := filepath.Abs(filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		radicalLogger.Log.Errorf("read file err %s, abs file %s", err, absFile)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		radicalLogger.Log.Errorf("db exec err %s", err)
	}
	radicalLogger.Log.Infof("db exec info %s", filePath)
}
