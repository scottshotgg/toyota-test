package handlers

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/scottshotgg/toyota-test/inmem"
	"github.com/scottshotgg/toyota-test/models"
	"github.com/scottshotgg/toyota-test/restapi/operations"
)

var (
	db  *inmem.DB
	err error
)

func init() {
	dir, derr := os.Getwd()
	if derr != nil {
		err = errors.Wrap(err, "os.Getwd")
		log.Printf("Error: %+v\n", err)

		os.Exit(9)
	}

	// Read in the file
	var file, err = ioutil.ReadFile(dir + "/handlers/get/symbols")
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadFile")
		log.Printf("Error: %+v\n", err)

		os.Exit(9)
	}

	var lines = strings.Split(string(file), "\n")

	db, err = inmem.New(lines)
	if err != nil {
		err = errors.Wrap(err, "inmem.New")
		log.Printf("Error: %+v\n", err)

		os.Exit(9)
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)

			err = db.Sync()
			if err != nil {
				err = errors.Wrap(err, "db.Sync")
				log.Printf("Error: %+v\n", err)

				os.Exit(9)
			}
		}
	}()
}

// GetSymbol handles retrieving data to a single endpoint
func GetSymbol(params operations.GetCurrencySymbolParams) middleware.Responder {
	log.Printf("User requested: %s\n", params.Symbol)

	var currency = db.Get(strings.ToUpper(params.Symbol))
	if currency == nil {
		return operations.NewGetCurrencySymbolNotFound()
	}

	return operations.NewGetCurrencySymbolOK().WithPayload(currency)
}

// GetAll handles retrieving data for every currency
func GetAll(params operations.GetCurrencyAllParams) middleware.Responder {
	log.Println("User requested all")

	return operations.NewGetCurrencyAllOK().WithPayload(&models.Currencies{
		Currencies: db.All(),
	})
}
