package controllers

import (
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type UserRepo struct {
	User string `xorm:"varchar(100)"`
	Repo string `xorm:"varchar(100)"`
}

type UserStarringRepo struct {
	User string `xorm:"varchar(100)"`
	Repo string `xorm:"varchar(100)"`
}

type UserHitter struct {
	User   string `xorm:"varchar(100)"`
	Hitter string `xorm:"varchar(100)"`
}

// Adapter represents the MySQL adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	engine         *xorm.Engine
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
	a.engine.Close()
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(driverName string, dataSourceName string) *Adapter {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

func (a *Adapter) createDatabase() error {
	engine, err := xorm.NewEngine(a.driverName, a.dataSourceName)
	if err != nil {
		return err
	}
	defer engine.Close()

	_, err = engine.Exec("CREATE DATABASE IF NOT EXISTS gitstar")
	return err
}

func (a *Adapter) open() {
	if err := a.createDatabase(); err != nil {
		panic(err)
	}

	var engine *xorm.Engine
	engine, err := xorm.NewEngine(a.driverName, a.dataSourceName+"gitstar")
	if err != nil {
		panic(err)
	}

	a.engine = engine
	a.createTable()
}

func (a *Adapter) close() {
	a.engine.Close()
	a.engine = nil
}

func (a *Adapter) createTable() {
	err := a.engine.Sync2(new(UserRepo))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(UserStarringRepo))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(UserHitter))
	if err != nil {
		panic(err)
	}
}
