package inmem_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/toyota-test/inmem"
)

func TestNew(t *testing.T) {
	var db, err = inmem.New()
	if err != nil {
		t.Fatalf("Error creating new database")
	}

	fmt.Println(db)
}
