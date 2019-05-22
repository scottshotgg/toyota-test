// Could do this with a websocket, but we'll just use a REST handler for now

package inmem

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/pkg/errors"
	"github.com/scottshotgg/toyota-test/models"
)

// DB is the structure used for in memory storage of the database
type DB struct {
	sync.Mutex
	store map[string]models.Currency
}

type symbolCurrency struct {
	BaseCurrency       string `json:"baseCurrency"`
	FeeCurrency string `json:"feeCurrency"`
}

type priceCurrency struct {
	Ask         string
	Bid         string
	Last        string
	Open        string
	Low         string
	High        string
	Volume      string
	VolumeQuote string
	Timestamp   string
	Symbol      string
}

const (
	base       = "https://api.hitbtc.com/api/2/public/"
	symbolPath = "symbol"
	pricePath  = "ticker"
)

var (
	db DB
)

func retrieveAll() (map[string][]models.Currency, error) {
	// Get all currencies
	var res, err = http.Get(base + symbolPath)
	if err != nil {
		err = errors.Wrap(err, "http.Get")
		log.Printf("Error: %+v\n", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadAll")
		log.Printf("Error: %+v\n", err)
		return nil, err
	}

	var symbols []symbolCurrency

	err = json.Unmarshal(body, &symbols)
	if err != nil {
		err = errors.Wrap(err, "json.Unmarshal")
		log.Printf("Error: %+v\n", err)
		return nil, err
	}

	// Get the prices for all tickers
	res, err = http.Get(base + pricePath)
	if err != nil {
		err = errors.Wrap(err, "http.Get")
		log.Printf("Error: %+v\n", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadAll")
		log.Printf("Error: %+v\n", err)
		return nil, err
	}

	var tickers []priceCurrency

	err = json.Unmarshal(body, &tickers)
	if err != nil {
		err = errors.Wrap(err, "json.Unmarshal")
		log.Printf("Error: %+v\n", err)
		return nil, err
	}

	// Convert the array to a map
	var currencyMap = map[string]models.Currency{}
	for _, symbol := range symbols {
		currencyMap[symbol.] = model.Currency{

		}
	}

	return nil, nil
}

var symbols = []string{"BTCUSD", "ETHBTC"} 

// New creates a new in memory db
func New(symbols []string) (*DB, error) {
	var (
		db = &DB{
			store: map[string]models.Currency{},
		}
		err = db.Sync()
	)

	if err != nil {
		err = errors.Wrap(err, "db.Sync")
		log.Printf("Error: %+v\n", err)
		return nil, err
	}

	return db, nil
}

// Sync serves to synchronize the DB
func (db *DB) Sync() error {
	var currencyArray, err = retrieveAll()
	if err != nil {
		err = errors.Wrap(err, "retrieveAll")
		log.Printf("Error: %+v\n", err)
		return err
	}

	for _, currency := range currencyArray {
		db.Insert(currency)
	}

	return nil
}

// Dump removes all data from the database
func (db *DB) Dump() {
	db.Lock()
	defer db.Unlock()

	db.store = map[string]models.Currency{}
}

// Insert is used to inject new data into the in memory db
func (db *DB) Insert(c models.Currency) {
	db.Lock()
	defer db.Unlock()

	db.store[c.ID] = c
}
