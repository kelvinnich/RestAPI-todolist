package manager

import (
	"authenctications/config"
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
)

type InfraManager interface {
	ConnecDB() *sql.DB
}

type infraManager struct {
	db *sql.DB
	cfg config.Config
}

func(i *infraManager) initDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.DBName)

	db,err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("failed to connect database %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("error ping %v", err)
	}

	i.db = db
	log.Println("succes to running app...")
}

func(i *infraManager)ConnecDB() *sql.DB{
	return i.db
}

func NewInfraManager(cfg config.Config) InfraManager{
	infra := infraManager{
		cfg: cfg,
	}

	infra.initDB()
	return &infra
}