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
	supportedSymbols []string
	store            map[string]*models.Currency
}

type symbolCurrency struct {
	BaseCurrency string `json:"baseCurrency"`
	FeeCurrency  string `json:"feeCurrency"`
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
	base         = "https://api.hitbtc.com/api/2/public/"
	symbolPath   = "symbol/"
	pricePath    = "ticker/"
	currencyPath = "currency/"
)

var (
	db DB
)

// New creates a new in memory db
func New(supportedSymbols []string) (*DB, error) {
	var (
		db = &DB{
			store:            map[string]*models.Currency{},
			supportedSymbols: supportedSymbols,
		}
	)

	for _, symbol := range supportedSymbols {
		// Get all currencies
		var res, err = http.Get(base + symbolPath + symbol)
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

		var s symbolCurrency

		err = json.Unmarshal(body, &s)
		if err != nil {
			err = errors.Wrap(err, "json.Unmarshal")
			log.Printf("Error: %+v\n", err)
			return nil, err
		}

		var currency = &models.Currency{
			FeeCurrency: s.FeeCurrency,
		}

		// Get all currencies
		res, err = http.Get(base + pricePath + symbol)
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

		var p priceCurrency

		err = json.Unmarshal(body, &p)
		if err != nil {
			err = errors.Wrap(err, "json.Unmarshal")
			log.Printf("Error: %+v\n", err)
			return nil, err
		}

		currency.Ask = p.Ask
		currency.Bid = p.Bid
		currency.Last = p.Last
		currency.High = p.High
		currency.Low = p.Low
		currency.Open = p.Open

		// Get all currencies
		res, err = http.Get(base + currencyPath + s.BaseCurrency)
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

		// Unmarshal the ID and FullName into currency
		err = json.Unmarshal(body, &currency)
		if err != nil {
			err = errors.Wrap(err, "json.Unmarshal")
			log.Printf("Error: %+v\n", err)
			return nil, err
		}

		db.Insert(currency)
	}

	return db, nil
}

// Sync serves to synchronize the DB
func (db *DB) Sync() error {
	var dbb, err = New(db.supportedSymbols)
	if err != nil {
		err = errors.Wrap(err, "New")
		log.Printf("Error: %+v\n", err)
		return err
	}

	db.Lock()
	db.store = dbb.store
	db.Unlock()

	return nil
}

func (db *DB) Get(name string) *models.Currency {
	db.Lock()
	defer db.Unlock()

	return db.store[name]
}

func (db *DB) All() []*models.Currency {
	db.Lock()
	defer db.Unlock()

	var currencies []*models.Currency

	for _, currency := range db.store {
		currencies = append(currencies, currency)
	}

	return currencies
}

// Dump removes all data from the database
func (db *DB) Dump() {
	db.Lock()
	defer db.Unlock()

	db.store = map[string]*models.Currency{}
}

// Insert is used to inject new data into the in memory db
func (db *DB) Insert(c *models.Currency) {
	db.Lock()
	defer db.Unlock()

	db.store[c.ID] = c
}
