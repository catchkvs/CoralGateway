package dbconnector

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cockroachdb/pebble"
)
var db *pebble.DB
func InitPebbleDB(dbDir string) {
	database, err := pebble.Open(dbDir, &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
	db = database

}

/**
 *
 */
func PutObject(collectionName, dataKey string, dataValue interface{}) {
	key := collectionName + "/" + dataKey
	keyBytes := []byte(key)
	data, e := json.Marshal(dataValue)
	if e!= nil {
		log.Fatal(e)
	}
	if err := db.Set(keyBytes, data, pebble.Sync); err != nil {
		log.Fatal(err)
	}
}

func GetObject(key string) interface{} {
	keyBytes := []byte(key)
	value, closer, err := db.Get(keyBytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", key, value)
	var result interface{}
	return json.Unmarshal(value, result)
	if err := closer.Close(); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
	return nil
}