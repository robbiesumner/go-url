package db_store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGet(t *testing.T) {
	db := NewDBClient()

	longUrl := "http://www.google.com"
	shortUrl := "myshort"

	db.Set(shortUrl, longUrl)
	result, err := db.Get(shortUrl)

	assert.Nil(t, err)
	assert.Equal(t, longUrl, result)
}

func TestGetNotFound(t *testing.T) {
    db := NewDBClient()

    result, err := db.Get("not set short url")

    assert.NotNil(t, err)
    assert.Equal(t, "", result)
}

func TestSetDuplicate(t *testing.T) {
    db := NewDBClient()

    longUrl := "http://www.google.com"
    shortUrl := "myshort"

    db.Set(shortUrl, longUrl)
    err := db.Set(shortUrl, longUrl)

    assert.NotNil(t, err)
}
