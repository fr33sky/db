// +build !generated

package postgresql

import (
	"log"
	"testing"

	"upper.io/db.v2/lib/sqlbuilder"
)

func mustOpen() sqlbuilder.Database {
	return nil
}

func TestMain(*testing.M) {
	log.Fatal(`Tests use generated code and a custom database, please use "make test".`)
}
