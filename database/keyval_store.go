package database

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type keyValue struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

// SetKey updates or inserts a key with the provided value
func SetKey(key string, value string) error {
	newPair := keyValue{Key: key, Value: value}
	err := mgoRequest("key_value_store", func(c *mgo.Collection) error {
		return c.Insert(newPair)
	})
	if err != nil {
		return err
	}

	return nil
}

// GetKey fetches the given key's value from the database
func GetKey(key string) (string, error) {
	var keyValues []keyValue
	err := mgoRequest("key_value_store", func(c *mgo.Collection) error {
		return c.Find(nil).Sort("-created_when").All(&keyValues)
	})
	if err != nil {
		return "", err
	}
	if len(keyValues) == 0 {
		return "", nil
	}

	return keyValues[0].Value, nil
}

// ClearKey removes the given key's row from the database
func ClearKey(key string) error {
	err := mgoRequest("key_value_store", func(c *mgo.Collection) error {
		_, err := c.RemoveAll(bson.M{"key": bson.M{"$eq": key}})
		return err
	})

	return err
}
