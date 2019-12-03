package sqldb

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"os"
	"sync"
	"testing"
)

var testdb *sqlx.DB
var once sync.Once

func TestMain(m *testing.M) {
	if err := beforeTest(); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func beforeTest() (err error) {
	once.Do(func() {
		testdb = sqlx.MustOpen("sqlite3", "file:test.db?cache=shared&mode=memory")
	})
	if _, err := testdb.Exec(`
		CREATE TABLE todo (
		  id TEXT PRIMARY KEY,
          title TEXT NOT NULL DEFAULT NULL,
		  description TEXT DEFAULT NULL,
		  done BOOLEAN NOT NULL CHECK (done IN (0,1)) DEFAULT 0,
		  created_at TEXT DEFAULT CURRENT_TIMESTAMP,
		  updated_at TEXT DEFAULT CURRENT_TIMESTAMP
		)
	`); err != nil {
		return errors.Wrap(err, "failed to create table 'todo'")
	}
	return nil
}
