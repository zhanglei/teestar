package api

import (
	"runtime"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var adapter *Adapter

func init() {
	adapter = NewAdapter("mysql", beego.AppConfig.String("dataSourceName"))
}

type UserRepo struct {
	Id   int64
	User string `xorm:"varchar(100) notnull"`
	Repo string `xorm:"varchar(100) notnull"`
}

type UserStarringRepo struct {
	User string `xorm:"varchar(100) notnull pk"`
	Repo string `xorm:"varchar(100) notnull pk"`
}

type User struct {
	User         string `xorm:"varchar(100) unique pk"`
	Password     string `xorm:"varchar(100)"`
	Hitter       string `xorm:"varchar(100)"`
	QQ           string `xorm:"varchar(100)"`
	CreatedAt    string `xorm:"varchar(100)"`
	Nickname     string `xorm:"varchar(100)"`
	Email        string `xorm:"varchar(100)"`
	IsAdmin      bool
	IsDisabled   bool
	IsFollowable bool
}

type UserFollowingTarget struct {
	User string `xorm:"varchar(100) notnull pk"`
	Target string `xorm:"varchar(100) notnull pk"`
}

type ExtendedUser struct {
	User          string
	Hitter        string
	QQ            string
	CreatedAt     string
	Nickname      string
	Email         string
	IsAdmin       bool
	IsDisabled    bool
	RepoCount     int
	StarringCount int
	StarredCount  int
	OweCount      int
}

type Repo struct {
	Name       string
	Stargazers []string
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

	_, err = engine.Exec("CREATE DATABASE IF NOT EXISTS gitstar default charset utf8 COLLATE utf8_general_ci")
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

	err = a.engine.Sync2(new(User))
	if err != nil {
		panic(err)
	}
}
