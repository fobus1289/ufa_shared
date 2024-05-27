package pg

import (
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var instance *instanceHolder

// global state instance database
type instanceHolder struct {
	*sqlx.DB
}

func (i *instanceHolder) clone() *instanceHolder {
	return &instanceHolder{DB: i.DB}
}

// TODO: add max connections | settings
// db.SetConnMaxIdleTime(n)
// db.SetMaxOpenConns(n)
// db.SetMaxIdleConns(n)
// db.SetConnMaxLifetime(n)
func connect(cfg Config) (*instanceHolder, error) {
	// singleton pattern
	if instance != nil {
		return instance.clone(), nil
	}

	db, err := sqlx.Connect("pgx", cfg.build())
	if err != nil {
		return nil, err
	}

	instance = &instanceHolder{db}

	return instance.clone(), nil
}

// max try 20.
// max timeout 10.
// reconnection time cannot be more than 10 seconds one step
// the number of attempts can be less than or equal to 20
// connection time out
func RetryConnect(cfg Config, tryCount uint, timeout uint64) (*instanceHolder, error) {
	if timeout > 10 {
		return nil, errors.New("reconnection time cannot be more than 10 seconds one step")
	}

	if tryCount > 20 {
		return nil, errors.New("the number of attempts can be less than or equal to 20")
	}

	timeoutSecond := time.Second * time.Duration(timeout)

	for i := uint(1); i <= tryCount; i++ {
		db, err := connect(cfg)

		if err == nil {
			return db, nil
		}

		time.Sleep(timeoutSecond)
	}

	return nil, errors.New("connection time out")
}

// value or error
func Connect(cfg Config) (*instanceHolder, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return db, err
}

// value or panic
func MustConnect(cfg Config) *instanceHolder {
	db, err := connect(cfg)
	if err != nil {
		panic(err)
	}

	return db
}
